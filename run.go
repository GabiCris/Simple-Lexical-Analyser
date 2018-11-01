package main

import (
	"fmt"
	"os"
)

func main() {
	reader, err := os.Open("sample")
	if err != nil {
		panic(err)
	}
	parser := getParser(reader)

	el := SymTableEntry{"hey", 1}
	symbolTableConstants = insertSort(symbolTableConstants, el)
	token, str := parser.scanWithoutSpaces()
	for token != EOF {
		fmt.Println(token, str)
		token, str = parser.scanWithoutSpaces()
	}

}
