package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToGrid2D(func(x, y int, s string) string {
		return s
	})

	var steps int
	turtle := Turtle{Location: FindStart(grid), Heading: Down}
	for {
		steps++
		if CanMoveForward(grid, turtle) {
			turtle.Forward(1)
			continue
		}

		turtle.TurnLeft()
		if CanMoveForward(grid, turtle) {
			turtle.Forward(1)
			continue
		}

		turtle.TurnLeft()
		turtle.TurnLeft()
		if CanMoveForward(grid, turtle) {
			turtle.Forward(1)
			continue
		}

		// We're out of moves
		break
	}

	fmt.Println(steps)
}

func CanMoveForward(g Grid2D[string], t Turtle) bool {
	next := t.Location.Move(t.Heading)
	return g.InBoundsPoint(next) && g.GetPoint(next) != Empty
}

func FindStart(g Grid2D[string]) Point2D {
	for x := 0; x < g.Width; x++ {
		if g.Get(x, 0) != Empty {
			return Point2D{X: x, Y: 0}
		}
	}
	return Point2D{}
}

const Empty string = " "
