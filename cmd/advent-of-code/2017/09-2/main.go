package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	length := GarbageLength(puz.InputToString(2017, 9))
	fmt.Println(length)
}

func GarbageLength(s string) int {
	var length int     // the length of the garbage encountered
	var depth int      // depth of the group we're currently in
	var inGarbage bool // whether we're currently within garbage

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
