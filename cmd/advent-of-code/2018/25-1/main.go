package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	ps := InputToPoints()

	var ds lib.DisjointSet[Point4D]
	for i := 0; i < len(ps); i++ {
		ds.Add(ps[i])

		for j := i + 1; j < len(ps); j++ {
			ds.Add(ps[j])

			if ps[i].ManhattanDistance(ps[j]) <= 3 {
				ds.Union(ps[i], ps[j])
			}
		}
	}

	var constellations lib.Set[Point4D]
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
	return lib.Abs(q.W-p.W) + lib.Abs(q.X-p.X) + lib.Abs(q.Y-p.Y) + lib.Abs(q.Z-p.Z)
}

func InputToPoints() []Point4D {
	return lib.InputLinesTo(func(line string) Point4D {
		var p Point4D
		fmt.Sscanf(line, "%d,%d,%d,%d", &p.W, &p.X, &p.Y, &p.Z)
		return p
	})
}
