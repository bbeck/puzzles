package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	ps := InputToPoints()

	var ds puz.DisjointSet[Point4D]
	for i := 0; i < len(ps); i++ {
		ds.Add(ps[i])

		for j := i + 1; j < len(ps); j++ {
			ds.Add(ps[j])

			if ps[i].ManhattanDistance(ps[j]) <= 3 {
				ds.Union(ps[i], ps[j])
			}
		}
	}

	var constellations puz.Set[Point4D]
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
	return puz.Abs(q.W-p.W) + puz.Abs(q.X-p.X) + puz.Abs(q.Y-p.Y) + puz.Abs(q.Z-p.Z)
}

func InputToPoints() []Point4D {
	return puz.InputLinesTo(2018, 25, func(line string) Point4D {
		var p Point4D
		fmt.Sscanf(line, "%d,%d,%d,%d", &p.W, &p.X, &p.Y, &p.Z)
		return p
	})
}
