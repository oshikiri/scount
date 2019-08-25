package main

// Printer is a base interface for TablePrinter and JSONPrinter
type Printer interface {
	print(counter Counter, nBytes int64, nChunks int64, forcePrint bool)
	exit(counter Counter)
}
