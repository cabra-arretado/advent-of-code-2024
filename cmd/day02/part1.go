package day02

import (
	"fmt"
	"strings"
)

func part1(file []string) int {
	fmt.Println("Day 02")
	total := 0
	for _, line := range file {
		if processLine(line) {
			total += 1
		}
	}
	return total
}

func processLine(line string) bool {
	line = strings.TrimSpace(line)
	numbers := strings.Split(line, " ")
	fmt.Println(numbers)
}
