package main

import "testing"

func TestApproximateCounterDoNotExist(t *testing.T) {
	counter := NewApproximateCounter(0.1, 0.1)
	got := counter.get("do_not_exist")
	if got != 0 {
		t.Errorf("counter.get(do_not_exist) != 0")
	}
}

func TestApproximateCounterExist(t *testing.T) {
	counter := NewApproximateCounter(0.1, 0.1)

	counter.increment("bbb")
	counter.increment("bbb")
	counter.increment("bbb")

	got := counter.get("bbb")
	if got != 3 {
		t.Errorf("counter.get(bbb) != 3")
	}
}

func assertCount(t *testing.T, counter Counter, item string, expected int) {
	actual := counter.get(item)
	if actual != expected {
		t.Errorf("counter.get(%v) != %v, got %v", item, expected, actual)
	}
}
func TestApproximateCounterExistWhenOverflow(t *testing.T) {
	counter := NewApproximateCounter(0.1, 0.2)

	counter.increment("aaa") // iIncrement = 1

	assertCount(t, counter, "aaa", 1)

	for i := 2; i < 10; i++ {
		counter.increment("bbb")
	} // iIncrement = 9

	assertCount(t, counter, "aaa", 1)
	assertCount(t, counter, "bbb", 8)

	counter.increment("bbb") // iIncrement = 9

	assertCount(t, counter, "aaa", 0)
	assertCount(t, counter, "bbb", 9)

	counter.increment("aaa") // iIncrement = 10

	// count = 1, error = 1, count + error = 2
	assertCount(t, counter, "aaa", 2)

	counter.truncateBySupport()
	assertCount(t, counter, "aaa", 0)
}

func TestApproximateCounterToJson(t *testing.T) {
	counter := NewApproximateCounter(0.01, 0.01)

	counter.increment("aaa")
	counter.increment("bbb")
	counter.increment("bbb")

	got := counter.toJSON()
	if got != "[{\"count\":2,\"item\":\"bbb\"},{\"count\":1,\"item\":\"aaa\"}]" {
		t.Errorf("Unexpected json: %v", got)
	}
}
