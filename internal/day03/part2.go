package day03

import (
	"regexp"
)

func part2(input []string) int {
	total := 0
	elabled := true
	for _, line := range input {
		total += processLine2(line, &elabled)
	}
	return total
}

func processLine2(input string, enabled *bool) int {
	pattern := `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(input, -1)
	total := 0
	for _, v := range matches {
		if v[0] == "do()" {
			*enabled = true
		} else if v[0] == "don't()" {
			*enabled = false
		} else {
			if *enabled {
				total += processMul(v[0])
			}
		}
	}
	return total
}
