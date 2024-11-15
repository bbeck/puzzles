package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"sort"
)

func main() {
	m := puz.InputToIntGrid2D()

	var ds puz.DisjointSet[puz.Point2D]
	m.ForEachPoint(func(p puz.Point2D, height int) {
		if height == 9 {
			return
		}

		m.ForEachOrthogonalNeighbor(p, func(n puz.Point2D, height int) {
			if height == 9 {
				return
			}

			ds.UnionWithAdd(p, n)
		})
	})

	index := make(map[puz.Point2D]int)
	m.ForEachPoint(func(p puz.Point2D, height int) {
		if root, ok := ds.Find(p); ok {
			index[root] = ds.Size(root)
		}
	})

	sizes := puz.GetMapValues(index)
	sort.Ints(sizes)

	fmt.Println(puz.Product(sizes[len(sizes)-3:]...))
}
