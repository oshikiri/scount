package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"code.cloudfoundry.org/bytefmt"
	"github.com/olekukonko/tablewriter"
)

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// TablePrinter is a printer of Counter object by tablewriter package
type TablePrinter struct {
	flushMilliSec       int64
	topnPrint           int
	lastFlushedDatetime time.Time
	table               tablewriter.Table
}

func (printer *TablePrinter) print(counter Counter, nBytes int64, nChunks int64, forcePrint bool) {
	currentDatetime := time.Now()
	diff := currentDatetime.Sub(printer.lastFlushedDatetime)
	if !forcePrint && diff.Nanoseconds() < 1000*1000*printer.flushMilliSec {
		return
	}

	printer.table.SetHeader([]string{"Name", "Count"})
	counts, countBase := counter.getCountingResult()
	sorted := sortMap(counts)
	end := Min(len(sorted), printer.topnPrint)

	ClearTerminal := "\033c"
	fmt.Fprint(os.Stderr, ClearTerminal)
	printer.table.ClearRows()

	for i, c := range sorted {
		if i >= end {
			break
		}
		printer.table.Append([]string{c.key, strconv.Itoa(c.value + countBase)})
	}

	byteSize := bytefmt.ByteSize(uint64(nBytes))
	caption := fmt.Sprintf("Read: %v", byteSize)
	printer.table.SetCaption(true, caption)

	printer.table.Render()
	printer.lastFlushedDatetime = currentDatetime
}

func (printer *TablePrinter) exit(counter Counter) {
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stdout, counter.toJSON())
}

// NewTablePrinter is a utility
func NewTablePrinter(flushMilliSec int64, topnPrint int) *TablePrinter {
	table := tablewriter.NewWriter(os.Stderr)

	printer := &TablePrinter{}
	printer.lastFlushedDatetime = time.Now()
	printer.flushMilliSec = flushMilliSec
	printer.topnPrint = topnPrint
	printer.table = *table
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
