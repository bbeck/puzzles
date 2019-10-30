package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	location := aoc.Point2D{0, 0}
	seen := map[aoc.Point2D]int{
		location: 1,
	}

	for _, b := range aoc.InputToBytes(2015, 3) {
		switch b {
		case '^':
			location = location.North()
		case '<':
			location = location.West()
		case '>':
			location = location.East()
		case 'v':
			location = location.South()
		default:
			log.Fatalf("unrecognized location: %s", string(b))
		}

		seen[location]++
	}

	fmt.Printf("number of locations: %d\n", len(seen))
}
