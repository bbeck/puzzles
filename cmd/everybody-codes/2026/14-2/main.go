package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid, names := InputToIdGrid()

	var costs = Make2D[int](len(names), len(names))
	grid.ForEach(func(x int, y int, s int) {
		if x == 0 || x == grid.Width-1 || y == 0 || y == grid.Height-1 {
			costs[s][0]++
			costs[0][s]++
		}

		grid.ForEachOrthogonalNeighbor(x, y, func(_ int, _ int, t int) {
			if s != t {
				costs[s][t]++
			}
		})
	})

	children := func(parent int) []int {
		var children []int
		for id, cost := range costs[parent] {
			if cost > 0 {
				children = append(children, id)
			}
		}
		return children
	}

	cost := func(from, to int) int {
		return costs[from][to]
	}

	n, _ := MinimumSpanningTree(Keys(names), children, cost)
	fmt.Println(n)
}

func InputToIdGrid() (Grid2D[int], map[int]string) {
	var sgrid = in.ToStringGrid2D()
	var ids = map[string]int{"ROOT": 0}
	var names = map[int]string{0: "ROOT"}

	var grid = NewGrid2D[int](sgrid.Width, sgrid.Height)
	sgrid.ForEach(func(x int, y int, s string) {
		if _, ok := ids[s]; !ok {
			ids[s] = len(ids)
			names[ids[s]] = s
		}
		grid.Set(x, y, ids[s])
	})

	return grid, names
}
