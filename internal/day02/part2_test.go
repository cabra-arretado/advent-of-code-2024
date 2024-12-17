package day02

import (
	"strings"
	"testing"
)

// Per line: T F F T T T
// Result for part 2: 4

func TestPart2(t *testing.T) {
	input := strings.Split(input, "\n")
	t.Run("test Day 2", func(t *testing.T) {
		got := ProcessLine(input[3])
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Test Day 2 Part 2", func(t *testing.T) {
		got := part1(input)
		want := 4
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
