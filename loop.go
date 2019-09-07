package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func loop(separator string, approximateCounting bool, theta float64, topnPrint int, flushMilliSec int64) {
	nBytes, nChunks := int64(0), int64(0)
	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)
	var approximateCountingThreshold = 1.0 / theta

	var counter Counter
	if approximateCounting {
		counter = NewApproximateCounter(approximateCountingThreshold)
	} else {
		counter = NewMapCounter()
	}

	printer := NewTablePrinter(flushMilliSec, topnPrint)

	previousTail := ""
	for {
		n, err := reader.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		nChunks++
		nBytes += int64(len(buf))
		s := previousTail + string(buf)
		previousTail = ""

		lines := strings.Split(s, "\n")
		if len(lines) > 0 {
			tail := lines[len(lines)-1]
			if tail != "" {
				previousTail = tail
			}
			lines = lines[:len(lines)-1]
		}

		for _, line := range lines {
			counter.increment(line)
		}

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		printer.print(counter, nBytes, nChunks, false)
	}

	printer.print(counter, nBytes, nChunks, true)
	printer.exit(counter)
}
