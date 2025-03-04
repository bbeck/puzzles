package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var seen Set[Point2D]
	var turtle Turtle

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

	fmt.Println(Origin2D.ManhattanDistance(turtle.Location))
}

type Direction struct {
	Turn  byte
	Steps int
}

func InputToDirections() []Direction {
	in.Remove(", ")

	var directions []Direction
	for in.HasNext() {
		directions = append(directions, Direction{
			Turn:  in.Byte(),
			Steps: in.Int(),
		})
	}
	return directions
}
