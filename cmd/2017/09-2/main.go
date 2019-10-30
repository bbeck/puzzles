package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var length int
	for _, line := range aoc.InputToLines(2017, 9) {
		length += GarbageLength(line)
	}

	fmt.Printf("length: %d\n", length)
}

func GarbageLength(s string) int {
	var depth int      // depth of the group we're currently in
	var inGarbage bool // whether or not we're currently within garbage
	var length int     // the length of the garbage encountered

	for i := 0; i < len(s); i++ {
		switch {
		case !inGarbage && s[i] == '{':
			depth++
		case !inGarbage && s[i] == '}':
			depth--
		case !inGarbage && s[i] == '<':
			inGarbage = true
		case inGarbage && s[i] == '>':
			inGarbage = false
		case inGarbage && s[i] == '!':
			i++
		case inGarbage:
			length++
		}
	}

	return length
}
