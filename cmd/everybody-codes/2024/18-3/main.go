package main

import (
	"fmt"
	"math"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToGrid2D(func(_, _ int, s string) string { return s })

	var plants []Point2D
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "P" {
			plants = append(plants, p)
		}
	})

	// Use each plant as the starting point for the flood fill and add together
	// the time each cell is visited.  The cell with the lowest total time is
	// the one that's closest to all plants.
	visitedAt := NewGrid2D[int](grid.Width, grid.Height)
	for _, entrance := range plants {
		for p, tm := range Flood(grid, entrance) {
			visitedAt.SetPoint(p, visitedAt.GetPoint(p)+tm)
		}
	}

	var tm = math.MaxInt
	visitedAt.ForEach(func(x int, y int, visited int) {
		if grid.Get(x, y) == "." {
			tm = min(tm, visited)
		}
	})
	fmt.Println(tm)
}

func Flood(grid Grid2D[string], start Point2D) map[Point2D]int {
	var visitedAt = map[Point2D]int{start: 0}
	children := func(p Point2D) []Point2D {
		var children []Point2D

		grid.ForEachOrthogonalNeighborPoint(p, func(q Point2D, s string) {
			if _, seen := visitedAt[q]; !seen && s != "#" {
				visitedAt[q] = visitedAt[p] + 1
				children = append(children, q)
			}
		})
		return children
	}
	BreadthFirstSearch(start, children, func(Point2D) bool { return false })

	return visitedAt
}
