package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var start, volcano Point2D
	grid := in.ToGrid2D(func(x, y int, s string) int {
		if s == "@" {
			volcano = Point2D{X: x, Y: y}
			return 0
		}
		if s == "S" {
			start = Point2D{X: x, Y: y}
			return 0
		}
		return ParseInt(s)
	})

	// We're going to try different radii away from the volcano.  Using the radius
	// we'll choose the bottom most point below the volcano and then determine
	// the cost of traveling from this bottom most point to the start via the left
	// and right sides of the volcano.  If we can accomplish the traversal before
	// the lava spreads then we have a candidate solution.
	var candidates []int
	for r := range Min(grid.Width, grid.Height) / 2 {
		bottom := Point2D{X: volcano.X, Y: volcano.Y + r + 1}
		cost, ok := Cost(grid, bottom, start, volcano, r)
		if !ok || cost >= 30*(r+1) {
			continue
		}
		candidates = append(candidates, r*cost)
	}
	fmt.Println(Min(candidates...))
}

func Cost(grid Grid2D[int], start, end, volcano Point2D, r int) (int, bool) {
	children := func(dir string) func(Point2D) []Point2D {
		return func(p Point2D) []Point2D {
			var children []Point2D
			grid.ForEachOrthogonalNeighborPoint(p, func(q Point2D, _ int) {
				// Ensure that when we reach the y-level of the volcan we're on the
				// correct side of it.
				if q.Y == volcano.Y && dir == "left" && q.X > volcano.X {
					return
				}
				if q.Y == volcano.Y && dir == "right" && q.X < volcano.X {
					return
				}

				// Ensure that we're outside the lava.
				if dx, dy := q.X-volcano.X, q.Y-volcano.Y; dx*dx+dy*dy <= r*r {
					return
				}

				children = append(children, q)
			})
			return children
		}
	}

	cost := func(_, p Point2D) int { return grid.GetPoint(p) }

	// Perform a Dijkstra's shortest path search in both directions.  If either
	// search fails then we don't have a solution.
	costsL, pathsL := Dijkstra(start, children("left"), cost)
	if pathsL[end] == nil {
		return -1, false
	}
	costsR, pathsR := Dijkstra(start, children("right"), cost)
	if pathsR[end] == nil {
		return -1, false
	}

	total := grid.GetPoint(start) + costsL[end] + costsR[end]
	return total, true
}
