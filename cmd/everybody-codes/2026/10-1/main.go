package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid, ghosts := InputToGridAndGhosts()

	var collisions int
	for range 10 {
		var seen = make(map[Point2D][]int)
		for idx := range ghosts {
			ghosts[idx].Forward(1)
			if grid.GetPoint(ghosts[idx].Location) == "#" {
				ghosts[idx].TurnLeft()
				ghosts[idx].TurnLeft()
				ghosts[idx].Forward(1)
				ghosts[idx].TurnLeft() // Clockwise relative to original direction
				collisions++
			}

			seen[ghosts[idx].Location] = append(seen[ghosts[idx].Location], idx)
		}

		for _, indices := range seen {
			if len(indices) > 1 {
				for _, idx := range indices {
					ghosts[idx].TurnRight()
					ghosts[idx].TurnRight()
					collisions++
				}
			}
		}
	}

	fmt.Println(collisions)
}

func InputToGridAndGhosts() (Grid2D[string], []Turtle) {
	var headings = map[string]Heading{"^": Up, "v": Down, ">": Right, "<": Left}

	var ghosts []Turtle
	grid := in.ToGrid2D(func(x, y int, s string) string {
		if s == "#" || s == "." {
			return s
		}

		ghosts = append(ghosts, Turtle{Location: Point2D{X: x, Y: y}, Heading: headings[s]})
		return "."
	})

	return grid, ghosts
}
