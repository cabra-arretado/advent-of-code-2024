package day01

import (
	"fmt"
	"advent-of-code-2024/internal/utils"
)

func Day01() {
	file := utils.ReadFileAsSlice("01")
	fmt.Println("Day 01")
	fmt.Println("Part 1:", part1(file))
	fmt.Println("Part 2:", part2(file))
}
