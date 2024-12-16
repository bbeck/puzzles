package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
)

func main() {
	grid, robot, moves := InputToGridRobotMoves()
	for _, move := range moves {
		robot = Move(grid, robot, move)
	}

	var sum int
	grid.ForEach(func(x int, y int, s string) {
		if s == "O" {
			sum += 100*y + x
		}
	})
	fmt.Println(sum)
}

func Move(grid Grid2D[string], r Point2D, move string) Point2D {
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

	switch grid.GetPoint(neighbor) {
	case ".":
		grid.SetPoint(neighbor, grid.GetPoint(r))
		grid.SetPoint(r, ".")
		return neighbor
	case "O":
		Move(grid, neighbor, move)
		if grid.GetPoint(neighbor) == "." {
			grid.SetPoint(neighbor, grid.GetPoint(r))
			grid.SetPoint(r, ".")
			return neighbor
		}
		return r
	default:
		return r
	}
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

	grid := NewGrid2D[string](len(lines[0]), blank)
	grid.ForEach(func(x int, y int, _ string) {
		s := string(lines[y][x])
		grid.Set(x, y, s)

	})

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
