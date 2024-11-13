package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	var seen puz.Set[puz.Point2D]
	var turtle puz.Turtle

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

	fmt.Println(puz.Origin2D.ManhattanDistance(turtle.Location))
}

type Direction struct {
	Turn  byte
	Steps int
}

func InputToDirections() []Direction {
	input := puz.InputToString(2016, 1)
	input = strings.ReplaceAll(input, ",", "")

	var directions []Direction
	for _, part := range strings.Fields(input) {
		directions = append(directions, Direction{
			Turn:  part[0],
			Steps: puz.ParseInt(part[1:]),
		})
	}

	return directions
}
