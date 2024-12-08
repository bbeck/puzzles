package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
)

func main() {
	grid := InputToStringGrid2D()

	m := make(map[string][]Point2D)
	grid.ForEachPoint(func(p Point2D, s string) {
		if s != "." {
			m[s] = append(m[s], p)
		}
	})

	var seen Set[Point2D]
	for _, ps := range m {
		for i, a := range ps {
			for _, b := range ps[i+1:] {
				dx, dy := b.X-a.X, b.Y-a.Y

				p1 := Point2D{X: a.X - dx, Y: a.Y - dy}
				if grid.InBoundsPoint(p1) {
					seen.Add(p1)
				}

				p2 := Point2D{X: a.X + 2*dx, Y: a.Y + 2*dy}
				if grid.InBoundsPoint(p2) {
					seen.Add(p2)
				}
			}
		}
	}
	fmt.Println(len(seen))
}
