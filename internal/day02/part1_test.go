package day02

import (
	"fmt"
	"strings"
	"testing"
)

var input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

// Per line: T F F F F T
// Result for part 1: 2

func TestPart1(t *testing.T) {
	input := strings.Split(input, "\n")
	fmt.Println("Day 02")
	fmt.Println(ProcessLine(input[0]))
	t.Run("test Day 2", func(t *testing.T) {
		got := ProcessLine(input[0])
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Test Day 2", func(t *testing.T) {
		got := part1(input)
		want := 2
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
