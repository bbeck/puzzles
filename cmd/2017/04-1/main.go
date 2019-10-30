package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, line := range aoc.InputToLines(2017, 4) {
		if IsValid(line) {
			count++
		}
	}

	fmt.Printf("count: %d\n", count)
}

func IsValid(line string) bool {
	seen := make(map[string]bool)

	for _, word := range strings.Split(line, " ") {
		if seen[word] {
			return false
		}

		seen[word] = true
	}

	return true
}
