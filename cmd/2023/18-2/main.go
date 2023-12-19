package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
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

func PlanToVertices(plan []Step) []aoc.Point2D {
	var current aoc.Point2D

	ps := []aoc.Point2D{current}
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

		current = aoc.Point2D{
			X: current.X + dx*step.Length,
			Y: current.Y + dy*step.Length,
		}
		ps = append(ps, current)
	}
	return ps
}

type Step struct {
	Heading aoc.Heading
	Length  int
}

func InputToPlan() []Step {
	parseHeading := func(s string) aoc.Heading {
		switch s[len(s)-1] {
		case '0':
			return aoc.Right
		case '1':
			return aoc.Down
		case '2':
			return aoc.Left
		case '3':
			return aoc.Up
		}
		return -1
	}

	parseLength := func(s string) int {
		return aoc.ParseIntWithBase(s[:5], 16)
	}

	return aoc.InputLinesTo(2023, 18, func(line string) Step {
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
