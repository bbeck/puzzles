package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	n := in.Int()
	c := SpiralCoordinate(n)
	fmt.Println(Origin2D.ManhattanDistance(c))
}

func SpiralCoordinate(n int) Point2D {
	var turtle Turtle
	turtle.TurnRight()

	// Edges represent the distance along the edges that we're traveling.
	// We use a container since we have to configure multiple edges at a time.
	var edges Stack[int]
	edges.Push(1)
	edges.Push(1)

	// Remaining represents how much further along the current edge we need to
	// travel before making a turn.
	remaining := edges.Peek()

	for n > 1 {
		turtle.Forward(1)
		remaining--
		n--

		if remaining == 0 {
			// We've completed an edge, see if we need to prepare the next set of
			// edges.
			if edge := edges.Pop(); edges.Empty() {
				edges.Push(edge + 1)
				edges.Push(edge + 1)
			}

			remaining = edges.Peek()
			turtle.TurnLeft()
		}
	}

	return turtle.Location
}
