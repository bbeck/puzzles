package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	m := puz.InputToIntGrid2D(2021, 9)

	isLowPoint := func(p puz.Point2D) bool {
		hp := m.GetPoint(p)

		isLow := true
		m.ForEachOrthogonalNeighbor(p, func(_ puz.Point2D, hn int) {
			isLow = isLow && (hp < hn)
		})
		return isLow
	}

	var risk int
	m.ForEachPoint(func(p puz.Point2D, height int) {
		if isLowPoint(p) {
			risk += height + 1
		}
	})
	fmt.Println(risk)
}
