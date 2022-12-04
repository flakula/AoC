package main

import (
	"fmt"
	"io/ioutil"
	// "math"
	"strconv"
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
		assignments := strings.Split(inputdata[i], ",")
		first_assignment := strings.Split(assignments[0], "-")
		second_assignment := strings.Split(assignments[1], "-")

		x_1, _ := strconv.ParseFloat(first_assignment[0], 64)
		y_1, _ := strconv.ParseFloat(first_assignment[1], 64)

		x_2, _ := strconv.ParseFloat(second_assignment[0], 64)
		y_2, _ := strconv.ParseFloat(second_assignment[1], 64)

		if x_1 >= x_2 && x_1 <= y_2 || y_1 >= x_2 && y_1 <= y_2 {
			// 1 is inside 2
			total_sum += 1
			fmt.Printf("1 overlaps 2: %5v %3v %3v\n", total_sum, first_assignment, second_assignment)
		} else if x_2 >= x_1 && x_2 <= y_1 || y_2 >= x_1 && y_2 <= y_1 {
			// 2 is inside 1
			total_sum += 1
			fmt.Printf("2 overlaps 1: %5v %3v %3v\n", total_sum, first_assignment, second_assignment)
		}

	}
	fmt.Println(total_sum)
}
