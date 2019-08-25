package main

import "testing"

func TestApproximateCounterDoNotExist(t *testing.T) {
	counter := NewApproximateCounter(0.5)
	got := counter.get("do_not_exist")
	if got != 0 {
		t.Errorf("counter.get(do_not_exist) != 0")
	}
}

func TestApproximateCounterExist(t *testing.T) {
	counter := NewApproximateCounter(2)

	counter.increment("bob")
	counter.increment("bob")
	counter.increment("bob")

	got := counter.get("bob")
	if got != 3 {
		t.Errorf("counter.get(bob) != 3")
	}
}

func assertBobCount(t *testing.T, expected int, actual int) {
	if actual != expected {
		t.Errorf("counter.get(bob) != %v, got %v", expected, actual)
	}
}
func TestApproximateCounterExistWhenOverflow(t *testing.T) {
	counter := NewApproximateCounter(2)

	for i := 0; i < 100; i++ {
		counter.increment("bob")
	}

	counter.increment("alice1") // added
	assertBobCount(t, 100, counter.get("bob"))
	counter.increment("alice2") // added and then deleted
	assertBobCount(t, 99, counter.get("bob"))
	counter.increment("alice3") // added
	assertBobCount(t, 99, counter.get("bob"))
}

func TestApproximateCounterToJson(t *testing.T) {
	counter := NewApproximateCounter(2)

	counter.increment("alice")
	counter.increment("bob")
	counter.increment("bob")

	got := counter.toJSON()
	if got != "{\"alice\":1,\"bob\":2}" {
		t.Errorf("Unexpected json: %v", got)
	}
}
