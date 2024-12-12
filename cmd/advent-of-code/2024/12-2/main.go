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
			if s1 == s2 {
				ds.UnionWithAdd(p, q)
			}
		})
	})

	regions := make(map[Point2D]Set[Point2D])
	grid.ForEachPoint(func(p Point2D, _ string) {
		f, _ := ds.Find(p)
		regions[f] = regions[f].UnionElems(p)
	})

	var total int
	for _, ps := range regions {
		var corners int
		for p := range ps {
			up, right, down, left := p.Up(), p.Right(), p.Down(), p.Left()

			if !ps.Contains(up) && !ps.Contains(right) {
				corners++
			}
			if !ps.Contains(up) && !ps.Contains(left) {
				corners++
			}
			if !ps.Contains(down) && !ps.Contains(right) {
				corners++
			}
			if !ps.Contains(down) && !ps.Contains(left) {
				corners++
			}
			if ps.Contains(up) && ps.Contains(right) && !ps.Contains(up.Right()) {
				corners++
			}
			if ps.Contains(up) && ps.Contains(left) && !ps.Contains(up.Left()) {
				corners++
			}
			if ps.Contains(down) && ps.Contains(right) && !ps.Contains(down.Right()) {
				corners++
			}
			if ps.Contains(down) && ps.Contains(left) && !ps.Contains(down.Left()) {
				corners++
			}
		}

		total += len(ps) * corners
	}
	fmt.Println(total)
}
