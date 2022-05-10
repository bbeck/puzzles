package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	floor := 0
	for _, c := range aoc.InputToString(2015, 1) {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	fmt.Println(floor)
}
