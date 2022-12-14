package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	world, springX := InputToWorld()

	var count int
	for Down(world, springX, 0) {
		count++
	}
	fmt.Println(count)
}

func Down(world World, x, y int) bool {
	if world.Get(x, y) != Empty || y >= world.Height-1 {
		return false
	}

	downY := y + 1

	if world.Get(x, downY) == Empty {
		return Down(world, x, downY)
	}

	if left := x - 1; world.InBounds(left, downY) && world.Get(left, downY) == Empty {
		return Down(world, left, downY)
	}

	if right := x + 1; world.InBounds(right, downY) && world.Get(right, downY) == Empty {
		return Down(world, right, downY)
	}

	world.Add(x, y, Sand)
	return true
}

const (
	Empty = iota
	Wall
	Sand
)

type World struct {
	aoc.Grid2D[int]
}

func InputToWorld() (World, int) {
	type Line []aoc.Point2D
	lines := aoc.InputLinesTo(2022, 14, func(s string) (Line, error) {
		var line Line
		for _, field := range strings.Split(s, " -> ") {
			x, y, _ := strings.Cut(field, ",")
			line = append(line, aoc.Point2D{X: aoc.ParseInt(x), Y: aoc.ParseInt(y)})
		}

		return line, nil
	})

	// Determine the bounding box of the line segments.
	var ps []aoc.Point2D
	for _, line := range lines {
		ps = append(ps, line...)
	}
	ps = append(ps, aoc.Point2D{X: 500, Y: 0})
	tl, br := aoc.GetBounds(ps)

	// Expand our bounding box to the left and right.  Since when sand flows
	// it'll form the y=x line, we can expand by the height and know that
	// we've expanded wide enough.
	tl.X -= br.Y
	br.X += br.Y

	// Determine the offsets to apply to each line segment that removes empty
	// space to the left of them.
	x0, y0 := tl.X-1, tl.Y

	// Build the world from the line segments.
	world := aoc.NewGrid2D[int](br.X-tl.X+2, br.Y-tl.Y+3)
	for _, line := range lines {
		p := line[0]
		world.Add(p.X-x0, p.Y-y0, Wall)

		for _, q := range line[1:] {
			dx, dy := aoc.Sign(q.X-p.X), aoc.Sign(q.Y-p.Y)
			for p != q {
				world.Add(p.X-x0, p.Y-y0, Wall)
				p.X += dx
				p.Y += dy
			}
			world.Add(p.X-x0, p.Y-y0, Wall)
		}
	}

	for x := 0; x < world.Width; x++ {
		world.Add(x, world.Height-1, Wall)
	}

	// Lastly, since we've shifted the coordinates around also compute the new
	// X coordinate of the spring.  We know the y coordinate will be 0.
	springX := 500 - x0

	return World{world}, springX
}
