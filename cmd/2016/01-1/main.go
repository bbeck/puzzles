package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	var turtle aoc.Turtle
	for _, direction := range InputToDirections() {
		switch direction.Turn {
		case 'L':
			turtle.TurnLeft()
		case 'R':
			turtle.TurnRight()
		}
		turtle.Forward(direction.Steps)
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
