package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	n := aoc.InputToInt(2017, 3)
	c := spiral(n)

	fmt.Printf("distance: %d\n", c.ManhattanDistance(aoc.Point2D{}))
}

// Determine the coordinate of the nth value in a spiral.
func spiral(n int) aoc.Point2D {
	var coord aoc.Point2D

	value := 1
	dir := "E"
	dist := 1
	left := 1

	for value < n {
		value++

		if left == 0 {
			if dir == "E" {
				dir = "N"
				left = dist
			} else if dir == "N" {
				dir = "W"
				dist++
				left = dist
			} else if dir == "W" {
				dir = "S"
				left = dist
			} else {
				dir = "E"
				dist++
				left = dist
			}
		}

		switch dir {
		case "E":
			coord = coord.East()
		case "N":
			coord = coord.North()
		case "W":
			coord = coord.West()
		case "S":
			coord = coord.South()
		}

		left--
	}

	return coord
}
