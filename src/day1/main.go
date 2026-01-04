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
	count := 0
	for _, line := range rotations {
		dir := line[0:1]
		distance, _ := strconv.Atoi(line[1:])

		switch dir {
		case "R":
			start = (((start + distance) % 100) + 100) % 100
		case "L":
			start = (((start - distance) % 100) + 100) % 100
		default:
			panic("NOOOOOOOOO!!!")
		}
		if start == 0 {
			count++
		}
	}
	fmt.Println("Part 1: What's the actual password to open the door?", count)
}
