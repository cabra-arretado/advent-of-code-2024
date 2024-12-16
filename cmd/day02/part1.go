package day02

import (
	"fmt"
	"strconv"
	"strings"
)

const MAX_DIFF = 3

func part1(file []string) int {
	fmt.Println("Day 02")
	total := 0
	for _, line := range file {
		if ProcessLine(line) {
			total += 1
		}
	}
	return total
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func passDiff(number1 string, number2 string, leftToRight bool) bool {
	num1, err := strconv.Atoi(number1)
	if err != nil {
		panic(err)
	}
	num2, err := strconv.Atoi(number2)
	if err != nil {
		panic(err)
	}
	if leftToRight {
		if num1 >= num2 {
			return false
		}
	} else {
		if num1 <= num2 {
			return false
		}
	}
	return abs(num1-num2) <= MAX_DIFF
}

func ProcessLine(line string) bool {
	line = strings.TrimSpace(line)
	numbers := strings.Split(line, " ")

	num0, err := strconv.Atoi(numbers[0])
	if err != nil {
		panic(err)
	}
	num1, err := strconv.Atoi(numbers[1])
	if err != nil {
		panic(err)
	}
	if num0 == num1 {
		return false
	}
	leftToRight := num0 < num1

	for i, _ := range numbers {
		if i == 0 {
			continue
		}
		if !passDiff(numbers[i-1], numbers[i], leftToRight) {
			return false
		}
	}
	return true
}
