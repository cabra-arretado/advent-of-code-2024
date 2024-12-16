package day01

import (
	"strconv"
	"strings"
)

func part2(input []string) int {
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
	mapOfNumbers := make(map[int]int)
	for _, rightNumber := range rightSlice {
		_, ok := mapOfNumbers[rightNumber]
		if ok {
			mapOfNumbers[rightNumber]++
		} else {
			mapOfNumbers[rightNumber] = 1
		}
	}

	total := 0

	for _, leftNumber := range leftSlice {
		_, ok := mapOfNumbers[leftNumber]
		if ok {
			total += leftNumber * mapOfNumbers[leftNumber]
		}
	}

	return total
}

