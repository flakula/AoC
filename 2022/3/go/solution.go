package main

import (
	"fmt"
	"io/ioutil"
	// "strconv"
	"strings"
)

func main() {
	filebuffer, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		return
	}
	inputdata := strings.Split(string(filebuffer), "\r\n")

	total_sum := 0

	for i := 0; i < len(inputdata); i++ {
		first_compartment := inputdata[i][:len(inputdata[i])/2]
		second_compartment := inputdata[i][len(inputdata[i])/2:]
		m := make(map[rune]rune)
		for _, c := range first_compartment {
			var value rune
			if c >= 96 {
				value = c - (96)
			} else {
				value = c - 65 + 27
			}
			m[c] = value
		}
		for _, c := range second_compartment {
			value, in := m[c]
			if in {
				total_sum += int(value)
				fmt.Printf("char=%1c rune=%3v priority=%2v ", c, c, value)
				break
			}
		}
		fmt.Printf("%5v %25v %25v %v\n", total_sum, first_compartment, second_compartment, m)
	}
	fmt.Println(total_sum)
}
