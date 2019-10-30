package main

import (
	"fmt"
	"math"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := InputToGrid(2017, 22)
	position := grid.Center()
	dir := "U"

	var infects int
	for round := 0; round < 10000000; round++ {
		// if the node is clean turn left
		// if the node is weakened do not turn
		// if the node is infected turn right
		// if the node is flagged reverse direction
		switch grid[position] {
		case Clean:
			if dir == "U" {
				dir = "L"
			} else if dir == "D" {
				dir = "R"
			} else if dir == "L" {
				dir = "D"
			} else if dir == "R" {
				dir = "U"
			}
		case Weakened:
		case Infected:
			if dir == "U" {
				dir = "R"
			} else if dir == "D" {
				dir = "L"
			} else if dir == "L" {
				dir = "U"
			} else if dir == "R" {
				dir = "D"
			}
		case Flagged:
			if dir == "U" {
				dir = "D"
			} else if dir == "D" {
				dir = "U"
			} else if dir == "L" {
				dir = "R"
			} else if dir == "R" {
				dir = "L"
			}
		}

		// clean -> weakened
		// weakened -> infected
		// infected -> flagged
		// flagged -> clean
		switch grid[position] {
		case Clean:
			grid[position] = Weakened
		case Weakened:
			grid[position] = Infected
			infects++
		case Infected:
			grid[position] = Flagged
		case Flagged:
			grid[position] = Clean
		}

		// move forward one node
		switch dir {
		case "U":
			position = position.Up()
		case "D":
			position = position.Down()
		case "L":
			position = position.Left()
		case "R":
			position = position.Right()
		}
	}

	fmt.Printf("number of infects: %d\n", infects)
}

type State int

const (
	Clean    State = 0
	Weakened State = 1
	Infected State = 2
	Flagged  State = 3
)

type Grid map[aoc.Point2D]State

func (g Grid) Center() aoc.Point2D {
	minX, minY := math.MaxInt64, math.MaxInt64
	maxX, maxY := math.MinInt64, math.MinInt64
	for p := range g {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	return aoc.Point2D{
		X: (maxX - minX) / 2,
		Y: (maxY - minY) / 2,
	}
}

func (g Grid) Print() {
	minX, minY := math.MaxInt64, math.MaxInt64
	maxX, maxY := math.MinInt64, math.MinInt64
	for p := range g {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {

			switch g[aoc.Point2D{X: x, Y: y}] {
			case Clean:
				fmt.Print(".")
			case Weakened:
				fmt.Print("W")
			case Infected:
				fmt.Print("#")
			case Flagged:
				fmt.Print("F")
			}
		}
		fmt.Println()
	}
}

func InputToGrid(year, day int) Grid {
	grid := make(Grid)
	for y, line := range aoc.InputToLines(year, day) {
		for x, c := range line {
			p := aoc.Point2D{x, y}
			if c == '#' {
				grid[p] = Infected
			} else {
				grid[p] = Clean
			}
		}
	}

	return grid
}
