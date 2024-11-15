package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	grid := puz.InputToStringGrid2D(2017, 19)

	var steps int
	turtle := puz.Turtle{Location: FindStart(grid), Heading: puz.Down}
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

func CanMoveForward(g puz.Grid2D[string], t puz.Turtle) bool {
	next := t.Location.Move(t.Heading)
	return g.InBoundsPoint(next) && g.GetPoint(next) != Empty
}

func FindStart(g puz.Grid2D[string]) puz.Point2D {
	for x := 0; x < g.Width; x++ {
		if g.Get(x, 0) != Empty {
			return puz.Point2D{X: x, Y: 0}
		}
	}
	return puz.Point2D{}
}

const Empty string = " "
