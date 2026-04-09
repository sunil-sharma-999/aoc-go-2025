package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Track() func() {
	fmt.Println()
	start := time.Now()
	return func() {
		fmt.Printf("\n--- %dms ---\n", time.Since(start).Milliseconds())
	}
}

func ReadLines() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
		panic("IDK WHATS HAPPENING")
	}
	return lines
}

func ReadAllText() string {
	return strings.TrimRight(strings.Join(ReadLines(), "\n"), "\n")
}

func ReadSections() [][]string {
	lines := ReadLines()
	var sections [][]string
	var current []string
	for _, line := range lines {
		if line == "" {
			if len(current) > 0 {
				sections = append(sections, current)
				current = nil
			}
		} else {
			current = append(current, line)
		}
	}
	if len(current) > 0 {
		sections = append(sections, current)
	}
	return sections
}

func ParseInts(strs []string) []int {
	nums := make([]int, len(strs))
	for i, s := range strs {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			panic("ParseInts: cannot parse \"" + s + "\": " + err.Error())
		}
		nums[i] = n
	}
	return nums
}

func SplitInts(s string, sep string) []int {
	return ParseInts(strings.Split(s, sep))
}

func ReadGrid() [][]byte {
	lines := ReadLines()
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	return grid
}
