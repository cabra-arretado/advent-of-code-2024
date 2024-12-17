package day02

import (
	"advent-of-code-2024/internal/utils"
	"fmt"
)

func Day02() {
	input := utils.ReadFileAsSlice("02")

	fmt.Println("Day 02")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
