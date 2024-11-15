package main

import (
	"fmt"
	"math"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var infections int

	grid, center := InputToGrid()
	turtle := puz.Turtle{Location: center}
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

type Grid map[puz.Point2D]string

const (
	Clean    = "."
	Infected = "#"
	Weakened = "W"
	Flagged  = "F"
)

func InputToGrid() (Grid, puz.Point2D) {
	grid := make(Grid)
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for y, line := range puz.InputToLines() {
		for x, c := range line {
			grid[puz.Point2D{X: x, Y: y}] = string(c)
			minX = puz.Min(minX, x)
			maxX = puz.Max(maxX, x)
		}
		minY = puz.Min(minY, y)
		maxY = puz.Max(maxY, y)
	}

	center := puz.Point2D{X: (maxX - minX) / 2, Y: (maxY - minY) / 2}
	return grid, center
}
