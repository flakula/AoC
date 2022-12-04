package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	filebuffer, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		return
	}
	inputdata := strings.Split(string(filebuffer), "\r\n\r\n")
	fmt.Println(inputdata[0])
	max_sum := 0
	for i := 0; i < len(inputdata); i++ {
		elfs := strings.Split(inputdata[i], "\r\n")
		sum := 0
		for j := 0; j < len(elfs); j++ {
			tenBaseSixteenBitInt, _ := strconv.Atoi(elfs[j])
			sum = sum + tenBaseSixteenBitInt
		}
		// fmt.Println(i, sum, elfs)
		if max_sum < sum {
			max_sum = sum
		}
	}
	fmt.Println(max_sum)
}
