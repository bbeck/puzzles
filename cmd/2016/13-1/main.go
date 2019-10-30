package main

import (
	"fmt"
	"log"
	"math/bits"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	offset := uint(aoc.InputToInt(2016, 13))
	start := Location{Point2D: aoc.Point2D{1, 1}, offset: offset}
	goal := Location{Point2D: aoc.Point2D{31, 39}, offset: offset}

	cost := func(from, to aoc.Node) int {
		return 1
	}

	heuristic := func(node aoc.Node) int {
		l := node.(Location)
		return l.ManhattanDistance(goal.Point2D)
	}

	visit := func(node aoc.Node) bool {
		l := node.(Location)
		return l.Point2D == goal.Point2D
	}

	path, distance, found := aoc.AStarSearch(start, visit, cost, heuristic)
	if !found {
		log.Fatal("no path found")
	}

	fmt.Printf("path (distance: %d)\n", distance)
	for _, l := range path {
		fmt.Printf("  %s\n", l.ID())
	}
}

type Location struct {
	aoc.Point2D
	offset uint
}

func (l Location) ID() string {
	return fmt.Sprintf("(%d, %d)", l.X, l.Y)
}

func (l Location) Children() []aoc.Node {
	isOpen := func(p aoc.Point2D) bool {
		if p.X < 0 || p.Y < 0 {
			return false
		}

		x, y := uint(p.X), uint(p.Y)
		return bits.OnesCount(x*x+3*x+2*x*y+y+y*y+l.offset)%2 == 0
	}

	var neighbors []aoc.Node
	if isOpen(l.Up()) {
		neighbors = append(neighbors, Location{Point2D: l.Up(), offset: l.offset})
	}

	if isOpen(l.Down()) {
		neighbors = append(neighbors, Location{Point2D: l.Down(), offset: l.offset})
	}

	if isOpen(l.Left()) {
		neighbors = append(neighbors, Location{Point2D: l.Left(), offset: l.offset})
	}

	if isOpen(l.Right()) {
		neighbors = append(neighbors, Location{Point2D: l.Right(), offset: l.offset})
	}

	return neighbors
}
