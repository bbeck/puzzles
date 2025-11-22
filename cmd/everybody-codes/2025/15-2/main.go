package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var walls Set[Point2D]

	var turtle Turtle
	for _, part := range in.Split(",") {
		dir, n := part[0], ParseInt(part[1:])
		if dir == 'L' {
			turtle.TurnLeft()
		}
		if dir == 'R' {
			turtle.TurnRight()
		}
		for range n {
			turtle.Forward(1)
			walls.Add(turtle.Location)
		}
	}

	// The end is where the turtle is currently at, remove it as a wall.
	start, end := Origin2D, turtle.Location
	walls.Remove(end)

	children := func(p Point2D) []Point2D {
		var children []Point2D
		for _, q := range p.OrthogonalNeighbors() {
			if !walls.Contains(q) {
				children = append(children, q)
			}
		}

		return children
	}

	goal := func(p Point2D) bool {
		return p == end
	}

	path, _ := BreadthFirstSearch(start, children, goal)
	fmt.Println(len(path) - 1)
}
