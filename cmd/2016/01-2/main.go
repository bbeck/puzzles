package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	var seen aoc.Set[aoc.Point2D]
	var turtle aoc.Turtle

outer:
	for _, direction := range InputToDirections() {
		switch direction.Turn {
		case 'L':
			turtle.TurnLeft()
		case 'R':
			turtle.TurnRight()
		}

		for i := 0; i < direction.Steps; i++ {
			turtle.Forward(1)
			if !seen.Add(turtle.Location) {
				break outer
			}
		}
	}

	fmt.Println(aoc.Origin2D.ManhattanDistance(turtle.Location))
}

type Direction struct {
	Turn  byte
	Steps int
}

func InputToDirections() []Direction {
	input := aoc.InputToString(2016, 1)
	input = strings.ReplaceAll(input, ",", "")

	var directions []Direction
	for _, part := range strings.Fields(input) {
		directions = append(directions, Direction{
			Turn:  part[0],
			Steps: aoc.ParseInt(part[1:]),
		})
	}

	return directions
}
