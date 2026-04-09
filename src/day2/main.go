package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sunil-sharma-999/aoc-go-2025/src/utils"
)

func repeatedNum(num int) int {
	str := strconv.Itoa(num)
	sLen := len(str)
	limit := sLen / 2
	for div := 1; div <= limit; div++ {
		if sLen%div == 0 && strings.Repeat(str[0:div], sLen/div) == str {
			return num
		}
	}
	return 0
}

func main() {
	line := utils.ReadAllText()

	invalidIDRanges := [][2]int{}
	for str := range strings.SplitSeq(line, ",") {
		invalidIDR := [2]int{}
		splitStrArr := strings.Split(str, "-")
		invalidIDR[0], _ = strconv.Atoi(splitStrArr[0])
		invalidIDR[1], _ = strconv.Atoi(splitStrArr[1])
		invalidIDRanges = append(invalidIDRanges, invalidIDR)
	}

	part1 := 0

	for _, row := range invalidIDRanges {
		for num := row[0]; num <= row[1]; num++ {
			numStr := strconv.Itoa(num)
			len := len(numStr)
			midIndex := len / 2
			if numStr[0:midIndex] == numStr[midIndex:] {
				part1 += num
			}
		}
	}

	part2 := 0
	for _, row := range invalidIDRanges {
		for num := row[0]; num <= row[1]; num++ {
			part2 += repeatedNum(num)
		}
	}

	fmt.Println("What do you get if you add up all of the invalid IDs?", part1)
	fmt.Println("What do you get if you add up all of the invalid IDs using these new rules?", part2)
}
