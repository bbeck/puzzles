package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"math"
)

func main() {
	var starts []aoc.Point2D
	var end aoc.Point2D
	grid := aoc.InputToGrid2D(2022, 12, func(x, y int, s string) byte {
		if s == "S" {
			s = "a"
		}
		if s == "a" {
			starts = append(starts, aoc.Point2D{X: x, Y: y})
		}
		if s == "E" {
			end = aoc.Point2D{X: x, Y: y}
			s = "z"
		}
		return s[0]
	})

	children := func(p aoc.Point2D) []aoc.Point2D {
		pv := grid.GetPoint(p)

		var children []aoc.Point2D
		grid.ForEachOrthogonalNeighbor(p, func(child aoc.Point2D, cv byte) {
			if cv <= pv || cv-pv == 1 {
				children = append(children, child)
			}
		})
		return children
	}

	goal := func(p aoc.Point2D) bool { return p == end }

	best := math.MaxInt
	for _, start := range starts {
		if path, ok := aoc.BreadthFirstSearch(start, children, goal); ok {
			best = aoc.Min(best, len(path)-1)
		}
	}
	fmt.Println(best)
}
