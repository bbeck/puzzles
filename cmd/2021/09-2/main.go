package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"sort"
)

func main() {
	m := InputToHeightMap()

	neighbors := func(p aoc.Point2D) []aoc.Point2D {
		var ns []aoc.Point2D
		for _, n := range []aoc.Point2D{p.Up(), p.Right(), p.Down(), p.Left()} {
			if _, ok := m[n]; ok {
				ns = append(ns, n)
			}
		}
		return ns
	}

	// Build a single element disjoint set for each point
	sets := make(map[aoc.Point2D]*aoc.DisjointSet)
	for p := range m {
		sets[p] = aoc.NewDisjointSet(p)
	}

	// Link a point's disjoint set to it's neighbor's if the neighbor is higher
	for p, h := range m {
		for _, neighbor := range neighbors(p) {
			if m[neighbor] > h {
				sets[p].Union(sets[neighbor])
			}
		}
	}

	// Now determine the sizes of each set and sort them to find the largest three.
	seen := aoc.NewSet()
	var sizes []int
	for _, set := range sets {
		parent := set.Find()
		if seen.Add(parent) {
			sizes = append(sizes, parent.Size)
		}
	}
	sort.Ints(sizes)

	N := len(sizes)
	fmt.Println(sizes[N-1] * sizes[N-2] * sizes[N-3])
}

func InputToHeightMap() map[aoc.Point2D]int {
	m := make(map[aoc.Point2D]int)
	for y, line := range aoc.InputToLines(2021, 9) {
		for x, c := range line {
			n := aoc.ParseInt(string(c))

			// Locations with a height of 9 aren't considered part of any basin
			if n == 9 {
				continue
			}

			m[aoc.Point2D{X: x, Y: y}] = n
		}
	}

	return m
}
