package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// JSONPrinter is a printer of Counter object by JSON format
type JSONPrinter struct {
	flushMilliSec       int64
	topnPrint           int
	lastFlushedDatetime time.Time
}

func (printer *JSONPrinter) print(counter Counter, nBytes int64, nChunks int64, forcePrint bool) {
	currentDatetime := time.Now()
	diff := currentDatetime.Sub(printer.lastFlushedDatetime)
	if !forcePrint && diff.Nanoseconds() < 1000*1000*printer.flushMilliSec {
		return
	}

	jsonString := counter.toJSON()
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(jsonString), "", "  ")
	if err != nil {
		panic(err)
	}
	indentJSON := buf.String()
	ClearTerminal := "\033c"

	fmt.Fprintf(os.Stderr, "%vBytes: %v,\t\t Chunks: %v\n%v", ClearTerminal, nBytes, nChunks, indentJSON)
	printer.lastFlushedDatetime = currentDatetime
}

func (printer *JSONPrinter) exit(counter Counter) {
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stdout, counter.toJSON())
}

// NewJSONPrinter is a utility
func NewJSONPrinter(flushMilliSec int64, topnPrint int) *JSONPrinter {
	printer := &JSONPrinter{}
	printer.flushMilliSec = flushMilliSec
	printer.topnPrint = topnPrint
	return printer
}
