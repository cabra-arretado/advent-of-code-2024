package day03

import (
	"regexp"
	"strconv"
)

func part1(input []string) int {
	total := 0
	for _, line := range input {
		total += processLine(line)
	}
	return total
}

func processMul(mulExpression string) int {
	pattern := "\\d+"
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(mulExpression, -1)
	// Pick the first number
	firstNum, err := strconv.Atoi(matches[0][0])
	if err != nil {
		panic(err)
	}
	// Pick the second number
	secondNum, err := strconv.Atoi(matches[1][0])
	if err != nil {
		panic(err)
	}
	return firstNum * secondNum
}

func processLine(input string) int {
	pattern := "mul\\(\\d+,\\d+\\)"
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(input, -1)
	total := 0
	for _, v := range matches {
		total += processMul(v[0])
	}
	return total
}
