package main

import (
	"bytes"
	"fmt"
	"os"

	"text/tabwriter"
	"time"

	"code.cloudfoundry.org/bytefmt"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

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
	elapsedNanoSecondsAfterFlushed := currentDatetime.Sub(printer.lastFlushedDatetime).Nanoseconds()
	if !forcePrint && elapsedNanoSecondsAfterFlushed < 1000*1000*printer.flushMilliSec {
		return
	}

	var buffer bytes.Buffer
	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', tabwriter.Debug)

	sortedCounts := sortMap(counter.getCountingResult())
	end := Min(len(sortedCounts), printer.topnPrint)

	formatter := message.NewPrinter(language.English)
	maxCountLength := len(formatter.Sprintf("%v", sortedCounts[0].value))
	countFormat := formatter.Sprintf("%%%dv", maxCountLength+1)

	for i, c := range sortedCounts {
		if i >= end {
			break
		}
		count := formatter.Sprintf(countFormat, c.value)
		line := fmt.Sprintf(" %v\t%v\n", c.key, count)
		writer.Write([]byte(line))
	}

	writer.Write([]byte(createCaption(nBytes, counter.getSize())))
	writer.Flush()

	printer.printOnTscreen(buffer.String())
	printer.lastFlushedDatetime = currentDatetime
}

func (printer TablePrinter) printOnTscreen(content string) {
	x := 0
	y := 0
	for _, r := range []rune(content) {
		if r != '\n' {
			printer.tscreen.SetCell(x, y, tcell.StyleDefault, r)
			x++
		} else {
			x = 0
			y++
		}
	}
	printer.tscreen.Show()
}

func (printer *TablePrinter) exit(counter Counter) {
	printer.tscreen.Fini()
}

func (printer *TablePrinter) initializeTscreen() {
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
}

// NewTablePrinter is a constructor of TablePrinter
func NewTablePrinter(flushMilliSec int64, topnPrint int) *TablePrinter {
	printer := &TablePrinter{}

	printer.lastFlushedDatetime = time.Now()
	printer.flushMilliSec = flushMilliSec
	printer.topnPrint = topnPrint
	printer.initializeTscreen()

	return printer
}

func createCaption(nBytes int64, nItems uint64) string {
	byteSize := bytefmt.ByteSize(uint64(nBytes))
	counterSizeString := bytefmt.ByteSize(nItems)
	caption := fmt.Sprintf("Read: %v, size of counter: %v", byteSize, counterSizeString)
	return caption
}
