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

	token, str := parser.scanParser()
	fmt.Println(token, str)
}
