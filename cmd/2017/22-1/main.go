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
	for round := 0; round < 10000; round++ {
		// if the current node is infected turn right, otherwise turn left
		switch {
		case grid[position] && dir == "U":
			dir = "R"
		case grid[position] && dir == "D":
			dir = "L"
		case grid[position] && dir == "L":
			dir = "U"
		case grid[position] && dir == "R":
			dir = "D"

		case !grid[position] && dir == "U":
			dir = "L"
		case !grid[position] && dir == "D":
			dir = "R"
		case !grid[position] && dir == "L":
			dir = "D"
		case !grid[position] && dir == "R":
			dir = "U"
		}

		// if the current node is clean, infect it, otherwise clean it
		if !grid[position] {
			infects++
		}
		grid[position] = !grid[position]

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

type Grid map[aoc.Point2D]bool

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
			if g[aoc.Point2D{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
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
			grid[p] = c == '#'
		}
	}

	return grid
}
