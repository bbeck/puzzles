package main

import (
	"fmt"
	"sort"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	stars := InputToStars()

	// Build the constellation groups.
	var ds DisjointSet[Point2D]
	for i := 0; i < len(stars); i++ {
		for j := i + 1; j < len(stars); j++ {
			if stars[i].ManhattanDistance(stars[j]) < 6 {
				ds.UnionWithAdd(stars[i], stars[j])
			}
		}
	}

	groups := make(map[Point2D]Set[Point2D])
	for _, star := range stars {
		s, _ := ds.Find(star)
		groups[s] = groups[s].UnionElems(star)
	}

	var costs []int
	for _, group := range groups {
		cost := MinimumSpanningTreeCost(group)
		costs = append(costs, cost+len(group))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(costs)))

	fmt.Println(Product(costs[0:3]...))
}

func MinimumSpanningTreeCost(vertices Set[Point2D]) int {
	children := func(p Point2D) []Point2D {
		var children []Point2D
		for q := range vertices {
			if p != q {
				children = append(children, q)
			}
		}
		return children
	}

	weight := func(p, q Point2D) int {
		return p.ManhattanDistance(q)
	}

	cost, _ := MinimumSpanningTree(vertices.Entries(), children, weight)
	return cost
}

func InputToStars() []Point2D {
	var stars []Point2D
	in.ToGrid2D(func(x, y int, s string) string {
		if s == "*" {
			stars = append(stars, Point2D{X: x, Y: y})
		}
		return s
	})

	return stars
}
