package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToGrid2D(func(_, _ int, s string) string {
		return s
	})

	var start Point2D
	for x := 0; x < grid.Width; x++ {
		if grid.Get(x, 0) == "." {
			start = Point2D{X: x}
			break
		}
	}

	children := func(p Point2D) []Point2D {
		var children []Point2D
		grid.ForEachOrthogonalNeighborPoint(p, func(q Point2D, s string) {
			if s != "#" {
				children = append(children, q)
			}
		})
		return children
	}

	goal := func(p Point2D) bool {
		s := grid.GetPoint(p)
		return s != "." && s != "#"
	}

	path, ok := BreadthFirstSearch(start, children, goal)
	if !ok {
		return
	}

	fmt.Println(2 * (len(path) - 1))
}
