package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
)

var Values = map[string]int{
	"#": 1,
}

func main() {
	grid := InputToGrid2D(func(_ int, _ int, s string) int {
		return Values[s]
	})

	for {
		var changed bool
		grid = grid.MapPoint(func(p Point2D, v int) int {
			if v == 0 {
				return 0
			}

			for _, q := range p.Neighbors() {
				if !grid.InBoundsPoint(q) || grid.GetPoint(q) != v {
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
