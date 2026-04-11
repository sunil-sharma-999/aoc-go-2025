package main

import (
	"fmt"

	"github.com/sunil-sharma-999/aoc-go-2025/src/utils"
)

const RoP = '@'

func canBeForkLifted(x, y, lineEnd, gridEnd int, grid [][]byte) bool {
	if grid[x][y] != RoP {
		return false
	}
	var count int8

	for n1 := range 3 {
		n1 -= 1
		for n2 := range 3 {
			n2 -= 1
			if n1 == 0 && n2 == 0 {
				continue
			}
			if x+n1 >= 0 && x+n1 <= lineEnd && y+n2 >= 0 && y+n2 <= gridEnd && grid[x+n1][y+n2] == RoP {
				count++
			}
		}
	}
	if count > 3 {
		return false
	}

	return true

}

func main() {
	grid := utils.ReadGrid()
	defer utils.Track()()

	part1 := 0
	part2 := 0

	gridEnd := len(grid) - 1
	lineEnd := len(grid[0]) - 1

	for x := range grid {
		for y := range grid[x] {
			if canBeForkLifted(x, y, lineEnd, gridEnd, grid) {
				part1++
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
