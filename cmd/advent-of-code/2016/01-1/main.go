package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	var turtle lib.Turtle
	for _, direction := range InputToDirections() {
		switch direction.Turn {
		case 'L':
			turtle.TurnLeft()
		case 'R':
			turtle.TurnRight()
		}
		turtle.Forward(direction.Steps)
	}

	fmt.Println(lib.Origin2D.ManhattanDistance(turtle.Location))
}

type Direction struct {
	Turn  byte
	Steps int
}

func InputToDirections() []Direction {
	input := lib.InputToString()
	input = strings.ReplaceAll(input, ",", "")

	var directions []Direction
	for _, part := range strings.Fields(input) {
		directions = append(directions, Direction{
			Turn:  part[0],
			Steps: lib.ParseInt(part[1:]),
		})
	}

	return directions
}
