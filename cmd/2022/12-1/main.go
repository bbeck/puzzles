package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var start, end aoc.Point2D
	grid := aoc.InputToGrid2D(2022, 12, func(x, y int, s string) byte {
		if s == "S" {
			start = aoc.Point2D{X: x, Y: y}
			return 'a'
		}
		if s == "E" {
			end = aoc.Point2D{X: x, Y: y}
			return 'z'
		}
		return s[0]
	})

	children := func(p aoc.Point2D) []aoc.Point2D {
		pv := grid.GetPoint(p)

		var children []aoc.Point2D
		grid.ForEachOrthogonalNeighbor(p, func(child aoc.Point2D, cv byte) {
			if cv <= pv+1 {
				children = append(children, child)
			}
		})
		return children
	}

	goal := func(p aoc.Point2D) bool { return p == end }

	path, _ := aoc.BreadthFirstSearch(start, children, goal)
	fmt.Println(len(path) - 1) // the path includes the starting node
}
