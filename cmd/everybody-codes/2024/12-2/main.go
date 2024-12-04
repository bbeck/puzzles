package main

import (
	"fmt"
	"math"

	. "github.com/bbeck/puzzles/lib"
)

var Segments = map[string]int{"A": 1, "B": 2, "C": 3}
var Multipliers = map[string]int{"T": 1, "H": 2}

func main() {
	grid := InputToGrid()

	catapults := make(map[string]Point2D)
	grid.ForEachPoint(func(p Point2D, s string) {
		if Segments[s] > 0 {
			catapults[s] = p
		}
	})

	hits := NewGrid2D[int](grid.Width, grid.Height)
	visit := func(p Point2D, ranking int) {
		h := hits.GetPoint(p)
		if h == 0 || ranking < h {
			hits.SetPoint(p, ranking)
		}
	}

	for power := 1; power < grid.Width/2; power++ {
		for id, c := range catapults {
			p := c

			// 45-degree angle up
			for n := 0; n < power; n++ {
				p = Point2D{X: p.X + 1, Y: p.Y - 1}
				visit(p, Segments[id]*power)
			}
			// Horizontal
			for n := 0; n < power; n++ {
				p = Point2D{X: p.X + 1, Y: p.Y}
				visit(p, Segments[id]*power)
			}
			// 45-degree angle down
			for p.X < grid.Width-1 && p.Y < grid.Height-1 {
				p = Point2D{X: p.X + 1, Y: p.Y + 1}
				visit(p, Segments[id]*power)
			}
		}
	}

	var sum int
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "T" || s == "H" {
			sum += Multipliers[s] * hits.GetPoint(p)
		}
	})
	fmt.Println(sum)
}

func InputToGrid() Grid2D[string] {
	grid := InputToStringGrid2D()

	minCatapultY := math.MaxInt
	grid.ForEachPoint(func(p Point2D, s string) {
		if Segments[s] > 0 {
			minCatapultY = Min(minCatapultY, p.Y)
		}
	})

	// We're going to shoot with enough power to reach the right side of the
	// grid from the lowest catapult.  The means that we need enough vertical
	// height from the perspective of the highest catapult.
	offset := grid.Width + (grid.Height - minCatapultY)

	taller := NewGrid2D[string](grid.Width, grid.Height+offset)
	taller.Fill(".")

	grid.ForEach(func(x int, y int, s string) {
		taller.Set(x, y+offset, s)
	})

	return taller
}
