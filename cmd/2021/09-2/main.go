package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"sort"
)

func main() {
	m := InputToHeightMap()

	var ds aoc.DisjointSet[aoc.Point2D]
	m.ForEachPoint(func(p aoc.Point2D, height int) {
		if height == 9 {
			return
		}

		m.ForEachOrthogonalNeighbor(p, func(n aoc.Point2D, height int) {
			if height == 9 {
				return
			}

			ds.UnionWithAdd(p, n)
		})
	})

	index := make(map[aoc.Point2D]int)
	m.ForEachPoint(func(p aoc.Point2D, height int) {
		if root, ok := ds.Find(p); ok {
			index[root] = ds.Size(root)
		}
	})

	sizes := aoc.GetMapValues(index)
	sort.Ints(sizes)

	fmt.Println(aoc.Product(sizes[len(sizes)-3:]...))
}

func InputToHeightMap() aoc.Grid2D[int] {
	return aoc.InputToGrid2D(2021, 9, func(x int, y int, s string) int {
		return aoc.ParseInt(s)
	})
}
