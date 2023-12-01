package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	// "math"
)

func main() {
	filebuffer, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		return
	}
	inputdata := strings.Split(string(filebuffer), "\r\n")

	total_sum := 0
	for i, value := range inputdata {
		pos := i
		if value[0] == '$' {
			fmt.Println(pos, "COMMAND", value)
			/* code */
		} else if value[0] == 'd' {
			fmt.Println(pos, "directory", value)
		} else {
			fmt.Println(pos, "file", value)
		}
		// fmt.Println(pos, value)
	}
	fmt.Println(total_sum)
}
