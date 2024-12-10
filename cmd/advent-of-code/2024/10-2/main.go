package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
)

func main() {
	var heads Set[Point2D]
	grid := InputToGrid2D(func(x int, y int, s string) int {
		n := ParseInt(s)
		if n == 0 {
			heads.Add(Point2D{X: x, Y: y})
		}
		return n
	})

	var sum int
	for head := range heads {
		sum += Count(grid, head, grid.GetPoint(head), make(map[Point2D]int))
	}
	fmt.Println(sum)
}

func Count(grid Grid2D[int], p Point2D, v int, memo map[Point2D]int) int {
	if answer, ok := memo[p]; ok {
		return answer
	}

	if v == 9 {
		memo[p] = 1
		return 1
	}

	var ways int
	for _, n := range p.OrthogonalNeighbors() {
		if grid.InBoundsPoint(n) && grid.GetPoint(n) == v+1 {
			ways += Count(grid, n, v+1, memo)
		}
	}

	memo[p] = ways
	return ways
}
