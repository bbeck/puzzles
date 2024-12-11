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
			grid.SetPoint(p, ".")
		}
	})

	// Figure out the guard's initial path without any additional obstacles.
	// Since these are the only locations the guard will visit, they are the only
	// places where we need to consider adding an obstacle.
	path, _ := Walk(grid, guard)

	var tried, looped Set[Point2D]
	for i := 1; i < len(path); i++ {
		// Don't place an obstacle at the same location twice.  Since the same
		// location came up again it means that the path already went through this
		// location.  Thus, it must remain open so that we could even get here the
		// second time.
		if !tried.Add(path[i].Location) {
			continue
		}

		grid.SetPoint(path[i].Location, "#")
		_, isLoop := Walk(grid, path[i-1])
		grid.SetPoint(path[i].Location, ".")

		if isLoop {
			looped.Add(path[i].Location)
		}
	}

	fmt.Println(len(looped))
}

func Walk(grid Grid2D[string], guard Turtle) ([]Turtle, bool) {
	var path []Turtle
	var seen Set[Turtle]
	for {
		if !grid.InBoundsPoint(guard.Location) {
			return path, false
		}

		if !seen.Add(guard) {
			return nil, true
		}

		for {
			_, s, _ := PeekForward(grid, guard)
			if s == "#" {
				guard.TurnRight()
				continue
			}

			break
		}

		path = append(path, guard)
		guard.Forward(1)
	}
}

func PeekForward(grid Grid2D[string], t Turtle) (Point2D, string, bool) {
	t.Forward(1)

	if !grid.InBoundsPoint(t.Location) {
		return Origin2D, "", false
	}
	return t.Location, grid.GetPoint(t.Location), true
}
