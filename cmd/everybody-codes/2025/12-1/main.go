package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToGrid2D(func(_, _ int, s string) int { return ParseInt(s) })

	var seen Set[Point2D]
	var queue = []Point2D{Origin2D}
	for len(queue) > 0 {
		var p Point2D

		p, queue = queue[0], queue[1:]
		if !seen.Add(p) {
			continue
		}

		pv := grid.GetPoint(p)
		grid.ForEachOrthogonalNeighborPoint(p, func(q Point2D, qv int) {
			if qv <= pv {
				queue = append(queue, q)
			}
		})
	}

	fmt.Println(len(seen))
}
