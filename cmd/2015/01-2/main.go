package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	position := 0
	floor := 0
	for p, c := range aoc.InputToString(2015, 1) {
		position = p + 1

		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 {
			break
		}
	}

	fmt.Println(position)
}
