package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	grid, start := InputToGridAndStart()
	grid, start = Triple(grid, start)

	// Remove any pipe that's not part of the loop.
	loop := GetLoopPoints(grid, start)
	grid.ForEachPoint(func(p lib.Point2D, _ Cell) {
		if !loop.Contains(p) {
			grid.SetPoint(p, Cell{})
		}
	})

	// Flood fill from the edge.
	var seen lib.Set[lib.Point2D]
	for x := 0; x < grid.Width; x++ {
		seen = Flood(grid, lib.Point2D{X: x, Y: 0}, seen)
		seen = Flood(grid, lib.Point2D{X: x, Y: grid.Height - 1}, seen)
	}
	for y := 0; y < grid.Height; y++ {
		seen = Flood(grid, lib.Point2D{X: 0, Y: y}, seen)
		seen = Flood(grid, lib.Point2D{X: grid.Width - 1, Y: y}, seen)
	}

	var enclosed int
	grid.ForEachPoint(func(p lib.Point2D, _ Cell) {
		// Only consider points at the center of a 3x3 block.  These are the points
		// that are in the original, unexpanded grid.
		if p.X%3 != 1 || p.Y%3 != 1 {
			return
		}

		// If the point was on the loop or was seen during the flood fill then it
		// can't be enclosed
		if loop.Contains(p) || seen.Contains(p) {
			return
		}

		enclosed++
	})
	fmt.Println(enclosed)
}

func Triple(grid lib.Grid2D[Cell], start lib.Point2D) (lib.Grid2D[Cell], lib.Point2D) {
	next := lib.NewGrid2D[Cell](3*grid.Width, 3*grid.Height)
	grid.ForEach(func(x int, y int, cell Cell) {
		center := lib.Point2D{X: 3*x + 1, Y: 3*y + 1}
		next.SetPoint(center, cell)
		if cell.N {
			next.SetPoint(center.Up(), Cell{N: true, S: true})
		}
		if cell.E {
			next.SetPoint(center.Right(), Cell{E: true, W: true})
		}
		if cell.S {
			next.SetPoint(center.Down(), Cell{N: true, S: true})
		}
		if cell.W {
			next.SetPoint(center.Left(), Cell{E: true, W: true})
		}
	})

	return next, lib.Point2D{X: 3*start.X + 1, Y: 3*start.Y + 1}
}

func GetLoopPoints(grid lib.Grid2D[Cell], start lib.Point2D) lib.Set[lib.Point2D] {
	children := func(p lib.Point2D) []lib.Point2D {
		cell := grid.GetPoint(p)

		var children []lib.Point2D
		if q := p.Up(); cell.N && grid.InBoundsPoint(q) {
			children = append(children, q)
		}
		if q := p.Right(); cell.E && grid.InBoundsPoint(q) {
			children = append(children, q)
		}
		if q := p.Down(); cell.S && grid.InBoundsPoint(q) {
			children = append(children, q)
		}
		if q := p.Left(); cell.W && grid.InBoundsPoint(q) {
			children = append(children, q)
		}
		return children
	}

	var loop lib.Set[lib.Point2D]
	isGoal := func(p lib.Point2D) bool {
		loop.Add(p)
		return false
	}

	lib.BreadthFirstSearch(start, children, isGoal)
	return loop
}

func Flood(grid lib.Grid2D[Cell], p lib.Point2D, seen lib.Set[lib.Point2D]) lib.Set[lib.Point2D] {
	if !seen.Add(p) || !grid.GetPoint(p).IsGround() {
		return seen
	}

	grid.ForEachOrthogonalNeighborPoint(p, func(q lib.Point2D, _ Cell) {
		seen = Flood(grid, q, seen)
	})
	return seen
}

type Cell struct {
	N, S, E, W bool
}

func (c Cell) IsGround() bool {
	return !c.N && !c.S && !c.E && !c.W
}

func InputToGridAndStart() (lib.Grid2D[Cell], lib.Point2D) {
	var start lib.Point2D
	grid := lib.InputToGrid2D(func(x int, y int, s string) Cell {
		switch s {
		case "|": // | is a vertical pipe connecting north and south.
			return Cell{N: true, S: true}
		case "-": // - is a horizontal pipe connecting east and west.
			return Cell{E: true, W: true}
		case "L": // L is a 90-degree bend connecting north and east.
			return Cell{N: true, E: true}
		case "J": // J is a 90-degree bend connecting north and west.
			return Cell{N: true, W: true}
		case "7": // 7 is a 90-degree bend connecting south and west.
			return Cell{S: true, W: true}
		case "F": // F is a 90-degree bend connecting south and east.
			return Cell{S: true, E: true}
		case "S": // S is the starting position of the animal
			start = lib.Point2D{X: x, Y: y}
		}
		return Cell{}
	})

	// Infer what kind of pipe is at the start.
	var cell Cell
	if p := start.Up(); grid.InBoundsPoint(p) && grid.GetPoint(p).S {
		cell.N = true
	}
	if p := start.Right(); grid.InBoundsPoint(p) && grid.GetPoint(p).W {
		cell.E = true
	}
	if p := start.Down(); grid.InBoundsPoint(p) && grid.GetPoint(p).N {
		cell.S = true
	}
	if p := start.Left(); grid.InBoundsPoint(p) && grid.GetPoint(p).E {
		cell.W = true
	}
	grid.SetPoint(start, cell)

	return grid, start
}
