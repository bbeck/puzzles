package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"strings"
)

func main() {
	grid := lib.InputToStringGrid2D()

	var visited strings.Builder
	turtle := lib.Turtle{Location: FindStart(grid), Heading: lib.Down}
	for {
		if c := grid.GetPoint(turtle.Location); c >= "A" && c <= "Z" {
			visited.WriteString(c)
		}

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

	fmt.Println(visited.String())
}

func CanMoveForward(g lib.Grid2D[string], t lib.Turtle) bool {
	next := t.Location.Move(t.Heading)
	return g.InBoundsPoint(next) && g.GetPoint(next) != Empty
}

func FindStart(g lib.Grid2D[string]) lib.Point2D {
	for x := 0; x < g.Width; x++ {
		if g.Get(x, 0) != Empty {
			return lib.Point2D{X: x, Y: 0}
		}
	}
	return lib.Point2D{}
}

const Empty string = " "
