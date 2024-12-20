package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	grid := InputToStringGrid2D()

	var start, end Point2D
	grid.ForEachPoint(func(p Point2D, s string) {
		switch s {
		case "S":
			start = p
		case "E":
			end = p
		}
	})

	path, _ := BreadthFirstSearch(
		start,
		func(p Point2D) []Point2D {
			var children []Point2D
			grid.ForEachOrthogonalNeighborPoint(p, func(q Point2D, s string) {
				if s != "#" {
					children = append(children, q)
				}
			})
			return children
		},
		func(p Point2D) bool {
			return p == end
		},
	)

	var count int
	for i := 0; i < len(path); i++ {
		for j := i + 1; j < len(path); j++ {
			dist := path[i].ManhattanDistance(path[j])
			if dist > 20 {
				continue
			}

			savings := j - i - dist
			if savings >= 100 {
				count++
			}
		}
	}
	fmt.Println(count)
}
