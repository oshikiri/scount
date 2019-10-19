package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func loop(approximateCounting bool, epsilon float64, support float64, topnPrint int, flushMilliSec int64, quietMode bool) {
	nBytes, nChunks := int64(0), int64(0)
	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)

	var counter Counter
	if approximateCounting {
		counter = NewApproximateCounter(epsilon, support)
	} else {
		counter = NewMapCounter()
	}

	var printer *TablePrinter
	if !quietMode {
		printer = NewTablePrinter(flushMilliSec, topnPrint)
	}

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

		if !quietMode {
			printer.print(counter, nBytes, nChunks, false)
		}
	}

	if !quietMode {
		printer.exit(counter)
	}

	fmt.Fprintf(os.Stdout, counter.toJSON())
}
