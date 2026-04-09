package main

import (
	"fmt"
	"strconv"

	"github.com/sunil-sharma-999/aoc-go-2025/src/utils"
)

func main() {
	banks := utils.ReadLines()

	part1 := 0

	for _, bank := range banks {
		maxIndex := -1
		var max rune
		for i, char := range bank[0 : len(bank)-1] {
			if char > max {
				max = char
				maxIndex = i
			}
		}
		var max2nd rune
		for _, char := range bank[maxIndex+1:] {
			if char > max2nd {
				max2nd = char
			}
		}
		numStr := string(max) + string(max2nd)
		num, _ := strconv.Atoi(numStr)
		part1 += num
	}

	fmt.Println("What do you get if you add up all of the invalid IDs?", part1)

}
