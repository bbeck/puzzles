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

var Deltas = map[aoc.Heading]aoc.Point2D{
	aoc.Up:    {X: 0, Y: -1},
	aoc.Right: {X: 1, Y: 0},
	aoc.Down:  {X: 0, Y: 1},
	aoc.Left:  {X: -1, Y: 0},
}

func CanMoveForward(g aoc.Grid2D[string], t aoc.Turtle) bool {
	delta := Deltas[t.Heading]
	next := aoc.Point2D{X: t.Location.X + delta.X, Y: t.Location.Y + delta.Y}
	return g.InBounds(next) && g.Get(next) != Empty
}

func FindStart(g aoc.Grid2D[string]) aoc.Point2D {
	for x := 0; x < g.Width; x++ {
		if g.GetXY(x, 0) != Empty {
			return aoc.Point2D{X: x, Y: 0}
		}
	}
	return aoc.Point2D{}
}

const Empty string = " "

func InputToGrid() aoc.Grid2D[string] {
	lines := aoc.InputToLines(2017, 19)

	grid := aoc.NewGrid2D[string](len(lines[0]), len(lines))
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			grid.AddXY(x, y, string(lines[y][x]))
		}
	}

	return grid
}
