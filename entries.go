package main

import "encoding/json"

// Entry is key-value pair to sort
type Entry struct {
	key   string
	value int
}

// MarshalJSON converts Entry to JSON element
func (entry Entry) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"item": entry.key, "count": entry.value})
}

// EntryList for sorting
type EntryList []Entry

func (l EntryList) swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func extractTopnItems(m map[string]int, topn int) EntryList {
	list := EntryList{}
	for k, v := range m {
		entry := Entry{k, v}
		list = append(list, entry)
	}

	// https://en.wikipedia.org/wiki/Selection_algorithm#Partial_selection_sort
	for i, e := range list {
		if i >= topn {
			return list[:topn]
		}
		iMax := i
		eMax := e
		for j := i; j < len(list); j++ {
			if list[j].value > eMax.value {
				iMax = j
				eMax = list[j]
			}
		}
		list.swap(i, iMax)
	}

	return list[:topn]
}
