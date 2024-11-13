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

type Grid map[puz.Point2D]bool

func InputToGrid() (Grid, puz.Point2D) {
	grid := make(Grid)
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for y, line := range puz.InputToLines(2017, 22) {
		for x, c := range line {
			grid[puz.Point2D{X: x, Y: y}] = c == '#'
			minX = puz.Min(minX, x)
			maxX = puz.Max(maxX, x)
		}
		minY = puz.Min(minY, y)
		maxY = puz.Max(maxY, y)
	}

	center := puz.Point2D{X: (maxX - minX) / 2, Y: (maxY - minY) / 2}
	return grid, center
}
