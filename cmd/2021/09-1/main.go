package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	m := InputToHeightMap()

	var risk int
	for p, n := range m {
		if u, ok := m[p.Up()]; ok && n >= u {
			continue
		}
		if u, ok := m[p.Down()]; ok && n >= u {
			continue
		}
		if u, ok := m[p.Left()]; ok && n >= u {
			continue
		}
		if u, ok := m[p.Right()]; ok && n >= u {
			continue
		}
		risk += n + 1
	}

	fmt.Println(risk)
}

func InputToHeightMap() map[aoc.Point2D]int {
	m := make(map[aoc.Point2D]int)
	for y, line := range aoc.InputToLines(2021, 9) {
		for x, c := range line {
			n := aoc.ParseInt(string(c))
			m[aoc.Point2D{X: x, Y: y}] = n
		}
	}

	return m
}
