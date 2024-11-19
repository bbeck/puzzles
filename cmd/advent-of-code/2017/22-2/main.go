package main

import (
	"fmt"
	"math"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	var infections int

	grid, center := InputToGrid()
	turtle := lib.Turtle{Location: center}
	for n := 0; n < 10000000; n++ {
		var state string
		if s, ok := grid[turtle.Location]; ok {
			state = s
		} else {
			state = Clean
		}

		switch state {
		case Clean:
			turtle.TurnLeft()
			grid[turtle.Location] = Weakened
		case Weakened:
			grid[turtle.Location] = Infected
			infections++
		case Infected:
			turtle.TurnRight()
			grid[turtle.Location] = Flagged
		case Flagged:
			turtle.TurnRight()
			turtle.TurnRight()
			grid[turtle.Location] = Clean
		}

		turtle.Forward(1)
	}

	fmt.Println(infections)
}

type Grid map[lib.Point2D]string

const (
	Clean    = "."
	Infected = "#"
	Weakened = "W"
	Flagged  = "F"
)

func InputToGrid() (Grid, lib.Point2D) {
	grid := make(Grid)
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for y, line := range lib.InputToLines() {
		for x, c := range line {
			grid[lib.Point2D{X: x, Y: y}] = string(c)
			minX = lib.Min(minX, x)
			maxX = lib.Max(maxX, x)
		}
		minY = lib.Min(minY, y)
		maxY = lib.Max(maxY, y)
	}

	center := lib.Point2D{X: (maxX - minX) / 2, Y: (maxY - minY) / 2}
	return grid, center
}
