package main

import "sort"

// Entry is key-value pair to sort
type Entry struct {
	key   string
	value int
}

// EntryList for sorting
type EntryList []Entry

func (l EntryList) Len() int {
	return len(l)
}

func (l EntryList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l EntryList) Less(i, j int) bool {
	if l[i].value == l[j].value {
		return (l[i].key < l[j].key)
	}
	return (l[i].value < l[j].value)
}

func sortMap(m map[string]int, topn int) EntryList {
	list := EntryList{}
	for k, v := range m {
		entry := Entry{k, v}
		list = append(list, entry)
	}

	sort.Sort(sort.Reverse(list))
	return list
}
