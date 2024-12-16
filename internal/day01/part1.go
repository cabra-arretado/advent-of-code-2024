package day01

import (
	"sort"
	"strconv"
	"strings"
)

func part1(input []string) int {
	leftSlice := make([]int, 0)
	rightSlice := make([]int, 0)
	for _, line := range input {
		splitedLine := strings.Split(line, "   ")
		leftNumber, err := strconv.Atoi(splitedLine[0])
		if err != nil {
			panic(err)
		}
		rightNumber, err := strconv.Atoi(splitedLine[1])
		if err != nil {
			panic(err)
		}
		leftSlice = append(leftSlice, leftNumber)
		rightSlice = append(rightSlice, rightNumber)
	}
	sort.Ints(leftSlice)
	sort.Ints(rightSlice)

	total := 0

	for i, leftNumber := range leftSlice {
		rightNumber := rightSlice[i]
		if rightNumber > leftNumber {
			total += rightNumber - leftNumber
		} else {
			total += leftNumber - rightNumber
		}
	}

	return total
}
