package main

import (
	"fmt"
	"strconv"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

const WALL = 1e10

func main() {
	grid, start, end := InputToGrid()

	children := func(n Point2D) []Point2D {
		var children []Point2D
		grid.ForEachOrthogonalNeighborPoint(n, func(c Point2D, height int) {
			if height != WALL {
				children = append(children, c)
			}
		})
		return children
	}

	goal := func(n Point2D) bool {
		return n == end
	}

	cost := func(from, to Point2D) int {
		a, b := grid.GetPoint(from), grid.GetPoint(to)
		a, b = Min(a, b), Max(a, b)

		// Consider distance going down from b to a as well as wrapping around from
		// a to b.
		return Min(b-a+1, 10+a-b+1)
	}

	heuristic := func(n Point2D) int {
		return n.ManhattanDistance(end)
	}

	_, c, _ := AStarSearch(start, children, goal, cost, heuristic)
	fmt.Println(c)
}

func InputToGrid() (Grid2D[int], Point2D, Point2D) {
	temp := in.ToGrid2D(func(_, _ int, s string) string {
		return s
	})

	var start, end Point2D
	grid := NewGrid2D[int](temp.Width, temp.Height)
	temp.ForEachPoint(func(p Point2D, s string) {
		switch s {
		case "S":
			start = p
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

	return grid, start, end
}
