package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	ps := InputToPoints()

	var ds DisjointSet[Point4D]
	for i := range ps {
		ds.Add(ps[i])

		for j := i + 1; j < len(ps); j++ {
			ds.Add(ps[j])

			if ps[i].ManhattanDistance(ps[j]) <= 3 {
				ds.Union(ps[i], ps[j])
			}
		}
	}

	var constellations Set[Point4D]
	for _, p := range ps {
		c, _ := ds.Find(p)
		constellations.Add(c)
	}
	fmt.Println(len(constellations))
}

type Point4D struct {
	W, X, Y, Z int
}

func (p Point4D) ManhattanDistance(q Point4D) int {
	return Abs(q.W-p.W) + Abs(q.X-p.X) + Abs(q.Y-p.Y) + Abs(q.Z-p.Z)
}

func InputToPoints() []Point4D {
	return in.LinesToS(func(in.Scanner[Point4D]) Point4D {
		return Point4D{W: in.Int(), X: in.Int(), Y: in.Int(), Z: in.Int()}
	})
}
