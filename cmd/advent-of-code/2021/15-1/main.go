package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	cave := lib.InputToIntGrid2D()

	start, end := lib.Origin2D, lib.Point2D{X: cave.Width - 1, Y: cave.Height - 1}

	children := func(p lib.Point2D) []lib.Point2D {
		var children []lib.Point2D
		cave.ForEachOrthogonalNeighbor(p, func(q lib.Point2D, _ int) {
			children = append(children, q)
		})
		return children
	}
	goal := func(p lib.Point2D) bool { return p == end }
	cost := func(from, to lib.Point2D) int { return cave.GetPoint(to) }
	heuristic := func(p lib.Point2D) int { return end.ManhattanDistance(p) }

	_, risk, _ := lib.AStarSearch(start, children, goal, cost, heuristic)
	fmt.Println(risk)
}
