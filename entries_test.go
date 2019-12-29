package main

import "testing"

func TestEntryListLen(t *testing.T) {
	entryList := []Entry{
		Entry{"b", 20},
		Entry{"a", 10},
		Entry{"c", 5},
	}
	if len(entryList) != 3 {
		t.Errorf("Unexpected length: len(%v) != 3", entryList)
	}
}

func TestSortMap(t *testing.T) {
	m := map[string]int{"a": 10, "b": 20, "c": 5}
	actual := sortMap(m)
	expected := []Entry{
		Entry{"b", 20},
		Entry{"a", 10},
		Entry{"c", 5},
	}
	if !equalEntryList(actual, expected) {
		t.Errorf("Unexpected return value: %v != %v", actual, expected)
	}
}

func equalEntryList(x1, x2 EntryList) bool {
	n1 := len(x1)
	n2 := len(x2)
	if n1 != n2 {
		return false
	}

	for i := 0; i < n1; i++ {
		if x1[i] != x2[i] {
			return false
		}
	}

	return true
}
