package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ps := InputToPoints()

	var ds aoc.DisjointSet[Point4D]
	for i := 0; i < len(ps); i++ {
		ds.Add(ps[i])

		for j := i + 1; j < len(ps); j++ {
			ds.Add(ps[j])

			if ps[i].ManhattanDistance(ps[j]) <= 3 {
				ds.Union(ps[i], ps[j])
			}
		}
	}

	var constellations aoc.Set[Point4D]
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
	return aoc.Abs(q.W-p.W) + aoc.Abs(q.X-p.X) + aoc.Abs(q.Y-p.Y) + aoc.Abs(q.Z-p.Z)
}

func InputToPoints() []Point4D {
	return aoc.InputLinesTo(2018, 25, func(line string) (Point4D, error) {
		var p Point4D
		_, err := fmt.Sscanf(line, "%d,%d,%d,%d", &p.W, &p.X, &p.Y, &p.Z)
		return p, err
	})
}
