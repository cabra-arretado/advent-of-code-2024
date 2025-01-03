package day03

import (
	"testing"
)

var testInput = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func TestPart1(t *testing.T) {
	t.Run("Part 1 - processLine", func(t *testing.T) {
		got := processLine(testInput)
		if got != 161 {
			t.Errorf("Expected 161, got %v", got)
		}
	})

	t.Run("Part 1 - processMul", func(t *testing.T) {
		got := processMul("mul(2,4)")
		if got != 8 {
			t.Errorf("Expected 8, got %v", got)
		}
	})

	t.Run("Part 1 - part1", func(t *testing.T) {
		got := part1([]string{testInput})
		if got != 161 {
			t.Errorf("Expected 161, got %v", got)
		}
	})
}
