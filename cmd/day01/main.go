package day01

import (
	"fmt"
	"advent-of-code-2024/utils"
)

func Day01() {
	file := utils.ReadFileAsSlice("01")
	fmt.Println("Day 01")
	fmt.Println("Part 1:", part1(file))
}
