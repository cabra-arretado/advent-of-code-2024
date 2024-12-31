package day03

import (
	"advent-of-code-2024/internal/utils"
	"fmt"
)

func Day03() {
	input := utils.ReadFileAsSlice("03")

	fmt.Println("Day 03")
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
