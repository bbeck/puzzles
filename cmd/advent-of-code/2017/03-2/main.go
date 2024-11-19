package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	target := lib.InputToInt()
	sums := map[lib.Point2D]int{
		lib.Origin2D: 1,
	}

	var sum int
	for n := 0; ; n++ {
		c := SpiralCoordinate(n)

		sum = sums[c]
		if sum == 0 {
			for _, neighbor := range c.Neighbors() {
				sum += sums[neighbor]
			}
			sums[c] = sum
		}

		if sum > target {
			break
		}
	}

	fmt.Println(sum)
}

func SpiralCoordinate(n int) lib.Point2D {
	var turtle lib.Turtle
	turtle.TurnRight()

	// Edges represent the distance along the edges that we're traveling.
	// We use a container since we have to configure multiple edges at a time.
	var edges lib.Stack[int]
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
