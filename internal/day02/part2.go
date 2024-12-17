package day02

import (
	"strconv"
	"strings"
)

func part2(file []string) int {
	total := 0
	for _, line := range file {
		if ProcessLine2(line) {
			total += 1
		}
	}
	return total
}

func passDiff2(number1 int, number2 int, leftToRight bool) bool {
	if leftToRight {
		if number1 >= number2 {
			return false
		}
	} else {
		if number1 <= number2 {
			return false
		}
	}
	return abs(number1-number2) <= MAX_DIFF
}

func ProcessLine2(line string) bool{
	line = strings.TrimSpace(line)
	numbersStr := strings.Split(line, " ")
	slice := make([]int, len(numbersStr))
	for i, v := range numbersStr {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		slice[i] = num
	}
	if ProcessNumbers(slice) {
		return true
	}
	for i, _ := range slice {
		newSlice := append([]int(nil), slice...)
		newSlice = append(newSlice[:i], newSlice[i+1:]...)
		if ProcessNumbers(newSlice) {
			return true
		}
	}
	return false
}


func ProcessNumbers(numbers []int) bool {

	if numbers[0] == numbers[1] {
		return false
	}
	leftToRight := numbers[0] < numbers[1]

	for i, _ := range numbers {
		if i == 0 {
			continue
		}
		if !passDiff2(numbers[i-1], numbers[i], leftToRight) {
			return false
		}
	}
	return true
}
