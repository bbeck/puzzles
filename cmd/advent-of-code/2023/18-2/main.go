package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	plan := InputToPlan()
	vertices := PlanToVertices(plan)

	// Use the shoelace formula to compute the interior area
	// https://en.wikipedia.org/wiki/Shoelace_formula#Shoelace_formula
	var interior int
	for i := 0; i < len(vertices); i++ {
		j := (i + 1) % len(vertices)
		interior += vertices[i].X * vertices[j].Y
		interior -= vertices[j].X * vertices[i].Y
	}
	interior /= 2

	// The interior area from the shoelace formula is missing 1/2 of each
	// perimeter block so also compute half of the perimeter area.  In addition,
	// because the polygon is defined as a series of steps, the starting block
	// isn't taken into account for in the perimeter, so make sure to add it back
	// in as well.
	var perimeter int
	for _, step := range plan {
		perimeter += step.Length
	}
	perimeter = (perimeter / 2) + 1

	fmt.Println(interior + perimeter)
}

func PlanToVertices(plan []Step) []puz.Point2D {
	var current puz.Point2D

	ps := []puz.Point2D{current}
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

		current = puz.Point2D{
			X: current.X + dx*step.Length,
			Y: current.Y + dy*step.Length,
		}
		ps = append(ps, current)
	}
	return ps
}

type Step struct {
	Heading puz.Heading
	Length  int
}

func InputToPlan() []Step {
	parseHeading := func(s string) puz.Heading {
		switch s[len(s)-1] {
		case '0':
			return puz.Right
		case '1':
			return puz.Down
		case '2':
			return puz.Left
		case '3':
			return puz.Up
		}
		return -1
	}

	parseLength := func(s string) int {
		return puz.ParseIntWithBase(s[:5], 16)
	}

	return puz.InputLinesTo(func(line string) Step {
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		line = strings.ReplaceAll(line, "#", "")
		fields := strings.Fields(line)

		return Step{
			Heading: parseHeading(fields[2]),
			Length:  parseLength(fields[2]),
		}
	})
}
