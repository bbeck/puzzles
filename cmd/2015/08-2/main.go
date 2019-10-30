package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var sum int
	for _, line := range aoc.InputToLines(2015, 8) {
		sum += EncodeLength(line) - len(line)
	}

	fmt.Printf("size: %d\n", sum)
}

func EncodeLength(s string) int {
	length := 2 // open and close "

	for i := 0; i < len(s); i++ {
		if s[i] == '\\' || s[i] == '"' {
			length += 2
			continue
		}

		length++
	}

	return length
}
