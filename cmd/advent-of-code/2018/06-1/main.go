package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"math"
)

func main() {
	points := InputToPoints()
	tl, br := puz.GetBounds(points)

	// This grid contains the index of the closest point, or -1 if no single
	// point is closest to this location.  The only portions of the grid that
	// have values are the ones within the bounds of our points.
	grid := puz.NewGrid2D[int](br.X+1, br.Y+1)
	for y := tl.Y; y <= br.Y; y++ {
		for x := tl.X; x <= br.X; x++ {
			cell := puz.Point2D{X: x, Y: y}
			grid.SetPoint(cell, -1) // initialize

			var best = math.MaxInt // how far away the closest point is
			var closest []int      // index of points that are closest
			for index, p := range points {
				d := cell.ManhattanDistance(p)

				if d < best {
					best = d
					closest = nil
				}

				if d == best {
					closest = append(closest, index)
				}
			}

			if len(closest) == 1 {
				grid.SetPoint(cell, closest[0])
			}
		}
	}

	// An area is infinite if it touches the edge
	var infinite puz.Set[int]
	for x := tl.X; x <= br.X; x++ {
		infinite.Add(grid.Get(x, tl.Y), grid.Get(x, br.Y))
	}
	for y := tl.Y; y <= br.Y; y++ {
		infinite.Add(grid.Get(tl.X, y), grid.Get(br.X, y))
	}

	var largest int
	sizes := make(map[int]int)
	for y := tl.Y; y <= br.Y; y++ {
		for x := tl.X; x <= br.X; x++ {
			n := grid.Get(x, y)
			if !infinite.Contains(n) {
				sizes[n]++
				largest = puz.Max(largest, sizes[n])
			}
		}
	}
	fmt.Println(largest)
}

func InputToPoints() []puz.Point2D {
	return puz.InputLinesTo(func(line string) puz.Point2D {
		var p puz.Point2D
		fmt.Sscanf(line, "%d, %d", &p.X, &p.Y)
		return p
	})
}
