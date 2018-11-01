package main

import (
	"sort"
)

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

type PIFEntry struct {
	Tok int
	Id  int
}

var constantTableCount = 0
var varTableCount = 0
var symbolTableConstants []SymTableEntry
var symbolTableVariables []SymTableEntry
var PIF []PIFEntry

func storeToken(tok Token, lit string) {
	updateSymTables(tok, lit)
	if tok == IDENTIFIER {
		_, pos := idInSymTable(symbolTableVariables, lit)
		PIF = append(PIF, PIFEntry{int(tok), pos})
	} else if tok == CONSTANT || tok == TRUE || tok == FALSE {
		_, pos := idInSymTable(symbolTableConstants, lit)
		PIF = append(PIF, PIFEntry{int(tok), pos})
	} else {
		PIF = append(PIF, PIFEntry{int(tok), -1})
	}
}

func updateSymTables(tok Token, lit string) {
	if tok == IDENTIFIER {
		exists, _ := idInSymTable(symbolTableVariables, lit)
		if !exists {
			symbolTableVariables = insertSort(symbolTableVariables, SymTableEntry{lit, varTableCount})
			varTableCount++
		}
	} else if tok == CONSTANT || tok == TRUE || tok == FALSE {
		exists, _ := idInSymTable(symbolTableConstants, lit)
		if !exists {
			symbolTableConstants = insertSort(symbolTableConstants, SymTableEntry{lit, constantTableCount})
			constantTableCount++
		}
	}
}
