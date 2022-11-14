package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	m := InputToHeightMap()

	isLowPoint := func(p aoc.Point2D) bool {
		hp := m.Get(p)

		isLow := true
		m.ForEachOrthogonalNeighbor(p, func(n aoc.Point2D, hn int) {
			isLow = isLow && (hp < hn)
		})
		return isLow
	}

	var risk int
	m.ForEach(func(p aoc.Point2D, height int) {
		if isLowPoint(p) {
			risk += height + 1
		}
	})
	fmt.Println(risk)
}

func InputToHeightMap() aoc.Grid2D[int] {
	lines := aoc.InputToLines(2021, 9)

	grid := aoc.NewGrid2D[int](len(lines[0]), len(lines))
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			grid.AddXY(x, y, aoc.ParseInt(string(lines[y][x])))
		}
	}

	return grid
}
