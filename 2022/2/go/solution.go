package main

import (
	"fmt"
	"io/ioutil"
	// "strconv"
	"strings"
)

func main() {
	filebuffer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return
	}
	inputdata := strings.Split(string(filebuffer), "\r\n")
	fmt.Println(inputdata[0])
	total_sum := 0
	for i := 0; i < len(inputdata); i++ {
		plays := strings.Split(inputdata[i], " ")
		enemy_play, my_play := plays[0], plays[1]
		round_points := 0

		// did I win ?
		if (enemy_play == "A" && my_play == "Y") || (enemy_play == "B" && my_play == "Z") || (enemy_play == "C" && my_play == "X") {
			round_points += 6
		}
		// it was a draw ?

		if (enemy_play == "A" && my_play == "X") || (enemy_play == "B" && my_play == "Y") || (enemy_play == "C" && my_play == "Z") {
			round_points += 3
		}
		// play points
		if my_play == "X" {
			round_points += 1
		} else if my_play == "Y" {
			round_points += 2
		} else {
			round_points += 3
		}
		total_sum += round_points
		fmt.Println(enemy_play, my_play, round_points)
	}
	fmt.Println(total_sum)
}
