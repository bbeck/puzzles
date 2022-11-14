package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"sort"
)

func main() {
	m := InputToHeightMap()

	var ds aoc.DisjointSet[aoc.Point2D]
	m.ForEach(func(p aoc.Point2D, height int) {
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
	m.ForEach(func(p aoc.Point2D, height int) {
		if root, ok := ds.Find(p); ok {
			index[root] = ds.Size(root)
		}
	})

	sizes := aoc.GetMapValues(index)
	sort.Ints(sizes)

	fmt.Println(aoc.Product(sizes[len(sizes)-3:]...))
}

func InputToHeightMap() aoc.Grid2D[int] {
	lines := aoc.InputToLines(2021, 9)

	grid := aoc.NewGrid2D[int](len(lines[0]), len(lines))
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			grid.AddXY(x, y, aoc.ParseInt(string(lines[y][x])))
		}
	}

	return grid
}
