package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

const MISSING = -1

func main() {
	original := in.ToGrid2D(func(x, y int, s string) int {
		if s == "." {
			return MISSING
		}
		return ParseInt(s)
	})

	var sum int
	for _, dir := range DIRECTIONS {
		grid := Copy(original)
		for _, current := range dir.Edges(grid) {
			shadow := grid.GetPoint(current)

			for x, y, dist := current.X+dir.dx, current.Y+dir.dy, 1; grid.InBounds(x, y); x, y, dist = x+dir.dx, y+dir.dy, dist+1 {
				height := grid.Get(x, y)
				if height != MISSING && shadow-dist >= height {
					grid.Set(x, y, MISSING)
				} else if height != MISSING && shadow-dist < height {
					shadow = height
					dist = 0
				}
			}
		}

		sum += Energy(grid)
	}
	fmt.Println(sum)
}

func Energy(grid Grid2D[int]) int {
	var energy int
	grid.ForEachPoint(func(_ Point2D, v int) {
		if v != MISSING {
			energy++
		}
	})
	return energy
}

func Copy[T any](grid Grid2D[T]) Grid2D[T] {
	var cells = make([]T, len(grid.Cells))
	copy(cells, grid.Cells)
	return Grid2D[T]{Cells: cells, Height: grid.Height, Width: grid.Width}
}

var DIRECTIONS = []Direction{
	{dx: 1, dy: 1},
	{dx: -1, dy: 1},
	{dx: -2, dy: 0},
	{dx: -1, dy: -1},
	{dx: 1, dy: -1},
	{dx: 2, dy: 0},
}

type Direction struct {
	dx, dy int
}

func (d Direction) Edges(grid Grid2D[int]) []Point2D {
	var edges Set[Point2D]
	grid.ForEachPoint(func(p Point2D, value int) {
		if value == MISSING {
			return
		}
		var q = Point2D{X: p.X - d.dx, Y: p.Y - d.dy}
		if !grid.InBoundsPoint(q) || grid.GetPoint(q) == MISSING {
			edges.Add(p)
		}
	})
	return edges.Entries()
}
