package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	m := InputToHeightMap()

	isLowPoint := func(p aoc.Point2D) bool {
		hp := m.GetPoint(p)

		isLow := true
		m.ForEachOrthogonalNeighbor(p, func(_ aoc.Point2D, hn int) {
			isLow = isLow && (hp < hn)
		})
		return isLow
	}

	var risk int
	m.ForEachPoint(func(p aoc.Point2D, height int) {
		if isLowPoint(p) {
			risk += height + 1
		}
	})
	fmt.Println(risk)
}

func InputToHeightMap() aoc.Grid2D[int] {
	return aoc.InputToGrid2D(2021, 9, func(x int, y int, s string) int {
		return aoc.ParseInt(s)
	})
}
