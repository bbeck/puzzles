package main

import (
	"fmt"
	"math"
	"sort"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	gaps, width, height := InputToGaps()

	var xs = []int{0}
	for x := range gaps {
		xs = append(xs, x)
	}
	sort.Ints(xs)

	var grid = NewGrid2D[string](width, height)
	grid.Fill(".")
	for x, gs := range gaps {
		for y := 0; y < height; y++ {
			grid.Set(x, y, "#")
		}

		for _, gap := range gs {
			for y := gap.Start; y <= gap.End; y++ {
				grid.Set(x, y, ".")
			}
		}
	}

	children := func(p Point2D) []Point2D {
		var children []Point2D

		p1 := p.Right().Up()
		if grid.InBoundsPoint(p1) && grid.GetPoint(p1) == "." {
			children = append(children, p1)
		}
		p2 := p.Right().Down()
		if grid.InBoundsPoint(p2) && grid.GetPoint(p2) == "." {
			children = append(children, p2)
		}

		return children
	}

	cost := func(from, to Point2D) int {
		if from.Y < to.Y {
			return 1
		}
		return 0
	}

	costs, paths := Dijkstra(Origin2D, children, cost)
	x := xs[len(xs)-1]

	var best = math.MaxInt
	for _, gap := range gaps[x] {
		for y := gap.Start; y <= gap.End; y++ {
			p := Point2D{X: x, Y: y}
			if len(paths[p]) > 0 {
				best = Min(best, costs[p])
			}
		}
	}
	fmt.Println(best)
}

type Gap struct {
	Start, End int // Inclusive
}

func InputToGaps() (map[int][]Gap, int, int) {
	var gaps = make(map[int][]Gap)
	var maxX, maxY int
	for in.HasNext() {
		x, y0, height := in.Int(), in.Int(), in.Int()
		gaps[x] = append(gaps[x], Gap{Start: y0, End: y0 + height - 1})
		maxX = Max(maxX, x)
		maxY = Max(maxY, y0+height-1)
	}

	return gaps, maxX + 1, maxY + 1
}
