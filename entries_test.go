package main

import "testing"

func TestEntryListLen(t *testing.T) {
	entryList := []Entry{
		{"b", 20},
		{"a", 10},
		{"c", 5},
	}
	if len(entryList) != 3 {
		t.Errorf("Unexpected length: len(%v) != 3", entryList)
	}
}

func TestExtractTopnItemsWhenTopnIsEqualToTheLength(t *testing.T) {
	m := map[string]int{"a": 10, "b": 20, "c": 5}
	actual := extractTopnItems(m, 3)
	expected := []Entry{
		{"b", 20},
		{"a", 10},
		{"c", 5},
	}
	if !equalEntryList(actual, expected) {
		t.Errorf("Unexpected return value: %v != %v", actual, expected)
	}
}

func TestExtractTopnItemsWhenTopnIsSmallerThanTheLength(t *testing.T) {
	m := map[string]int{"a": 10, "b": 20, "c": 5}
	actual := extractTopnItems(m, 2)
	expected := []Entry{
		{"b", 20},
		{"a", 10},
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
