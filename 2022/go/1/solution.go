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
	max_sum := int64(0)
	for i := 0; i < len(inputdata); i++ {
		elfs := strings.Split(inputdata[i], "\r\n")
		sum := int64(0)
		for j := 0; j < len(elfs); j++ {
			tenBaseSixteenBitInt, _ := strconv.ParseInt(elfs[j], 10, 64)
			sum = sum + tenBaseSixteenBitInt
		}
		fmt.Println(i, sum, elfs)
		// fmt.Println(i, sum)
		if max_sum < sum {
			max_sum = sum
		}
	}
	fmt.Println(max_sum)
}
