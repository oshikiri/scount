package main

// Counter is a base interface for counters
type Counter interface {
	initialize()
	increment(item string) int
	get(item string) int
	getCountingResult() map[string]int
	getSize() uint64
	toJSON() string
}
