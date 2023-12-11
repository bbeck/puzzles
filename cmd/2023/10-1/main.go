package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid, start := InputToGridAndStart()

	children := func(p aoc.Point2D) []aoc.Point2D {
		cell := grid.GetPoint(p)

		var children []aoc.Point2D
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

	var loop aoc.Set[aoc.Point2D]
	isGoal := func(p aoc.Point2D) bool {
		loop.Add(p)
		return false
	}

	aoc.BreadthFirstSearch(start, children, isGoal)
	fmt.Println(len(loop) / 2)
}

type Cell struct {
	N, S, E, W bool
}

func InputToGridAndStart() (aoc.Grid2D[Cell], aoc.Point2D) {
	var start aoc.Point2D
	grid := aoc.InputToGrid2D(2023, 10, func(x int, y int, s string) Cell {
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
			start = aoc.Point2D{X: x, Y: y}
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
