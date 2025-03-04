package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var grid, center = InputToGrid()
	var turtle = Turtle{Location: center}

	var infections int
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

func InputToGrid() (map[Point2D]bool, Point2D) {
	var grid = make(map[Point2D]bool)

	var x, y int
	var ch rune
	for in.HasNext() {
		for x, ch = range in.Line() {
			grid[Point2D{X: x, Y: y}] = ch == '#'
		}

		y++
	}

	return grid, Point2D{X: x / 2, Y: y / 2}
}
