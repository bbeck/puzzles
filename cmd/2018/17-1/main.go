package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	world := InputToWorld()

	// Find the initial spring of water.
	var springX int
	for x := 0; x < world.Width; x++ {
		if world.Get(x, 0) == Flow {
			springX = x
		}
	}

	Down(world, springX, 1)

	var wet int
	world.ForEach(func(x, y int, value int) {
		if value == Flow || value == Water {
			wet++
		}
	})
	fmt.Println(wet)
}

func Down(world World, x, y int) {
	if world.Get(x, y) != Empty {
		return
	}

	// Start by making our new point water flow.
	world.Add(x, y, Flow)

	downY := y + 1
	if !world.InBounds(x, downY) {
		return
	}

	// If below us is empty, then flow down into it.
	if world.Get(x, downY) == Empty {
		Down(world, x, downY)
	}

	// At this point everything below us has been determined.  If we're
	// sitting on top of water or a wall, then we should spread out to the
	// sides.
	if below := world.Get(x, downY); below == Wall || below == Water {
		cl := Side(world, x, y, -1)
		cr := Side(world, x, y, +1)

		// If the flow at this level is contained on both sides by a wall, then
		// convert it into standing water.
		if cl && cr {
			for q := x; world.InBounds(q, y) && world.Get(q, y) != Wall; q-- {
				world.Add(q, y, Water)
			}
			for q := x; world.InBounds(q, y) && world.Get(q, y) != Wall; q++ {
				world.Add(q, y, Water)
			}
		}
	}
}

func Side(world World, x, y, dx int) bool {
	if current := world.Get(x, y); current == Wall || current == Water {
		return true
	}

	below := world.Get(x, y+1)
	switch {
	case below == Wall || below == Water:
		world.Add(x, y, Flow)
		return Side(world, x+dx, y, dx)
	case below == Empty:
		Down(world, x, y)
	}

	return false
}

const (
	Empty = iota
	Wall
	Water
	Flow
)

type World struct{ aoc.Grid2D[int] }

func InputToWorld() World {
	// Convert the input into line segments
	type Line []aoc.Point2D
	lines := aoc.InputLinesTo(2018, 17, func(s string) (Line, error) {
		var x1, x2, y1, y2 int
		var line Line
		if _, err := fmt.Sscanf(s, "x=%d, y=%d..%d", &x1, &y1, &y2); err == nil {
			for y := y1; y <= y2; y++ {
				line = append(line, aoc.Point2D{X: x1, Y: y})
			}
			return line, nil
		}
		if _, err := fmt.Sscanf(s, "y=%d, x=%d..%d", &y1, &x1, &x2); err == nil {
			for x := x1; x <= x2; x++ {
				line = append(line, aoc.Point2D{X: x, Y: y1})
			}
			return line, nil
		}
		return line, fmt.Errorf("unable to parse line: %s", s)
	})

	// Determine the bounding box of the line segments.
	var ps []aoc.Point2D
	for _, line := range lines {
		ps = append(ps, line...)
	}
	tl, br := aoc.GetBounds(ps)

	// Determine the offsets to apply to each line segment that removes empty
	// space to the left of them.
	x0, y0 := tl.X-1, tl.Y

	// Build the world from the line segments.
	world := aoc.NewGrid2D[int](br.X-tl.X+2, br.Y-tl.Y+1)
	for _, line := range lines {
		for _, p := range line {
			world.Add(p.X-x0, p.Y-y0, Wall)
		}
	}

	// Lastly, since we've shifted the coordinates around also compute the new
	// X coordinate of the spring.  We know the y coordinate will be 0.
	world.Add(500-x0, 0, Flow)

	return World{world}
}
