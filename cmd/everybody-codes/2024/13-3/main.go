package main

import (
	"fmt"
	"strconv"

	. "github.com/bbeck/puzzles/lib"
)

const WALL = 1e10

func main() {
	grid, starts, end := InputToGrid()

	children := func(n Point2D) []Point2D {
		var children []Point2D
		grid.ForEachOrthogonalNeighborPoint(n, func(c Point2D, height int) {
			if height != WALL {
				children = append(children, c)
			}
		})
		return children
	}

	// Instead of searching from each start to the end, search from the end to
	// any of the starts.
	goal := func(n Point2D) bool {
		return starts.Contains(n)
	}

	cost := func(from, to Point2D) int {
		a, b := grid.GetPoint(from), grid.GetPoint(to)
		a, b = Min(a, b), Max(a, b)

		// Consider distance going down from b to a as well as wrapping around from
		// a to b.
		return Min(b-a+1, 10+a-b+1)
	}

	heuristic := func(n Point2D) int {
		return Min(
			n.X,             // Distance to left edge
			n.Y,             // Distance to top edge
			grid.Width-n.X,  // Distance to right edge
			grid.Height-n.Y, // Distance to bottom edge
		)
	}

	_, c, _ := AStarSearch(end, children, goal, cost, heuristic)
	fmt.Println(c)
}

func InputToGrid() (Grid2D[int], Set[Point2D], Point2D) {
	temp := InputToStringGrid2D()

	var starts Set[Point2D]
	var end Point2D
	grid := NewGrid2D[int](temp.Width, temp.Height)
	temp.ForEachPoint(func(p Point2D, s string) {
		switch s {
		case "S":
			starts.Add(p)
		case "E":
			end = p
		case "#":
			grid.SetPoint(p, WALL)
		default:
			if n, err := strconv.Atoi(s); err == nil {
				grid.SetPoint(p, n)
			}
		}
	})

	return grid, starts, end
}
