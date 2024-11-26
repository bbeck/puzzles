package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
)

func main() {
	grid := InputToGrid2D(func(_ int, _ int, s string) int {
		if s == "#" {
			return 1
		}
		return 0
	})

	for {
		var changed bool
		grid = grid.MapPoint(func(p Point2D, v int) int {
			if v == 0 {
				return 0
			}

			for _, q := range p.OrthogonalNeighbors() {
				if grid.GetPoint(q) != v {
					return v
				}
			}

			changed = true
			return v + 1
		})

		if !changed {
			break
		}
	}

	var sum int
	grid.ForEach(func(_ int, _ int, v int) {
		sum += v
	})
	fmt.Println(sum)
}
