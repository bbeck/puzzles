package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	plan := InputToPlan()
	grid := PlanToGrid(plan)

	// Find the starting point for the flood fill.  Look for a point in the top
	// row that has an empty point below it.
	var start puz.Point2D
	for x := 0; x < grid.Width; x++ {
		if grid.Get(x, 0) == "#" && grid.Get(x, 1) == "." {
			start = puz.Point2D{X: x, Y: 1}
			break
		}
	}

	// Use a BFS to find the interior points of the polygon.
	var interior puz.Set[puz.Point2D]
	puz.BreadthFirstSearch(
		start,
		func(p puz.Point2D) []puz.Point2D {
			var children []puz.Point2D
			grid.ForEachOrthogonalNeighbor(p, func(q puz.Point2D, s string) {
				if s == "." {
					children = append(children, q)
				}
			})
			return children
		},
		func(p puz.Point2D) bool {
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

func PlanToGrid(plan []Step) puz.Grid2D[string] {
	var points puz.Set[puz.Point2D]
	var current puz.Point2D
	for _, step := range plan {
		var dx, dy int
		switch step.Heading {
		case puz.Up:
			dy = -1
		case puz.Right:
			dx = 1
		case puz.Down:
			dy = 1
		case puz.Left:
			dx = -1
		}

		for n := 0; n < step.Length; n++ {
			current = puz.Point2D{X: current.X + dx, Y: current.Y + dy}
			points.Add(current)
		}
	}

	tl, br := puz.GetBounds(points.Entries())

	grid := puz.NewGrid2D[string](br.X-tl.X+1, br.Y-tl.Y+1)
	grid.ForEach(func(x int, y int, _ string) {
		grid.Set(x, y, ".")
	})
	for p := range points {
		grid.Set(p.X-tl.X, p.Y-tl.Y, "#")
	}

	return grid
}

type Step struct {
	Heading puz.Heading
	Length  int
	Color   string
}

func InputToPlan() []Step {
	parseHeading := func(s string) puz.Heading {
		switch s {
		case "U":
			return puz.Up
		case "R":
			return puz.Right
		case "D":
			return puz.Down
		case "L":
			return puz.Left
		}
		return -1
	}

	return puz.InputLinesTo(func(line string) Step {
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
