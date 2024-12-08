package main

import (
	// "advent-of-code-2024/utils"
	"advent-of-code-2024/cmd/day01"
	"advent-of-code-2024/cmd/day02"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the day: (e.g. 01, 10, 20) ")
	input, _ := reader.ReadString('\n') // Read input until newline
	input = strings.TrimSpace(input)
	switch input {
	case "01":
		day01.Day01()
	case "02":
		day02.Day02()
	default:
		fmt.Println("Day not found")
	}
}
