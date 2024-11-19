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
	for n := 0; n < 10000; n++ {
		if grid[turtle.Location] {
			turtle.TurnRight()
			grid[turtle.Location] = false
		} else {
			turtle.TurnLeft()
			grid[turtle.Location] = true
			infections++
		}

		turtle.Forward(1)
	}

	fmt.Println(infections)
}

type Grid map[lib.Point2D]bool

func InputToGrid() (Grid, lib.Point2D) {
	grid := make(Grid)
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for y, line := range lib.InputToLines() {
		for x, c := range line {
			grid[lib.Point2D{X: x, Y: y}] = c == '#'
			minX = lib.Min(minX, x)
			maxX = lib.Max(maxX, x)
		}
		minY = lib.Min(minY, y)
		maxY = lib.Max(maxY, y)
	}

	center := lib.Point2D{X: (maxX - minX) / 2, Y: (maxY - minY) / 2}
	return grid, center
}
