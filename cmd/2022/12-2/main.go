package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var start puz.Point2D
	grid := puz.InputToGrid2D(2022, 12, func(x, y int, s string) byte {
		if s == "S" {
			return 'a'
		}
		if s == "E" {
			start = puz.Point2D{X: x, Y: y}
			return 'z'
		}
		return s[0]
	})

	children := func(p puz.Point2D) []puz.Point2D {
		pv := grid.GetPoint(p)

		var children []puz.Point2D
		grid.ForEachOrthogonalNeighbor(p, func(child puz.Point2D, cv byte) {
			if cv >= pv-1 {
				children = append(children, child)
			}
		})
		return children
	}

	goal := func(p puz.Point2D) bool { return grid.GetPoint(p) == 'a' }

	path, _ := puz.BreadthFirstSearch(start, children, goal)
	fmt.Println(len(path) - 1)
}
