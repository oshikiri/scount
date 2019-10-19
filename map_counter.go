package main

import (
	"encoding/json"
)

// MapCounter is a simple counter utility using hash map
type MapCounter struct {
	counts map[string]int
	Counter
}

func (counter *MapCounter) getCountingResult() map[string]int {
	return counter.counts
}

func (counter *MapCounter) initialize() {
	counter.counts = map[string]int{}
}

func (counter MapCounter) increment(item string) int {
	v, ok := counter.counts[item]
	if ok {
		counter.counts[item] = v + 1
	} else {
		counter.counts[item] = 1
	}
	return counter.counts[item]
}

func (counter MapCounter) get(item string) int {
	return counter.counts[item]
}

func (counter MapCounter) getSize() uint64 {
	return uint64(len(counter.counts))
}

func (counter MapCounter) toJSON() string {
	s, _ := json.Marshal(counter.counts)
	return string(s)
}

// NewMapCounter constructs MapCounter struct
func NewMapCounter() *MapCounter {
	counter := &MapCounter{}
	counter.initialize()
	return counter
}
