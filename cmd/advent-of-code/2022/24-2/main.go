package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	children := func(s State) []State {
		next := GetGrid(s.Time + 1)

		var children []State

		// We can only remain in our current location if nothing hits it
		if next.GetPoint(s.Location) == 0 {
			children = append(children, State{Time: s.Time + 1, Location: s.Location})
		}

		next.ForEachOrthogonalNeighbor(s.Location, func(n lib.Point2D, bs lib.BitSet) {
			if bs == 0 {
				children = append(children, State{Time: s.Time + 1, Location: n})
			}
		})

		return children
	}

	W, H := GetGrid(0).Width, GetGrid(0).Height
	path1, _ := lib.BreadthFirstSearch(
		State{Location: lib.Point2D{X: 1, Y: 0}},
		children,
		func(s State) bool { return s.Location.X == W-2 && s.Location.Y == H-1 },
	)

	path2, _ := lib.BreadthFirstSearch(
		path1[len(path1)-1],
		children,
		func(s State) bool { return s.Location.X == 1 && s.Location.Y == 0 },
	)

	path3, _ := lib.BreadthFirstSearch(
		path2[len(path2)-1],
		children,
		func(s State) bool { return s.Location.X == W-2 && s.Location.Y == H-1 },
	)

	fmt.Println(len(path1) + len(path2) + len(path3) - 3)
}

type State struct {
	Time     int
	Location lib.Point2D
}

var grids []lib.Grid2D[lib.BitSet]

func GetGrid(tm int) lib.Grid2D[lib.BitSet] {
	if len(grids) == 0 {
		grids = append(grids, InputToGrid())
	}

	for len(grids) <= tm {
		grids = append(grids, Next(grids[len(grids)-1]))
	}

	return grids[tm]
}

func Next(g lib.Grid2D[lib.BitSet]) lib.Grid2D[lib.BitSet] {
	next := lib.NewGrid2D[lib.BitSet](g.Width, g.Height)
	g.ForEach(func(x, y int, bs lib.BitSet) {
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

func InputToGrid() lib.Grid2D[lib.BitSet] {
	var bs lib.BitSet
	return lib.InputToGrid2D(func(x int, y int, s string) lib.BitSet {
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
