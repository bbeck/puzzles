package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	grid := puz.InputToStringGrid2D(2023, 23)
	graph, start, end := BuildGraph(grid)

	var longest int

	var dfs func(current int, seen puz.BitSet, length int)
	dfs = func(current int, seen puz.BitSet, length int) {
		if current == end {
			longest = puz.Max(longest, length)
		}

		for p, cost := range graph[current] {
			if cost != 0 && !seen.Contains(p) {
				dfs(p, seen.Add(p), length+cost)
			}
		}
	}
	dfs(start, 0, 0)

	fmt.Println(longest)
}

type Graph [][]int

func BuildGraph(grid puz.Grid2D[string]) (Graph, int, int) {
	vertices, start, end := FindVertices(grid)
	graph := puz.Make2D[int](len(vertices), len(vertices))

	blocked := puz.SetFrom(vertices...)
	for i, v1 := range vertices {
		blocked.Remove(v1)
		for j, v2 := range vertices {
			blocked.Remove(v2)
			if d, ok := Distance(grid, v1, v2, blocked); ok {
				graph[i][j] = d
				graph[j][i] = d
			}
			blocked.Add(v2)
		}
		blocked.Add(v1)
	}
	return graph, start, end
}

func FindVertices(grid puz.Grid2D[string]) ([]puz.Point2D, int, int) {
	var vertices []puz.Point2D
	var start, end int
	grid.ForEachPoint(func(p puz.Point2D, s string) {
		if s == "#" {
			return
		}

		// Include the start/end points as if they were forks
		if p.Y == 0 {
			start = len(vertices)
			vertices = append(vertices, p)
			return
		}
		if p.Y == grid.Height-1 {
			end = len(vertices)
			vertices = append(vertices, p)
			return
		}

		var count int
		grid.ForEachOrthogonalNeighbor(p, func(q puz.Point2D, s string) {
			if s != "#" {
				count++
			}
		})
		if count > 2 {
			vertices = append(vertices, p)
		}
	})

	return vertices, start, end
}

func Distance(grid puz.Grid2D[string], start, end puz.Point2D, blocked puz.Set[puz.Point2D]) (int, bool) {
	children := func(p puz.Point2D) []puz.Point2D {
		var children []puz.Point2D
		grid.ForEachOrthogonalNeighbor(p, func(q puz.Point2D, ch string) {
			if ch != "#" && !blocked.Contains(q) {
				children = append(children, q)
			}
		})
		return children
	}
	goal := func(p puz.Point2D) bool { return p == end }

	path, ok := puz.BreadthFirstSearch(start, children, goal)
	return len(path) - 1, ok
}
