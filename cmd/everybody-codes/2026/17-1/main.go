package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToStringGrid2D()

	children := func(p Point2D) []Point2D {
		var children []Point2D
		grid.ForEachOrthogonalNeighborPoint(p, func(q Point2D, s string) {
			if s == "#" {
				children = append(children, q)
			}
		})
		return children
	}

	goal := func(p Point2D) bool {
		return p.X == grid.Width-1 && p.Y == grid.Height-1
	}

	steps, _ := BreadthFirstSearch(Origin2D, children, goal)
	fmt.Println(len(steps) - 1)
}
