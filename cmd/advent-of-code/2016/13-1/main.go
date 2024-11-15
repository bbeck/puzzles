package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"math/bits"
)

func main() {
	m := Maze(puz.InputToInt())
	start := puz.Point2D{X: 1, Y: 1}
	target := puz.Point2D{X: 31, Y: 39}

	children := func(p puz.Point2D) []puz.Point2D {
		var children []puz.Point2D
		for _, n := range p.OrthogonalNeighbors() {
			if n.X < 0 || n.Y < 0 || !m.IsOpen(n) {
				continue
			}
			children = append(children, n)
		}

		return children
	}

	goal := func(p puz.Point2D) bool {
		return p == target
	}

	cost := func(from, to puz.Point2D) int {
		return 1
	}

	_, length, found := puz.AStarSearch(start, children, goal, cost, target.ManhattanDistance)
	if !found {
		fmt.Println("no path found")
	}
	fmt.Println(length)
}

type Maze int

func (m Maze) IsOpen(p puz.Point2D) bool {
	n := uint(p.X*p.X + 3*p.X + 2*p.X*p.Y + p.Y + p.Y*p.Y + int(m))
	return bits.OnesCount(n)%2 == 0
}
