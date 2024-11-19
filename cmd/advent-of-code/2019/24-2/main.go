package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

const N = 5

func main() {
	grids := []lib.BitSet{InputToGrid()}
	for tm := 0; tm < 200; tm++ {
		grids = Next(grids)
	}

	var count int
	for _, grid := range grids {
		count += grid.Size()
	}
	fmt.Println(count)
}

func Next(grids []lib.BitSet) []lib.BitSet {
	get := func(depth, x, y int) int {
		if depth < 1 || depth > len(grids) || x < 0 || x >= N || y < 0 || y >= N {
			return 0
		}

		if grids[depth-1].Contains(y*N + x) {
			return 1
		}
		return 0
	}

	next := make([]lib.BitSet, len(grids)+2)
	for n := 0; n < len(next); n++ {
		for x := 0; x < N; x++ {
			for y := 0; y < N; y++ {
				if x == 2 && y == 2 {
					continue
				}

				var count int
				ForEachNeighbor(x, y, func(x, y, delta int) {
					count += get(n+delta, x, y)
				})

				if get(n, x, y) == 1 && count == 1 {
					next[n] = next[n].Add(y*N + x)
				} else if get(n, x, y) == 0 && (count == 1 || count == 2) {
					next[n] = next[n].Add(y*N + x)
				}
			}
		}
	}

	return next
}

func ForEachNeighbor(x, y int, fn func(x, y, delta int)) {
	// N
	if y == 0 {
		fn(2, 1, -1)
	} else if y == 3 && x == 2 {
		fn(0, 4, +1)
		fn(1, 4, +1)
		fn(2, 4, +1)
		fn(3, 4, +1)
		fn(4, 4, +1)
	} else {
		fn(x, y-1, 0)
	}

	// W
	if x == 0 {
		fn(1, 2, -1)
	} else if x == 3 && y == 2 {
		fn(4, 0, +1)
		fn(4, 1, +1)
		fn(4, 2, +1)
		fn(4, 3, +1)
		fn(4, 4, +1)
	} else {
		fn(x-1, y, 0)
	}

	// E
	if x == 4 {
		fn(3, 2, -1)
	} else if x == 1 && y == 2 {
		fn(0, 0, +1)
		fn(0, 1, +1)
		fn(0, 2, +1)
		fn(0, 3, +1)
		fn(0, 4, +1)
	} else {
		fn(x+1, y, 0)
	}

	// S
	if y == 4 {
		fn(2, 3, -1)
	} else if y == 1 && x == 2 {
		fn(0, 0, +1)
		fn(1, 0, +1)
		fn(2, 0, +1)
		fn(3, 0, +1)
		fn(4, 0, +1)
	} else {
		fn(x, y+1, 0)
	}
}

func InputToGrid() lib.BitSet {
	var grid lib.BitSet
	for y, line := range lib.InputToLines() {
		for x, c := range line {
			if c == '#' {
				grid = grid.Add(y*N + x)
			}
		}
	}

	return grid
}
