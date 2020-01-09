package main

import "testing"

func TestDoNotExist(t *testing.T) {
	counter := NewMapCounter()
	got := counter.get("do_not_exist")
	if got != 0 {
		t.Errorf("counter.get(do_not_exist) != 0")
	}
}

func TestExist(t *testing.T) {
	counter := NewMapCounter()

	counter.increment("bob")
	counter.increment("bob")
	counter.increment("bob")

	got := counter.get("bob")
	if got != 3 {
		t.Errorf("counter.get(bob) != 3")
	}
}

func TestToJson(t *testing.T) {
	counter := NewMapCounter()

	counter.increment("alice")
	counter.increment("bob")
	counter.increment("bob")

	got := counter.toJSON()
	if got != "[{\"count\":2,\"item\":\"bob\"},{\"count\":1,\"item\":\"alice\"}]" {
		t.Errorf("Unexpected json: %v", got)
	}
}
