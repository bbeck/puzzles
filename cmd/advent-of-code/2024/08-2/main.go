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

				for n := -grid.Width; n < grid.Width; n++ {
					p := Point2D{X: a.X + n*dx, Y: a.Y + n*dy}
					if grid.InBoundsPoint(p) {
						seen.Add(p)
					}
				}
			}
		}
	}
	fmt.Println(len(seen))
}
