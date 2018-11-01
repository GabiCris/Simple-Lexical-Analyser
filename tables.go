package main

import "sort"

func insertSort(data []SymTableEntry, el SymTableEntry) []SymTableEntry {
	index := sort.Search(len(data), func(i int) bool { return data[i].Identifier > el.Identifier })
	data = append(data, SymTableEntry{})
	copy(data[index+1:], data[index:])
	data[index] = el
	return data
}

func idInSymTable(list []SymTableEntry, id string) (bool, int) {
	for _, b := range list {
		if b.Identifier == id {
			return true, b.Position
		}
	}
	return false, 0
}

type SymTableEntry struct {
	Identifier string
	Position   int
}

var symbolTableConstants []SymTableEntry
var symbolTableVariables []SymTableEntry
