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

		my_plays_points := []int{0, 3, 6} // win draw loose

		if enemy_play == "A" { // rock
			my_plays_points[0] += 3 // scissors
			my_plays_points[1] += 1 // rock
			my_plays_points[2] += 2 // paper
		} else if enemy_play == "B" { // paper
			my_plays_points[0] += 1 // rock
			my_plays_points[1] += 2 // paper
			my_plays_points[2] += 3 // scissors
		} else if enemy_play == "C" { // scissors
			my_plays_points[0] += 2 // paper
			my_plays_points[1] += 3 // scissors
			my_plays_points[2] += 1 // rock
		}

		if my_play == "X" { //win
			round_points = my_plays_points[0]
		} else if my_play == "Y" { // draw
			round_points = my_plays_points[1]
		} else if my_play == "Z" { //loose
			round_points = my_plays_points[2]
		}

		total_sum += round_points
		fmt.Println(enemy_play, my_play, round_points, total_sum)
	}
	fmt.Println(total_sum)
}
