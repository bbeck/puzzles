package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	plan := InputToPlan()
	grid := PlanToGrid(plan)

	// Find the starting point for the flood fill.  Look for a point in the top
	// row that has an empty point below it.
	var start aoc.Point2D
	for x := 0; x < grid.Width; x++ {
		if grid.Get(x, 0) == "#" && grid.Get(x, 1) == "." {
			start = aoc.Point2D{X: x, Y: 1}
			break
		}
	}

	// Use a BFS to find the interior points of the polygon.
	var interior aoc.Set[aoc.Point2D]
	aoc.BreadthFirstSearch(
		start,
		func(p aoc.Point2D) []aoc.Point2D {
			var children []aoc.Point2D
			grid.ForEachOrthogonalNeighbor(p, func(q aoc.Point2D, s string) {
				if s == "." {
					children = append(children, q)
				}
			})
			return children
		},
		func(p aoc.Point2D) bool {
			interior.Add(p)
			return false
		},
	)

	// Determine the perimeter of the polygon.
	var perimeter int
	grid.ForEach(func(x int, y int, s string) {
		if s == "#" {
			perimeter++
		}
	})

	// The area is the interior plus the perimeter
	fmt.Println(len(interior) + perimeter)
}

func PlanToGrid(plan []Step) aoc.Grid2D[string] {
	var points aoc.Set[aoc.Point2D]
	var current aoc.Point2D
	for _, step := range plan {
		var dx, dy int
		switch step.Heading {
		case aoc.Up:
			dy = -1
		case aoc.Right:
			dx = 1
		case aoc.Down:
			dy = 1
		case aoc.Left:
			dx = -1
		}

		for n := 0; n < step.Length; n++ {
			current = aoc.Point2D{X: current.X + dx, Y: current.Y + dy}
			points.Add(current)
		}
	}

	tl, br := aoc.GetBounds(points.Entries())

	grid := aoc.NewGrid2D[string](br.X-tl.X+1, br.Y-tl.Y+1)
	grid.ForEach(func(x int, y int, _ string) {
		grid.Set(x, y, ".")
	})
	for p := range points {
		grid.Set(p.X-tl.X, p.Y-tl.Y, "#")
	}

	return grid
}

type Step struct {
	Heading aoc.Heading
	Length  int
	Color   string
}

func InputToPlan() []Step {
	parseHeading := func(s string) aoc.Heading {
		switch s {
		case "U":
			return aoc.Up
		case "R":
			return aoc.Right
		case "D":
			return aoc.Down
		case "L":
			return aoc.Left
		}
		return -1
	}

	return aoc.InputLinesTo(2023, 18, func(line string) Step {
		var heading, color string
		var length int

		fmt.Sscanf(line, "%s %d %s", &heading, &length, &color)
		return Step{
			Heading: parseHeading(heading),
			Length:  length,
			Color:   color,
		}
	})
}
