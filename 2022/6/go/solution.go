package main

import (
	"fmt"
	"io/ioutil"
	// "math"
)

func main() {
	filebuffer, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		return
	}
	inputdata := string(filebuffer)

	array_copy := []rune{'0', '0', '0', '0'}

	pos := 0

	for i, value := range inputdata {
		pos = i
		fmt.Printf("%c", value)
		array_copy[i%4] = value
		if i < 3 {
			continue
		}
		if array_copy[0] != array_copy[1] && array_copy[1] != array_copy[2] && array_copy[0] != array_copy[2] && array_copy[1] != array_copy[3] && array_copy[0] != array_copy[3] && array_copy[2] != array_copy[3] {
			break
		}

	}
	fmt.Println()
	fmt.Printf(
		"%d %c %c %c %c",
		pos+1,
		array_copy[0], array_copy[1], array_copy[2], array_copy[3],
	)
}
