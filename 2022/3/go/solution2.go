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

	for i := 0; i < len(inputdata); i += 3 {
		first_rucksack := inputdata[i]
		second_rucksack := inputdata[i+1]
		third_rucksack := inputdata[i+2]
		m := make(map[rune]rune)
		for _, c := range first_rucksack + second_rucksack + third_rucksack {
			var value rune
			if c >= 96 {
				value = c - (96)
			} else {
				value = c - 65 + 27
			}
			m[c] = value
		}
		for k, v := range m {
			if strings.Contains(first_rucksack, string(k)) && strings.Contains(second_rucksack, string(k)) && strings.Contains(third_rucksack, string(k)) {
				fmt.Printf("char=%1c rune=%3v priority=%2v ", k, k, v)
				total_sum += int(v)
				break
			}
		}
		fmt.Printf("%5v %25v %25v %25v %v\n", total_sum, first_rucksack, second_rucksack, third_rucksack, m)
	}
	fmt.Println(total_sum)
}
