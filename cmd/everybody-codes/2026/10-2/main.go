package main

import (
	"fmt"
	"slices"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid, ghosts := InputToGridAndGhosts()

	var round int
	for round = 1; ; round++ {
		var seen = make(map[Point2D][]int)
		for idx := range ghosts {
			ghosts[idx].Forward(1)
			if grid.GetPoint(ghosts[idx].Location) == "#" {
				ghosts[idx].TurnAround()
				ghosts[idx].Forward(1)
				ghosts[idx].TurnLeft() // Clockwise relative to original direction
				ghosts[idx].Collisions++
			}

			seen[ghosts[idx].Location] = append(seen[ghosts[idx].Location], idx)
		}

		for _, indices := range seen {
			if len(indices) > 1 {
				for _, idx := range indices {
					ghosts[idx].TurnAround()
					ghosts[idx].Collisions++
				}
			}
		}

		ghosts = slices.DeleteFunc(ghosts, func(ghost Ghost) bool {
			return ghost.Collisions >= 10000
		})

		if len(ghosts) == 0 {
			break
		}
	}

	fmt.Println(round)
}

type Ghost struct {
	Turtle
	Collisions int
}

func (g *Ghost) TurnAround() {
	g.TurnRight()
	g.TurnRight()
}

func InputToGridAndGhosts() (Grid2D[string], []Ghost) {
	var headings = map[string]Heading{"^": Up, "v": Down, ">": Right, "<": Left}

	var ghosts []Ghost
	grid := in.ToGrid2D(func(x, y int, s string) string {
		if s == "#" || s == "." {
			return s
		}

		ghosts = append(ghosts, Ghost{Location: Point2D{X: x, Y: y}, Heading: headings[s]})
		return "."
	})

	return grid, ghosts
}
