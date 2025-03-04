package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid, center := InputToGrid()
	turtle := Turtle{Location: center}

	var infections int
	for n := 0; n < 10000000; n++ {
		var state = Clean
		if s, ok := grid[turtle.Location]; ok {
			state = s
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

const (
	Clean    = "."
	Infected = "#"
	Weakened = "W"
	Flagged  = "F"
)

func InputToGrid() (map[Point2D]string, Point2D) {
	var grid = make(map[Point2D]string)

	var x, y int
	var ch rune
	for in.HasNext() {
		for x, ch = range in.Line() {
			grid[Point2D{X: x, Y: y}] = string(ch)
		}

		y++
	}

	return grid, Point2D{X: x / 2, Y: y / 2}
}
