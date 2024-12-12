package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	grid := InputToStringGrid2D()

	var ds DisjointSet[Point2D]
	grid.ForEachPoint(func(p Point2D, s1 string) {
		ds.Add(p)

		grid.ForEachOrthogonalNeighborPoint(p, func(q Point2D, s2 string) {
			ds.Add(q)
			if s1 == s2 {
				ds.Union(p, q)
			}
		})
	})

	regions := make(map[Point2D][]Point2D)
	grid.ForEachPoint(func(p Point2D, _ string) {
		f, _ := ds.Find(p)
		regions[f] = append(regions[f], p)
	})

	var total int
	for _, ps := range regions {
		var perimeter int
		for _, p := range ps {
			s, _ := ds.Find(p)
			for _, q := range p.OrthogonalNeighbors() {
				if t, found := ds.Find(q); !found || s != t {
					perimeter++
				}
			}
		}

		total += len(ps) * perimeter
	}
	fmt.Println(total)
}
