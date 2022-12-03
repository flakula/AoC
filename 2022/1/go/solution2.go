package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, world")
	filebuffer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return
	}
	inputdata := strings.Split(string(filebuffer), "\r\n\r\n")
	fmt.Println(inputdata[0])
	max_sum := []int{0, 0, 0}
	for i := 0; i < len(inputdata); i++ {
		elfs := strings.Split(inputdata[i], "\r\n")
		sum := 0
		for j := 0; j < len(elfs); j++ {
			tenBaseSixteenBitInt, _ := strconv.Atoi(elfs[j])
			sum = sum + tenBaseSixteenBitInt
		}
		fmt.Println(i, sum, elfs)
		// fmt.Println(i, sum)
		if max_sum[2] < sum {
			max_sum[2] = sum
			if max_sum[2] > max_sum[1] {
				t := max_sum[2]
				max_sum[2] = max_sum[1]
				max_sum[1] = t
				if max_sum[1] > max_sum[0] {
					t := max_sum[1]
					max_sum[1] = max_sum[0]
					max_sum[0] = t
				}
			}
		}
	}
	fmt.Println(max_sum[0]+max_sum[1]+max_sum[2], max_sum)
}
