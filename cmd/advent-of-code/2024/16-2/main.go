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

	paths, _ := AllShortestPaths(
		start,
		func(t Turtle) []Turtle { return Children(grid, t) },
		func(t Turtle) bool { return t.Location == end },
		func(u, v Turtle) int {
			if u.Heading == v.Heading {
				return 1
			}
			return 1001
		},
	)

	var seen Set[Point2D]
	for _, p := range paths {
		for _, t := range p {
			seen.Add(t.Location)
		}
	}
	fmt.Println(len(seen))
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
