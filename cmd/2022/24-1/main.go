package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	children := func(s State) []State {
		next := GetGrid(s.Time + 1)

		var children []State

		// We can only remain in our current location if nothing hits it
		if next.GetPoint(s.Location) == 0 {
			children = append(children, State{Time: s.Time + 1, Location: s.Location})
		}

		next.ForEachOrthogonalNeighbor(s.Location, func(n aoc.Point2D, bs aoc.BitSet) {
			if bs == 0 {
				children = append(children, State{Time: s.Time + 1, Location: n})
			}
		})

		return children
	}

	W, H := GetGrid(0).Width, GetGrid(0).Height
	path, _ := aoc.BreadthFirstSearch(
		State{Location: aoc.Point2D{X: 1, Y: 0}},
		children,
		func(s State) bool { return s.Location.X == W-2 && s.Location.Y == H-1 },
	)

	fmt.Println(len(path) - 1)
}

type State struct {
	Time     int
	Location aoc.Point2D
}

var grids []aoc.Grid2D[aoc.BitSet]

func GetGrid(tm int) aoc.Grid2D[aoc.BitSet] {
	if len(grids) == 0 {
		grids = append(grids, InputToGrid())
	}

	for len(grids) <= tm {
		grids = append(grids, Next(grids[len(grids)-1]))
	}

	return grids[tm]
}

func Next(g aoc.Grid2D[aoc.BitSet]) aoc.Grid2D[aoc.BitSet] {
	next := aoc.NewGrid2D[aoc.BitSet](g.Width, g.Height)
	g.ForEach(func(x, y int, bs aoc.BitSet) {
		if bs.Contains(WALL) {
			next.Set(x, y, next.Get(x, y).Add(WALL))
			return
		}

		if bs.Contains(UP) {
			if !g.Get(x, y-1).Contains(WALL) {
				next.Set(x, y-1, next.Get(x, y-1).Add(UP))
			} else {
				next.Set(x, g.Height-2, next.Get(x, g.Height-2).Add(UP))
			}
		}

		if bs.Contains(DOWN) {
			if !g.Get(x, y+1).Contains(WALL) {
				next.Set(x, y+1, next.Get(x, y+1).Add(DOWN))
			} else {
				next.Set(x, 1, next.Get(x, 1).Add(DOWN))
			}
		}

		if bs.Contains(LEFT) {
			if !g.Get(x-1, y).Contains(WALL) {
				next.Set(x-1, y, next.Get(x-1, y).Add(LEFT))
			} else {
				next.Set(g.Width-2, y, next.Get(g.Width-2, y).Add(LEFT))
			}
		}

		if bs.Contains(RIGHT) {
			if !g.Get(x+1, y).Contains(WALL) {
				next.Set(x+1, y, next.Get(x+1, y).Add(RIGHT))
			} else {
				next.Set(1, y, next.Get(1, y).Add(RIGHT))
			}
		}
	})

	return next
}

const (
	WALL = iota
	UP
	DOWN
	LEFT
	RIGHT
)

func InputToGrid() aoc.Grid2D[aoc.BitSet] {
	var bs aoc.BitSet
	return aoc.InputToGrid2D(2022, 24, func(x int, y int, s string) aoc.BitSet {
		switch s {
		case "#":
			return bs.Add(WALL)
		case "^":
			return bs.Add(UP)
		case ">":
			return bs.Add(RIGHT)
		case "<":
			return bs.Add(LEFT)
		case "v":
			return bs.Add(DOWN)
		}
		return bs
	})
}
