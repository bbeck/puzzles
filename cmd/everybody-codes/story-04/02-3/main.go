package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	current, points := InputToPoints()

	// We'll use a breadth first search to explore the space of what would happen
	// if at every location we tried moving towards every point.
	var visited Set[Point2D]
	var queue = []Point2D{current}
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if !visited.Add(current) {
			continue
		}

		for _, q := range points {
			queue = append(queue, Point2D{
				X: (current.X + q.X) / 2,
				Y: (current.Y + q.Y) / 2,
			})
		}
	}

	var neighbors Set[Point2D]
	for v := range visited {
		neighbors.Add(v.OrthogonalNeighbors()...)
	}
	neighbors = neighbors.Difference(visited)
	fmt.Println(len(neighbors))
}

func InputToPoints() (Point2D, []Point2D) {
	var start Point2D
	var points []Point2D

	for in.HasNext() {
		lhs, rhs := in.CutS[any]("=")
		if name := lhs.String(); name != "START" {
			points = append(points, Point2D{X: rhs.Int(), Y: rhs.Int()})
		} else {
			start = Point2D{X: rhs.Int(), Y: rhs.Int()}
		}
	}

	return start, points
}
