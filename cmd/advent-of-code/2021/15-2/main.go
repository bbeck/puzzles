package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	tile := puz.InputToIntGrid2D()

	cave := puz.NewGrid2D[int](5*tile.Width, 5*tile.Height)
	for y := 0; y < cave.Height; y++ {
		for x := 0; x < cave.Width; x++ {
			tx, ty := x%tile.Width, y%tile.Height
			offset := x/tile.Width + y/tile.Height

			value := tile.Get(tx, ty) + offset
			for value > 9 {
				value -= 9
			}
			cave.Set(x, y, value)
		}
	}

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
