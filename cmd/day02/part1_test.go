package day02

import (
	"fmt"
	"testing"
	"strings"
)
var input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
// Result for part 1: 2

func TestPart1(t *testing.T) {
	input := strings.Split(input, "\n")
	result := part1(input)
	fmt.Println(result)
}
