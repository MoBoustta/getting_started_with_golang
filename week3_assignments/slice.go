package main

import (
	"fmt"
	"sort"
)

func main() {
	input := 0
	slc := make([]int, 3)
	count := 0
	for true {
		fmt.Println("Please enter a number of 'X' to quit")
		_, _ = fmt.Scan(&input)
		if input == 0 { break }

		if count < 3 {
			for i, v := range slc { if v == 0 { slc[i] = input; break } }
		} else { slc = append(slc, input) }
		sort.Ints(slc)
		fmt.Println(slc)
		count++
	}
}