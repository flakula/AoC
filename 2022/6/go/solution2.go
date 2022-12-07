package main

import (
	"fmt"
	"io/ioutil"
	// "math"
)

func main() {
	filebuffer, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		return
	}
	inputdata := string(filebuffer)

	array_copy := []rune{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}
	m := make(map[rune]int)
	pos := 0

	for i, value := range inputdata {
		pos = i
		// fmt.Printf("%c", value)
		// si ya esta lo quitas
		v, is_in := m[array_copy[i%14]]
		if is_in {
			m[array_copy[i%14]] = v - 1
			if m[array_copy[i%14]] == 0 {
				delete(m, array_copy[i%14])
			}
		}

		array_copy[i%14] = value

		v, is_in = m[value]
		if is_in {
			m[array_copy[i%14]] = v + 1
		} else {
			m[value] = 1
		}

		fmt.Printf("%d %d ", len(m), pos+1)
		for i := 0; i < 14; i++ {
			fmt.Printf("%c ", array_copy[i])
		}
		fmt.Println(m)

		if len(m) == 14 {
			break
		}
	}
}
