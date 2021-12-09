package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	m := InputToHeightMap()

	isLowPoint := func(p aoc.Point2D) bool {
		neighbors := []aoc.Point2D{p.Up(), p.Right(), p.Down(), p.Left()}
		for _, n := range neighbors {
			if nh, ok := m[n]; ok && nh <= m[p] {
				return false
			}
		}

		return true
	}

	var risk int
	for p := range m {
		if isLowPoint(p) {
			risk += m[p] + 1
		}
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
