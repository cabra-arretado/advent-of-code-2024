package day03

import (
	"testing"
)

var testInput2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestPart2(t *testing.T) {
	t.Run("Part 2 - processLine2", func(t *testing.T) {
		got := part2([]string{testInput2})
		if got != 48 {
			t.Errorf("Got %d, expected 48", got)
		}
	})
}
