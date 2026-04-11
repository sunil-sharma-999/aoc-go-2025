package main

import (
	"fmt"
	"strconv"

	"github.com/sunil-sharma-999/aoc-go-2025/src/utils"
)

func outputJoltage(bank string, joltageLen int) int {
	maxIndex := -1
	var max rune
	var numStr string
	for cell := range joltageLen {
		s := maxIndex + 1
		e := len(bank) - (joltageLen - cell) + 1
		for index, char := range bank[s:e] {
			if char > max {
				max = char
				maxIndex = s + index
			}
		}
		numStr = numStr + string(max)
		max = 0
	}
	num, _ := strconv.Atoi(numStr)

	return num
}

func main() {
	banks := utils.ReadLines()
	defer utils.Track()()

	part1 := 0
	part2 := 0

	for _, bank := range banks {
		part1 += outputJoltage(bank, 2)
		part2 += outputJoltage(bank, 12)
	}

	fmt.Println("what is the total output joltage?", part1)
	fmt.Println("What is the new total output joltage?", part2)

}
