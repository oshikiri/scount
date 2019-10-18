package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"

	"text/tabwriter"
	"time"

	"code.cloudfoundry.org/bytefmt"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// TablePrinter is a printer of Counter object
type TablePrinter struct {
	flushMilliSec       int64
	topnPrint           int
	lastFlushedDatetime time.Time
	tscreen             tcell.Screen
}

func (printer *TablePrinter) print(counter Counter, nBytes int64, nChunks int64, forcePrint bool) {
	printer.tscreen.Clear()

	currentDatetime := time.Now()
	diff := currentDatetime.Sub(printer.lastFlushedDatetime)
	if !forcePrint && diff.Nanoseconds() < 1000*1000*printer.flushMilliSec {
		return
	}

	var buffer bytes.Buffer
	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', tabwriter.Debug)

	counts := counter.getCountingResult()
	sorted := sortMap(counts)
	end := Min(len(sorted), printer.topnPrint)

	formatter := message.NewPrinter(language.English)
	maxCountLength := len(formatter.Sprintf("%v", sorted[0].value))
	countFormat := formatter.Sprintf("%%%dv", maxCountLength+1)

	for i, c := range sorted {
		if i >= end {
			break
		}
		count := formatter.Sprintf(countFormat, c.value)
		line := fmt.Sprintf(" %v\t%v\n", c.key, count)
		writer.Write([]byte(line))
	}

	byteSize := bytefmt.ByteSize(uint64(nBytes))
	caption := fmt.Sprintf("Read: %v", byteSize)
	writer.Write([]byte(caption))

	writer.Flush()

	s := buffer.String()
	runes := []rune(s)
	x := 0
	y := 0
	for _, r := range runes {
		if r != '\n' {
			printer.tscreen.SetCell(x, y, tcell.StyleDefault, r)
			x++
		} else {
			x = 0
			y++
		}
	}
	printer.tscreen.Show()

	printer.lastFlushedDatetime = currentDatetime
}

func (printer *TablePrinter) exit(counter Counter) {
	printer.tscreen.Fini()
}

// NewTablePrinter is a utility
func NewTablePrinter(flushMilliSec int64, topnPrint int) *TablePrinter {
	printer := &TablePrinter{}
	printer.lastFlushedDatetime = time.Now()
	printer.flushMilliSec = flushMilliSec
	printer.topnPrint = topnPrint

	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	printer.tscreen, _ = tcell.NewScreen()
	if err := printer.tscreen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	encoding.Register()

	printer.tscreen.SetStyle(tcell.StyleDefault)
	printer.tscreen.Clear()
	printer.tscreen.Show()

	return printer
}

func sortMap(m map[string]int) List {
	a := List{}
	for k, v := range m {
		e := Entry{k, v}
		a = append(a, e)
	}

	sort.Sort(sort.Reverse(a))
	return a
}

// Entry is key-value pair to sort
type Entry struct {
	key   string
	value int
}

// List for sort
type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].value == l[j].value {
		return (l[i].key < l[j].key)
	}
	return (l[i].value < l[j].value)
}
