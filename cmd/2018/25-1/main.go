package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var sets []*aoc.DisjointSet
	for _, point := range InputToPoints(2018, 25) {
		sets = append(sets, aoc.NewDisjointSet(point))
	}

	for i := 0; i < len(sets); i++ {
		a := sets[i].Data.(Point4D)
		for j := i + 1; j < len(sets); j++ {
			b := sets[j].Data.(Point4D)
			if a.Distance(b) <= 3 {
				sets[i].Union(sets[j])
			}
		}
	}

	constellations := make(map[Point4D]bool)
	for _, set := range sets {
		constellations[set.Find().Data.(Point4D)] = true
	}

	fmt.Printf("num constellations: %d\n", len(constellations))
}

type Point4D struct {
	w, x, y, z int
}

func (p Point4D) Distance(other Point4D) int {
	dw := p.w - other.w
	if dw < 0 {
		dw = -dw
	}

	dx := p.x - other.x
	if dx < 0 {
		dx = -dx
	}

	dy := p.y - other.y
	if dy < 0 {
		dy = -dy
	}

	dz := p.z - other.z
	if dz < 0 {
		dz = -dz
	}

	return dw + dx + dy + dz
}

func InputToPoints(year, day int) []Point4D {
	var points []Point4D
	for _, line := range aoc.InputToLines(year, day) {
		var w, x, y, z int
		if _, err := fmt.Sscanf(line, "%d,%d,%d,%d", &w, &x, &y, &z); err != nil {
			log.Fatalf("unble to parse line: %s", line)
		}

		points = append(points, Point4D{w, x, y, z})
	}

	return points
}
