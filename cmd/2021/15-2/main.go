package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	tile := InputToCave()

	cave := aoc.NewGrid2D[int](5*tile.Width, 5*tile.Height)
	for y := 0; y < cave.Height; y++ {
		for x := 0; x < cave.Width; x++ {
			tx, ty := x%tile.Width, y%tile.Height
			offset := x/tile.Width + y/tile.Height

			value := tile.GetXY(tx, ty) + offset
			for value > 9 {
				value -= 9
			}
			cave.AddXY(x, y, value)
		}
	}

	start, end := aoc.Origin2D, aoc.Point2D{X: cave.Width - 1, Y: cave.Height - 1}

	children := func(p aoc.Point2D) []aoc.Point2D {
		var children []aoc.Point2D
		cave.ForEachOrthogonalNeighbor(p, func(q aoc.Point2D, _ int) {
			children = append(children, q)
		})
		return children
	}
	goal := func(p aoc.Point2D) bool { return p == end }
	cost := func(from, to aoc.Point2D) int { return cave.Get(to) }
	heuristic := func(p aoc.Point2D) int { return end.ManhattanDistance(p) }

	_, risk, _ := aoc.AStarSearch(start, children, goal, cost, heuristic)
	fmt.Println(risk)
}

func InputToCave() aoc.Grid2D[int] {
	lines := aoc.InputToLines(2021, 15)

	cave := aoc.NewGrid2D[int](len(lines[0]), len(lines))
	for y, line := range lines {
		for x, c := range line {
			cave.AddXY(x, y, int(c-'0'))
		}
	}
	return cave
}
