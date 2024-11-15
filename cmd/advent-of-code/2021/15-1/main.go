package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	cave := puz.InputToIntGrid2D(2021, 15)

	start, end := puz.Origin2D, puz.Point2D{X: cave.Width - 1, Y: cave.Height - 1}

	children := func(p puz.Point2D) []puz.Point2D {
		var children []puz.Point2D
		cave.ForEachOrthogonalNeighbor(p, func(q puz.Point2D, _ int) {
			children = append(children, q)
		})
		return children
	}
	goal := func(p puz.Point2D) bool { return p == end }
	cost := func(from, to puz.Point2D) int { return cave.GetPoint(to) }
	heuristic := func(p puz.Point2D) int { return end.ManhattanDistance(p) }

	_, risk, _ := puz.AStarSearch(start, children, goal, cost, heuristic)
	fmt.Println(risk)
}
