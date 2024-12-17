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
}

