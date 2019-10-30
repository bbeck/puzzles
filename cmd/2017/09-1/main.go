package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var score int
	for _, line := range aoc.InputToLines(2017, 9) {
		score += Score(line)
	}

	fmt.Printf("score: %d\n", score)
}

func Score(s string) int {
	var score int      // overall score for the string
	var depth int      // depth of the group we're currently in
	var inGarbage bool // whether or not we're currently within garbage

	for i := 0; i < len(s); i++ {
		switch {
		case !inGarbage && s[i] == '{':
			depth++
			score += depth
		case !inGarbage && s[i] == '}':
			depth--
		case !inGarbage && s[i] == '<':
			inGarbage = true
		case inGarbage && s[i] == '>':
			inGarbage = false
		case inGarbage && s[i] == '!':
			i++
		}
	}

	return score
}
