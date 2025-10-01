package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	stars := InputToStars()

	children := func(p Point2D) []Point2D {
		var children []Point2D
		for _, q := range stars {
			if p != q {
				children = append(children, q)
			}
		}
		return children
	}

	weight := func(p, q Point2D) int {
		return p.ManhattanDistance(q)
	}

	cost, _ := MinimumSpanningTree(stars, children, weight)
	fmt.Println(cost + len(stars))
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
