package main

import (
	"fmt"
	"math/bits"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	m := Maze(in.Int())
	start := Point2D{X: 1, Y: 1}
	target := Point2D{X: 31, Y: 39}

	children := func(p Point2D) []Point2D {
		var children []Point2D
		for _, n := range p.OrthogonalNeighbors() {
			if n.X < 0 || n.Y < 0 || !m.IsOpen(n) {
				continue
			}
			children = append(children, n)
		}

		return children
	}

	goal := func(p Point2D) bool {
		return p == target
	}

	cost := func(from, to Point2D) int {
		return 1
	}

	_, length, found := AStarSearch(start, children, goal, cost, target.ManhattanDistance)
	if !found {
		fmt.Println("no path found")
	}
	fmt.Println(length)
}

type Maze int

func (m Maze) IsOpen(p Point2D) bool {
	n := uint(p.X*p.X + 3*p.X + 2*p.X*p.Y + p.Y + p.Y*p.Y + int(m))
	return bits.OnesCount(n)%2 == 0
}
