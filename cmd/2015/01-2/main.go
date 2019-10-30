package main

import (
	"fmt"
	"log"

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
		default:
			log.Fatalf("unrecognized character: %s", string(c))
		}

		if floor == -1 {
			break
		}
	}

	fmt.Printf("enters basement at position: %d\n", position)
}
