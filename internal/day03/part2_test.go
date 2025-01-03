package day03

import (
	"testing"
)

func TestPart2(t *testing.T) {
	t.Run("Part 2 - processLine2", func(t *testing.T) {
		got := part2([]string{testInput})
		if got != 48 {
			t.Errorf("Got %d, expected 48", got)
		}
	})
}
