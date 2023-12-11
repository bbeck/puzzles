package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := aoc.InputToGrid2D(2023, 11, func(x int, y int, s string) string {
		return s
	})
	grid = Expand(grid)

	var seen aoc.Set[aoc.Point2D]
	var sum int
	grid.ForEachPoint(func(p aoc.Point2D, s string) {
		if s != "#" {
			return
		}

		for _, q := range seen.Entries() {
			sum += Distance(grid, p, q)
		}
		seen.Add(p)
	})

	fmt.Println(sum)
}

func Expand(g aoc.Grid2D[string]) aoc.Grid2D[string] {
	for x := 0; x < g.Width; x++ {
		var cells aoc.Set[string]
		for y := 0; y < g.Height; y++ {
			cells.Add(g.Get(x, y))
		}

		if !cells.Contains("#") {
			for y := 0; y < g.Height; y++ {
				g.Set(x, y, "G")
			}
		}
	}

	for y := 0; y < g.Height; y++ {
		var cells aoc.Set[string]
		for x := 0; x < g.Width; x++ {
			cells.Add(g.Get(x, y))
		}

		if !cells.Contains("#") {
			for x := 0; x < g.Width; x++ {
				g.Set(x, y, "G")
			}
		}
	}

	return g
}

var Distances = map[string]int{
	".": 1,
	"#": 1,
	"G": 2,
}

func Distance(grid aoc.Grid2D[string], p, q aoc.Point2D) int {
	var d int
	for {
		switch {
		case p.X < q.X:
			q = q.Left()
		case p.X > q.X:
			q = q.Right()
		case p.Y < q.Y:
			q = q.Up()
		case p.Y > q.Y:
			q = q.Down()
		default:
			return d
		}

		d += Distances[grid.GetPoint(q)]
	}
}
