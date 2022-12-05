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
	inputdata := strings.Split(string(filebuffer), "\r\n\r\n")

	stacks := strings.Split(inputdata[0], "\r\n")
	moves := strings.Split(inputdata[1], "\r\n")

	last_line := strings.Trim(stacks[len(stacks)-1], " ")
	stack_index := strings.Split(last_line, " ")
	total_stacks, _ := strconv.Atoi(stack_index[len(stack_index)-1])

	m := make(map[int][]string)

	for i := len(stacks) - 2; i >= 0; i-- {
		for j := 0; j < total_stacks; j++ {
			_, in := m[i+1]
			if !in {
				m[i+1] = []string{}
			}
			box := string(stacks[i][(j*4)+1])
			if box == " " {
				continue
			}
			m[j+1] = append(m[j+1], box)
		}
	}

	fmt.Println(m)

	for i := 0; i < len(moves); i++ {
		assignments := strings.Split(moves[i], " ")
		if len(assignments) < 5 {
			continue
		}
		amount, _ := strconv.Atoi(assignments[1])
		from, _ := strconv.Atoi(assignments[3])
		to, _ := strconv.Atoi(assignments[5])
		fmt.Println(amount, from, to)
		for i := 0; i < amount; i++ {
			box := m[from][len(m[from])-1]
			m[from] = m[from][:len(m[from])-1]
			m[to] = append(m[to], box)
			fmt.Println(m)
		}
	}

	for i := 1; i <= total_stacks; i++ {
		fmt.Printf(m[i][len(m[i])-1])
	}
	fmt.Println()
}
