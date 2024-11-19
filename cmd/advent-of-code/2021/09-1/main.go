package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	m := lib.InputToIntGrid2D()

	isLowPoint := func(p lib.Point2D) bool {
		hp := m.GetPoint(p)

		isLow := true
		m.ForEachOrthogonalNeighbor(p, func(_ lib.Point2D, hn int) {
			isLow = isLow && (hp < hn)
		})
		return isLow
	}

	var risk int
	m.ForEachPoint(func(p lib.Point2D, height int) {
		if isLowPoint(p) {
			risk += height + 1
		}
	})
	fmt.Println(risk)
}
