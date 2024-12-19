package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
)

func main() {
	grid := InputToStringGrid2D()

	var start Turtle
	var end Point2D
	grid.ForEachPoint(func(p Point2D, s string) {
		switch s {
		case "S":
			start = Turtle{Location: p, Heading: Right}
		case "E":
			end = p
		}
	})

	_, cost, _ := AStarSearch(
		start,
		func(t Turtle) []Turtle { return Children(grid, t) },
		func(t Turtle) bool { return t.Location == end },
		func(from, to Turtle) int {
			if from.Heading == to.Heading {
				return 1
			}
			return 1001
		},
		func(t Turtle) int { return t.Location.ManhattanDistance(end) },
	)

	fmt.Println(cost)
}

func Children(grid Grid2D[string], t Turtle) []Turtle {
	var children []Turtle

	// forward
	{
		u := t
		u.Forward(1)
		if grid.GetPoint(u.Location) != "#" {
			children = append(children, u)
		}
	}

	// right
	{
		u := t
		u.TurnRight()
		u.Forward(1)
		if grid.GetPoint(u.Location) != "#" {
			children = append(children, u)
		}
	}

	// forward
	{
		u := t
		u.TurnLeft()
		u.Forward(1)
		if grid.GetPoint(u.Location) != "#" {
			children = append(children, u)
		}
	}

	return children
}
