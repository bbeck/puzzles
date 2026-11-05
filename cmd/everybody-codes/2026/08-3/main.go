package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

const MISSING = -1000

func main() {
	grid := in.ToGrid2D(func(x, y int, s string) int {
		if s == "." {
			return MISSING
		}
		return ParseInt(s)
	})

	var warriors, energy int
	for i := 0; warriors < 100000; i++ {
		dir := DIRECTIONS[i%len(DIRECTIONS)]

		// Determine which are in shadows
		var inShadow Set[Point2D]
		for _, current := range dir.Edges(grid) {
			shadow := grid.GetPoint(current)

			for x, y, dist := current.X+dir.dx, current.Y+dir.dy, 1; grid.InBounds(x, y); x, y, dist = x+dir.dx, y+dir.dy, dist+1 {
				height := grid.Get(x, y)
				if height != MISSING && shadow-dist >= height {
					inShadow.Add(Point2D{X: x, Y: y})
				} else if height != MISSING && shadow-dist < height {
					shadow, dist = height, 0
				}
			}
		}

		// Grow everything not in a shadow
		grid.ForEachPoint(func(p Point2D, height int) {
			if height == MISSING || (height > 0 && inShadow.Contains(p)) {
				return
			}

			if height > 0 {
				energy++
			}

			if height == 9 {
				grid.SetPoint(p, -2)
				warriors++
			} else {
				grid.SetPoint(p, height+1)
			}
		})
	}
	fmt.Println(energy)
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
	var edges []Point2D
	grid.ForEachPoint(func(p Point2D, value int) {
		if value == MISSING {
			return
		}
		var q = Point2D{X: p.X - d.dx, Y: p.Y - d.dy}
		if !grid.InBoundsPoint(q) || grid.GetPoint(q) == MISSING {
			edges = append(edges, p)
		}
	})
	return edges
}
