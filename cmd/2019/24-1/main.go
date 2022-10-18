package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

const N = 5

func main() {
	// Since there are so few cells in the grid we'll pack its bits into a
	// number instead of using a dynamically sized data structure.  This will
	// make determining repeated states trivial.  Additionally, we'll organize
	// the bits such that the value of the set is the biodiversity rating that
	// we seek.
	grid := InputToGrid()

	var seen aoc.Set[aoc.BitSet]
	seen.Add(grid)

	for {
		grid = Next(grid)
		if !seen.Add(grid) {
			break
		}
	}

	fmt.Println(grid)
}

func Next(grid aoc.BitSet) aoc.BitSet {
	get := func(x, y int) int {
		if 0 <= x && x < N && 0 <= y && y < N && grid.Contains(y*N+x) {
			return 1
		}
		return 0
	}

	var next aoc.BitSet
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			count := get(x-1, y) + get(x+1, y) + get(x, y-1) + get(x, y+1)
			if get(x, y) == 1 && count == 1 {
				next = next.Add(y*N + x)
			} else if get(x, y) == 0 && (count == 1 || count == 2) {
				next = next.Add(y*N + x)
			}
		}
	}

	return next
}

func InputToGrid() aoc.BitSet {
	var grid aoc.BitSet
	for y, line := range aoc.InputToLines(2019, 24) {
		for x, c := range line {
			if c == '#' {
				grid = grid.Add(y*N + x)
			}
		}
	}

	return grid
}
