package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	n := aoc.InputToInt(2017, 3)
	fmt.Printf("value: %d\n", SpiralSum(n))
}

func SpiralSum(n int) int {
	var coord aoc.Point2D
	grid := map[aoc.Point2D]int{
		coord: 1,
	}

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

		// We know that value goes into coord.  Let's compute the grid value for
		// coord which is the sum of the neighboring values.
		neighbors := []aoc.Point2D{
			coord.North(),
			coord.North().West(),
			coord.West(),
			coord.West().South(),
			coord.South(),
			coord.South().East(),
			coord.East(),
			coord.East().North(),
		}

		var v int
		for _, neighbor := range neighbors {
			v += grid[neighbor]
		}

		if v > n {
			return v
		}

		grid[coord] = v
	}

	log.Fatal("couldn't find sum")
	return 0
}
