package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	banks := []string{}
	for scanner.Scan() {
		banks = append(banks, scanner.Text())
	}
	if scanner.Err() != nil {
		panic("IDK WHATS HAPPENING")
	}

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
