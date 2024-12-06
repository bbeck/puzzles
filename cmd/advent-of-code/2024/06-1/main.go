package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
)

func main() {
	grid := InputToStringGrid2D()

	var guard Turtle
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "^" {
			guard.Location = p
			guard.Heading = Up
		}
	})

	var seen Set[Point2D]
	for {
		seen.Add(guard.Location)

		mark := guard
		guard.Forward(1)
		if !grid.InBoundsPoint(guard.Location) {
			break
		}

		if s := grid.GetPoint(guard.Location); s != "." && s != "^" {
			guard = mark
			guard.TurnRight()
		}
	}
	fmt.Println(len(seen))
}
