package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	world := InputToWorld()

	var count int
	for Pour(world) {
		count++
	}
	fmt.Println(count)
}

func Pour(world World) bool {
	p := aoc.Point2D{X: 500, Y: 0}

	for p.Y < world.Height-1 {
		if q := p.Down(); world.InBoundsPoint(q) && world.GetPoint(q) == Empty {
			p = q
			continue
		}
		if q := p.Down().Left(); world.InBoundsPoint(q) && world.GetPoint(q) == Empty {
			p = q
			continue
		}
		if q := p.Down().Right(); world.InBoundsPoint(q) && world.GetPoint(q) == Empty {
			p = q
			continue
		}
		break
	}

	if world.GetPoint(p) == Empty {
		world.SetPoint(p, Sand)
		return true
	}
	return false
}

const (
	Empty = iota
	Wall
	Sand
)

type World struct {
	aoc.Grid2D[int]
}

func InputToWorld() World {
	var walls aoc.Set[aoc.Point2D]
	for _, line := range aoc.InputToLines(2022, 14) {
		points := strings.Split(line, " -> ")

		current := ParsePoint(points[0])
		walls.Add(current)

		for _, s := range points[1:] {
			end := ParsePoint(s)

			dx, dy := aoc.Sign(end.X-current.X), aoc.Sign(end.Y-current.Y)
			for current != end {
				current.X += dx
				current.Y += dy
				walls.Add(current)
			}
		}
	}

	// Add the floor at the bottom of our world.  It's supposed to expand
	// infinitely to the left and right, but we know that in the worst case the
	// flow of sand goes along the line y=x, so we only need to go as wide as we
	// are high.
	tl, br := aoc.GetBounds(walls.Entries())
	for x := tl.X - br.Y; x <= br.X+br.Y; x++ {
		walls.Add(aoc.Point2D{X: x, Y: br.Y + 2})
	}

	_, br = aoc.GetBounds(walls.Entries())
	grid := aoc.NewGrid2D[int](br.X+1, br.Y+1)
	for p := range walls {
		grid.SetPoint(p, Wall)
	}

	return World{grid}
}

func ParsePoint(s string) aoc.Point2D {
	var p aoc.Point2D
	_, _ = fmt.Sscanf(s, "%d,%d", &p.X, &p.Y)
	return p
}
