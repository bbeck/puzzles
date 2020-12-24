package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := make(map[HexPoint]bool)
	for _, p := range InputToCoordinates(2020, 24) {
		grid[p] = !grid[p]
	}

	for day := 1; day <= 100; day++ {
		grid = Next(grid)
	}

	var sum int
	for _, value := range grid {
		if value {
			sum++
		}
	}
	fmt.Println(sum)
}

func Next(grid map[HexPoint]bool) map[HexPoint]bool {
	var maxDistance int
	for p := range grid {
		distance := p.Distance(Origin)
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	next := make(map[HexPoint]bool)
	for radius := 0; radius <= maxDistance+1; radius++ {
		for _, p := range Spiral(Origin, radius) {
			// Count the number of neighbors that have their black side facing up.
			var neighbors int
			for _, q := range p.Neighbors() {
				if grid[q] {
					neighbors++
				}
			}

			result := grid[p]
			if grid[p] && (neighbors == 0 || neighbors > 2) {
				result = false
			}
			if !grid[p] && neighbors == 2 {
				next[p] = true
			}

			// Only save the tiles that have their black side facing up.
			if result {
				next[p] = result
			}
		}
	}

	return next
}

// Compute the points in a ring with the given center and radius.
func Ring(p HexPoint, radius int) []HexPoint {
	if radius == 0 {
		return []HexPoint{Origin}
	}

	// We'll start at a point on the ring, and then follow a path around the
	// center.  To generate the path we'll take radius steps in each direction.
	// The order of the directions is important, they need to be ordered so that
	// we keep moving around the center.
	directions := []HexPoint{
		{X: 1, Y: -1, Z: 0},
		{X: 1, Y: 0, Z: -1},
		{X: 0, Y: 1, Z: -1},
		{X: -1, Y: 1, Z: 0},
		{X: -1, Y: 0, Z: 1},
		{X: 0, Y: -1, Z: 1},
	}

	// Move to a point on the ring, since the first direction we'll move is east,
	// the point on the ring we'll choose is the most southwest one.
	p.X += radius * directions[4].X
	p.Y += radius * directions[4].Y
	p.Z += radius * directions[4].Z

	var ring []HexPoint
	for _, direction := range directions {
		for r := 0; r < radius; r++ {
			ring = append(ring, p)

			// Move to the next point in this direction
			p.X += direction.X
			p.Y += direction.Y
			p.Z += direction.Z
		}
	}

	return ring
}

// Compute a spiral of points starting at the given center and moving outwards
// until all points within the specified radius are produced.
func Spiral(center HexPoint, radius int) []HexPoint {
	var spiral []HexPoint
	for r := 0; r <= radius; r++ {
		spiral = append(spiral, Ring(center, r)...)
	}

	return spiral
}

func InputToCoordinates(year, day int) []HexPoint {
	var coordinates []HexPoint
	for _, line := range aoc.InputToLines(year, day) {
		var p HexPoint

		for i := 0; i < len(line); {
			switch {
			case line[i:i+1] == "e":
				p = p.East()
				i += 1
			case i+2 <= len(line) && line[i:i+2] == "se":
				p = p.SouthEast()
				i += 2
			case i+2 <= len(line) && line[i:i+2] == "sw":
				p = p.SouthWest()
				i += 2
			case line[i:i+1] == "w":
				p = p.West()
				i += 1
			case i+2 <= len(line) && line[i:i+2] == "nw":
				p = p.NorthWest()
				i += 2
			case i+2 <= len(line) && line[i:i+2] == "ne":
				p = p.NorthEast()
				i += 2
			}
		}

		coordinates = append(coordinates, p)
	}

	return coordinates
}

type HexPoint struct {
	X, Y, Z int
}

var Origin = HexPoint{}

func (p HexPoint) East() HexPoint {
	return HexPoint{X: p.X + 1, Y: p.Y - 1, Z: p.Z}
}

func (p HexPoint) SouthEast() HexPoint {
	return HexPoint{X: p.X, Y: p.Y - 1, Z: p.Z + 1}
}

func (p HexPoint) SouthWest() HexPoint {
	return HexPoint{X: p.X - 1, Y: p.Y, Z: p.Z + 1}
}

func (p HexPoint) West() HexPoint {
	return HexPoint{X: p.X - 1, Y: p.Y + 1, Z: p.Z}
}

func (p HexPoint) NorthWest() HexPoint {
	return HexPoint{X: p.X, Y: p.Y + 1, Z: p.Z - 1}
}

func (p HexPoint) NorthEast() HexPoint {
	return HexPoint{X: p.X + 1, Y: p.Y, Z: p.Z - 1}
}

func (p HexPoint) Distance(q HexPoint) int {
	dx := aoc.MaxInt(p.X, q.X) - aoc.MinInt(p.X, q.X)
	dy := aoc.MaxInt(p.Y, q.Y) - aoc.MinInt(p.Y, q.Y)
	dz := aoc.MaxInt(p.Z, q.Z) - aoc.MinInt(p.Z, q.Z)
	return (dx + dy + dz) / 2
}

func (p HexPoint) Neighbors() []HexPoint {
	return []HexPoint{
		p.East(),
		p.SouthEast(),
		p.SouthWest(),
		p.West(),
		p.NorthWest(),
		p.NorthEast(),
	}
}
