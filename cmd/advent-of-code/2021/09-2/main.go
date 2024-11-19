package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"sort"
)

func main() {
	m := lib.InputToIntGrid2D()

	var ds lib.DisjointSet[lib.Point2D]
	m.ForEachPoint(func(p lib.Point2D, height int) {
		if height == 9 {
			return
		}

		m.ForEachOrthogonalNeighbor(p, func(n lib.Point2D, height int) {
			if height == 9 {
				return
			}

			ds.UnionWithAdd(p, n)
		})
	})

	index := make(map[lib.Point2D]int)
	m.ForEachPoint(func(p lib.Point2D, height int) {
		if root, ok := ds.Find(p); ok {
			index[root] = ds.Size(root)
		}
	})

	sizes := lib.GetMapValues(index)
	sort.Ints(sizes)

	fmt.Println(lib.Product(sizes[len(sizes)-3:]...))
}
