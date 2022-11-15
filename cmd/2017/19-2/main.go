package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := InputToGrid()

	var steps int
	turtle := aoc.Turtle{Location: FindStart(grid), Heading: aoc.Down}
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

func CanMoveForward(g aoc.Grid2D[string], t aoc.Turtle) bool {
	next := t.Location.Move(t.Heading)
	return g.InBoundsPoint(next) && g.GetPoint(next) != Empty
}

func FindStart(g aoc.Grid2D[string]) aoc.Point2D {
	for x := 0; x < g.Width; x++ {
		if g.Get(x, 0) != Empty {
			return aoc.Point2D{X: x, Y: 0}
		}
	}
	return aoc.Point2D{}
}

const Empty string = " "

func InputToGrid() aoc.Grid2D[string] {
	return aoc.InputToGrid2D[string](2017, 19, func(x int, y int, s string) string {
		return s
	})
}
