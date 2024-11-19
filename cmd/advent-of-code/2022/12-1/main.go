package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	var start, end lib.Point2D
	grid := lib.InputToGrid2D(func(x, y int, s string) byte {
		if s == "S" {
			start = lib.Point2D{X: x, Y: y}
			return 'a'
		}
		if s == "E" {
			end = lib.Point2D{X: x, Y: y}
			return 'z'
		}
		return s[0]
	})

	children := func(p lib.Point2D) []lib.Point2D {
		pv := grid.GetPoint(p)

		var children []lib.Point2D
		grid.ForEachOrthogonalNeighbor(p, func(child lib.Point2D, cv byte) {
			if cv <= pv+1 {
				children = append(children, child)
			}
		})
		return children
	}

	goal := func(p lib.Point2D) bool { return p == end }

	path, _ := lib.BreadthFirstSearch(start, children, goal)
	fmt.Println(len(path) - 1) // the path includes the starting node
}
