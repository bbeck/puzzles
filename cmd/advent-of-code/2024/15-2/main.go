package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
)

func main() {
	grid, robot, moves := InputToGridRobotMoves()
	for _, move := range moves {
		grid, robot = Move(grid, robot, move)
	}

	var sum int
	grid.ForEach(func(x int, y int, s string) {
		if s == "[" {
			sum += 100*y + x
		}
	})
	fmt.Println(sum)
}

func Move(grid Grid2D[string], r Point2D, move string) (Grid2D[string], Point2D) {
	var neighbor Point2D
	switch move {
	case "^":
		neighbor = r.Up()
	case ">":
		neighbor = r.Right()
	case "v":
		neighbor = r.Down()
	case "<":
		neighbor = r.Left()
	}

	if move == "<" || move == ">" {
		// We're moving horizontally, we don't ever have to worry about splitting a
		// box.
		switch grid.GetPoint(neighbor) {
		case "#":
			return grid, r

		case ".":
			grid.SetPoint(neighbor, grid.GetPoint(r))
			grid.SetPoint(r, ".")
			return grid, neighbor

		default: // One side of a box
			Move(grid, neighbor, move)
			if grid.GetPoint(neighbor) == "." {
				grid.SetPoint(neighbor, grid.GetPoint(r))
				grid.SetPoint(r, ".")
				return grid, neighbor
			}
			return grid, r // We couldn't move everything out of the way
		}
	}

	if move == "^" || move == "v" {
		// We're moving vertically, we need to make sure we don't split boxes.
		switch grid.GetPoint(neighbor) {
		case "#":
			return grid, r

		case ".":
			grid.SetPoint(neighbor, grid.GetPoint(r))
			grid.SetPoint(r, ".")
			return grid, neighbor

		case "[":
			save := Clone(grid)
			save, _ = Move(save, neighbor, move)
			save, _ = Move(save, neighbor.Right(), move)
			if save.GetPoint(neighbor) == "." && save.GetPoint(neighbor.Right()) == "." {
				save.SetPoint(neighbor, save.GetPoint(r))
				save.SetPoint(r, ".")
				return save, neighbor
			}
			return grid, r

		case "]":
			save := Clone(grid)
			save, _ = Move(save, neighbor, move)
			save, _ = Move(save, neighbor.Left(), move)
			if save.GetPoint(neighbor) == "." && save.GetPoint(neighbor.Left()) == "." {
				save.SetPoint(neighbor, save.GetPoint(r))
				save.SetPoint(r, ".")
				return save, neighbor
			}
			return grid, r
		}
	}

	return grid, Origin2D // Should never happen
}

func Clone(g Grid2D[string]) Grid2D[string] {
	grid := NewGrid2D[string](g.Width, g.Height)
	g.ForEachPoint(func(p Point2D, s string) {
		grid.SetPoint(p, s)
	})
	return grid
}

func InputToGridRobotMoves() (Grid2D[string], Point2D, []string) {
	lines := InputToLines()

	var blank int
	for i, line := range lines {
		if line == "" {
			blank = i
			break
		}
	}

	grid := NewGrid2D[string](2*len(lines[0]), blank)
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < len(lines[y]); x++ {
			switch s := string(lines[y][x]); s {
			case "O":
				grid.Set(2*x, y, "[")
				grid.Set(2*x+1, y, "]")
			case "@":
				grid.Set(2*x, y, "@")
				grid.Set(2*x+1, y, ".")
			default:
				grid.Set(2*x, y, s)
				grid.Set(2*x+1, y, s)
			}
		}
	}

	var robot Point2D
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "@" {
			robot = p
		}
	})

	var moves []string
	for _, line := range lines[blank+1:] {
		for _, c := range line {
			moves = append(moves, string(c))
		}
	}

	return grid, robot, moves
}
