package main

import "encoding/json"

// ApproximateCounter is a counter using streaming counting algorithm
//
// Lossy counting algorithm [Manku-Motwani2002]
// bucketSize = 1 / epsilon
//
// [Manku-Motwani2002]:
//   Gurmeet Singh Manku & Rajeev Motwani. (2002)
//   Approximate Frequency Counts over Data Streams.
//   VLDB'02, 346 - 357.
//   http://dl.acm.org/citation.cfm?id=1287400
type ApproximateCounter struct {
	counts           map[string]int
	errors           map[string]int
	epsilon          float64
	supportThreshold float64
	iBucket          int
	iItem            int
	bucketSize       int
	Counter
}

func (counter ApproximateCounter) getCountingResult() (map[string]int, int) {
	return counter.counts, 0
}

func (counter *ApproximateCounter) initialize() {
	counter.counts = make(map[string]int, counter.bucketSize)
	counter.errors = make(map[string]int, counter.bucketSize)
}

func (counter *ApproximateCounter) truncateNegligibleItems() {
	for k := range counter.counts {
		if counter.counts[k]+counter.errors[k] <= counter.iBucket {
			delete(counter.counts, k)
			delete(counter.errors, k)
		}
	}
	counter.iBucket++
}

func (counter *ApproximateCounter) truncateBySupport() {
	for k := range counter.counts {
		if float64(counter.counts[k]) < (counter.supportThreshold-counter.epsilon)*float64(counter.iItem) {
			delete(counter.counts, k)
			delete(counter.errors, k)
		}
	}
}

func (counter *ApproximateCounter) increment(item string) int {
	counter.iItem++

	_, itemExists := counter.counts[item]
	if itemExists {
		counter.counts[item]++
	} else {
		counter.counts[item] = 1
		counter.errors[item] = counter.iBucket - 1
	}

	if counter.iItem%counter.bucketSize == 0 {
		counter.truncateNegligibleItems()
	}

	return counter.counts[item]
}

func (counter ApproximateCounter) get(item string) int {
	c, exists := counter.counts[item]
	if exists {
		return c + counter.errors[item]
	}
	return 0
}

func (counter ApproximateCounter) toJSON() string {
	for k, v := range counter.counts {
		counter.counts[k] = v + counter.errors[k]
	}
	s, _ := json.Marshal(counter.counts)
	return string(s)
}

// NewApproximateCounter constructs ApproximateCounter struct
func NewApproximateCounter(epsilon float64, supportThreshold float64) *ApproximateCounter {
	counter := &ApproximateCounter{}
	counter.iBucket = 1
	counter.iItem = 0
	counter.supportThreshold = supportThreshold
	counter.epsilon = epsilon
	counter.bucketSize = int(1.0 / epsilon)
	counter.initialize()
	return counter
}
