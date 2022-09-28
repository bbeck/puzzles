package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	world := InputToWorld()

	// Find the initial spring of water.
	var spring aoc.Point2D
	for x := 0; x < world.Width; x++ {
		if world.GetXY(x, 0) == Flow {
			spring = aoc.Point2D{X: x, Y: 0}
		}
	}

	Down(world, spring.Down())

	var wet int
	for y := 0; y < world.Height; y++ {
		for x := 0; x < world.Width; x++ {
			if current := world.GetXY(x, y); current == Flow || current == Water {
				wet++
			}
		}
	}
	fmt.Println(wet)
}

func Down(world World, p aoc.Point2D) {
	if world.Get(p) != Empty {
		return
	}

	// Start by making our new point water flow.
	world.Add(p, Flow)

	down := p.Down()
	if !world.InBounds(down) {
		return
	}

	// If below us is empty, then flow down into it.
	if world.Get(down) == Empty {
		Down(world, down)
	}

	// At this point everything below us has been determined.  If we're
	// sitting on top of water or a wall, then we should spread out to the
	// sides.
	if below := world.Get(down); below == Wall || below == Water {
		cl := Side(world, p, -1)
		cr := Side(world, p, +1)

		// If the flow at this level is contained on both sides by a wall, then
		// convert it into standing water.
		if cl && cr {
			for q := p; world.InBounds(q) && world.Get(q) != Wall; q = q.Left() {
				world.Add(q, Water)
			}
			for q := p; world.InBounds(q) && world.Get(q) != Wall; q = q.Right() {
				world.Add(q, Water)
			}
		}
	}
}

func Side(world World, p aoc.Point2D, dx int) bool {
	if current := world.Get(p); current == Wall || current == Water {
		return true
	}

	below := world.Get(p.Down())
	switch {
	case below == Wall || below == Water:
		world.Add(p, Flow)
		return Side(world, aoc.Point2D{X: p.X + dx, Y: p.Y}, dx)
	case below == Empty:
		Down(world, p)
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
			world.AddXY(p.X-x0, p.Y-y0, Wall)
		}
	}

	// Lastly, since we've shifted the coordinates around also compute the new
	// X coordinate of the spring.  We know the y coordinate will be 0.
	world.AddXY(500-x0, 0, Flow)

	return World{world}
}
