package main

import (
	"fmt"
	"os"
)

func fileRead() {
	reader, err := os.Open("sample")
	if err != nil {
		panic(err)
	}
	parser := getParser(reader)
	token, str := parser.scanWithoutSpaces()
	for token != EOF {
		fmt.Println(token, str)
		storeToken(token, str)
		token, str = parser.scanWithoutSpaces()
	}
	fmt.Println(symbolTableVariables)
	fmt.Println(symbolTableConstants)
	fmt.Println(PIF)
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %+v\n", r)
		}
	}()

	fileRead()
}
