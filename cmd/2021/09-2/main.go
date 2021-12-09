package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"sort"
)

func main() {
	m := InputToHeightMap()

	sets := make(map[aoc.Point2D]*aoc.DisjointSet)
	for p := range m {
		sets[p] = aoc.NewDisjointSet(p)
	}

	for p, n := range m {
		if n == 9 {
			continue
		}

		if v, ok := m[p.Up()]; ok && v != 9 && v > n {
			sets[p].Union(sets[p.Up()])
		}
		if v, ok := m[p.Down()]; ok && v != 9 && v > n {
			sets[p].Union(sets[p.Down()])
		}
		if v, ok := m[p.Left()]; ok && v != 9 && v > n {
			sets[p].Union(sets[p.Left()])
		}
		if v, ok := m[p.Right()]; ok && v != 9 && v > n {
			sets[p].Union(sets[p.Right()])
		}
	}

	sizes := make(map[*aoc.DisjointSet]int)
	for _, set := range sets {
		sizes[set.Find()]++
	}

	var ordering []int
	for _, n := range sizes {
		ordering = append(ordering, n)
	}

	sort.Ints(ordering)
	N := len(ordering)
	fmt.Println(ordering[N-1] * ordering[N-2] * ordering[N-3])
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
