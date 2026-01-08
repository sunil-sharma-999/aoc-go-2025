package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	rotations := []string{}

	for scanner.Scan() {
		rotations = append(rotations, scanner.Text())
	}

	if scanner.Err() != nil {
		panic("error reading standard input")
	}
	start := 50
	part1C := 0
	part2C := 0
	for _, line := range rotations {
		dir := line[0:1]
		distance, _ := strconv.Atoi(line[1:])

		switch dir {
		case "L":
			for range distance {
				start = (start - 1) % 100
				if start == 0 {
					part2C++
				}
			}
		case "R":
			for range distance {
				start = (start + 1) % 100
				if start == 0 {
					part2C++
				}
			}
		}
		if start == 0 {
			part1C++
		}
	}
	fmt.Println("Part 1: What's the actual password to open the door?", part1C)
	fmt.Println("Part 2: what is the password to open the door?", part2C)
}
