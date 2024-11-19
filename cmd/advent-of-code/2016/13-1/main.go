package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"math/bits"
)

func main() {
	m := Maze(lib.InputToInt())
	start := lib.Point2D{X: 1, Y: 1}
	target := lib.Point2D{X: 31, Y: 39}

	children := func(p lib.Point2D) []lib.Point2D {
		var children []lib.Point2D
		for _, n := range p.OrthogonalNeighbors() {
			if n.X < 0 || n.Y < 0 || !m.IsOpen(n) {
				continue
			}
			children = append(children, n)
		}

		return children
	}

	goal := func(p lib.Point2D) bool {
		return p == target
	}

	cost := func(from, to lib.Point2D) int {
		return 1
	}

	_, length, found := lib.AStarSearch(start, children, goal, cost, target.ManhattanDistance)
	if !found {
		fmt.Println("no path found")
	}
	fmt.Println(length)
}

type Maze int

func (m Maze) IsOpen(p lib.Point2D) bool {
	n := uint(p.X*p.X + 3*p.X + 2*p.X*p.Y + p.Y + p.Y*p.Y + int(m))
	return bits.OnesCount(n)%2 == 0
}
