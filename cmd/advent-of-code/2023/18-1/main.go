package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	plan := InputToPlan()
	grid := PlanToGrid(plan)

	// Find the starting point for the flood fill.  Look for a point in the top
	// row that has an empty point below it.
	var start lib.Point2D
	for x := 0; x < grid.Width; x++ {
		if grid.Get(x, 0) == "#" && grid.Get(x, 1) == "." {
			start = lib.Point2D{X: x, Y: 1}
			break
		}
	}

	// Use a BFS to find the interior points of the polygon.
	var interior lib.Set[lib.Point2D]
	lib.BreadthFirstSearch(
		start,
		func(p lib.Point2D) []lib.Point2D {
			var children []lib.Point2D
			grid.ForEachOrthogonalNeighborPoint(p, func(q lib.Point2D, s string) {
				if s == "." {
					children = append(children, q)
				}
			})
			return children
		},
		func(p lib.Point2D) bool {
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

func PlanToGrid(plan []Step) lib.Grid2D[string] {
	var points lib.Set[lib.Point2D]
	var current lib.Point2D
	for _, step := range plan {
		var dx, dy int
		switch step.Heading {
		case lib.Up:
			dy = -1
		case lib.Right:
			dx = 1
		case lib.Down:
			dy = 1
		case lib.Left:
			dx = -1
		}

		for n := 0; n < step.Length; n++ {
			current = lib.Point2D{X: current.X + dx, Y: current.Y + dy}
			points.Add(current)
		}
	}

	tl, br := lib.GetBounds(points.Entries())

	grid := lib.NewGrid2D[string](br.X-tl.X+1, br.Y-tl.Y+1)
	grid.ForEach(func(x int, y int, _ string) {
		grid.Set(x, y, ".")
	})
	for p := range points {
		grid.Set(p.X-tl.X, p.Y-tl.Y, "#")
	}

	return grid
}

type Step struct {
	Heading lib.Heading
	Length  int
	Color   string
}

func InputToPlan() []Step {
	parseHeading := func(s string) lib.Heading {
		switch s {
		case "U":
			return lib.Up
		case "R":
			return lib.Right
		case "D":
			return lib.Down
		case "L":
			return lib.Left
		}
		return -1
	}

	return lib.InputLinesTo(func(line string) Step {
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
