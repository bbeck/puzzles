package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, line := range aoc.InputToLines(2015, 5) {
		if IsNice(line) {
			count++
		}
	}

	fmt.Printf("count: %d\n", count)
}

func IsNice(s string) bool {
	return ContainsRepeatPair(s) && ContainsSplitRepeat(s)
}

func ContainsRepeatPair(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		substr := s[i : i+2]
		if strings.Contains(s[i+2:], substr) {
			return true
		}
	}

	return false
}

func ContainsSplitRepeat(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}

	return false
}
