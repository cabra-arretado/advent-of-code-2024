package day01

import (
	"advent-of-code-2024/utils"
	"strings"
	"testing"
)

var input = `3   4
4   3
2   5
1   3
3   9
3   3`

// Expected part 1 = 11

func TestPart1(t *testing.T) {
	trimmed := strings.Trim(input, "")
	splitedInput := strings.Split(trimmed, "\n")
	result := part1(splitedInput)
	utils.PassesTest(result == 11)

}
