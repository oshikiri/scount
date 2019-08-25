package main

import "encoding/json"

// ApproximateCounter is a counter using streaming counting algorithm
//
// One path KSP algorithm [KSP2003]
// countThreshold = 1/theta
//
// [KSP2003] Richard M. Karp, Scott Shenker, and Christos H. Papadimitriou. 2003.
//   A simple algorithm for finding frequent elements in streams and bags.
//   ACM Trans. Database Syst. 28, 1 (March 2003), 51-55.
//   DOI=http://dx.doi.org/10.1145/762471.762473
type ApproximateCounter struct {
	counts         map[string]int
	countThreshold float64
	nSubtracted    int
	nPassed        int
	Counter
}

func (counter *ApproximateCounter) getCountingResult() (map[string]int, int) {
	return counter.counts, counter.nSubtracted
}

func (counter *ApproximateCounter) initialize() {
	counter.nPassed = 0
	counter.nSubtracted = 0
	counter.counts = map[string]int{}
}

func (counter ApproximateCounter) increment(item string) int {
	counter.nPassed++

	// addition step
	v, ok := counter.counts[item]
	if ok {
		counter.counts[item] = v + 1
	} else {
		counter.counts[item] = 1
	}

	// deletion step
	if float64(len(counter.counts)) > counter.countThreshold {
		counter.nSubtracted++
		for k, v := range counter.counts {
			if v > 1 {
				counter.counts[k] = v - 1
			} else {
				delete(counter.counts, k)
			}
		}
	}

	return counter.counts[item] + counter.nSubtracted
}

func (counter ApproximateCounter) get(item string) int {
	return counter.counts[item] + counter.nSubtracted // WARN: is it appropriate?
}

func (counter ApproximateCounter) toJSON() string {
	for k, v := range counter.counts {
		counter.counts[k] = v + counter.nSubtracted // WARN: is it appropriate?
	}
	s, _ := json.Marshal(counter.counts)
	return string(s)
}

// NewApproximateCounter constructs ApproximateCounter struct
func NewApproximateCounter(countThreshold float64) *ApproximateCounter {
	counter := &ApproximateCounter{}
	counter.countThreshold = countThreshold
	counter.initialize()
	return counter
}
