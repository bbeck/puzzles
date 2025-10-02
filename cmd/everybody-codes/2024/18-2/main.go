package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToGrid2D(func(_, _ int, s string) string { return s })

	var entrances []Point2D
	var plants []Point2D
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "P" {
			plants = append(plants, p)
		}
		if p.X == 0 || p.Y == 0 || p.X == grid.Width-1 || p.Y == grid.Height-1 {
			if s == "." {
				entrances = append(entrances, p)
			}
		}
	})

	var visitedAt = make(map[Point2D]int)
	for _, entrance := range entrances {
		for p, tm := range Flood(grid, entrance) {
			if tmOld, present := visitedAt[p]; !present || tm < tmOld {
				visitedAt[p] = tm
			}
		}
	}

	var tm int
	for _, p := range plants {
		tm = max(tm, visitedAt[p])
	}
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
