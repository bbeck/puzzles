package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	m := aoc.InputToIntGrid2D(2021, 9)

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
