package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkRead(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("sample")
	checkRead(err)

	buffer := bufio.NewReader(file)
	value, err := buffer.ReadString(' ')
	checkRead(err)
	fmt.Println(value)

}
