package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
)

func main() {
	grid := InputToStringGrid2D()

	var guard Turtle
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "^" {
			guard.Location = p
			guard.Heading = Up
		}
	})

	// Figure out the guard's initial path without any additional obstacles.
	// Since these are the only locations the guard will visit, they are the only
	// places where we need to consider adding an obstacle.
	candidates, _ := Walk(grid, guard)

	var ps Set[Point2D]
	for p := range candidates {
		if grid.GetPoint(p) == "." && IsLoop(grid, guard, p) {
			ps.Add(p)
		}
	}
	fmt.Println(len(ps))
}

func Walk(grid Grid2D[string], guard Turtle) (Set[Point2D], bool) {
	var seen Set[Turtle]
	for {
		if !seen.Add(guard) {
			return nil, true
		}

		mark := guard
		guard.Forward(1)
		if !grid.InBoundsPoint(guard.Location) {
			break
		}

		if s := grid.GetPoint(guard.Location); s != "." && s != "^" {
			guard = mark
			guard.TurnRight()
		}
	}

	var pts Set[Point2D]
	for t := range seen {
		pts.Add(t.Location)
	}
	return pts, false
}

func IsLoop(grid Grid2D[string], guard Turtle, obstacle Point2D) bool {
	grid.SetPoint(obstacle, "#")
	defer grid.SetPoint(obstacle, ".")

	_, loop := Walk(grid, guard)
	return loop
}
