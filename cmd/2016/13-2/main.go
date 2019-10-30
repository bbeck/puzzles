package main

import (
	"fmt"
	"math/bits"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	offset := uint(aoc.InputToInt(2016, 13))
	start := Location{Point2D: aoc.Point2D{1, 1}, offset: offset}

	seen := make(map[Location]bool)
	aoc.BreadthFirstSearch(start, func(node aoc.Node) bool {
		seen[node.(Location)] = true
		return false
	})

	fmt.Printf("number of reachable locations: %d\n", len(seen))
}

type Location struct {
	aoc.Point2D
	offset   uint
	distance int
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

	var children []aoc.Node
	if l.distance < 50 && isOpen(l.Up()) {
		child := Location{Point2D: l.Up(), offset: l.offset, distance: l.distance + 1}
		children = append(children, child)
	}

	if l.distance < 50 && isOpen(l.Down()) {
		child := Location{Point2D: l.Down(), offset: l.offset, distance: l.distance + 1}
		children = append(children, child)
	}

	if l.distance < 50 && isOpen(l.Left()) {
		child := Location{Point2D: l.Left(), offset: l.offset, distance: l.distance + 1}
		children = append(children, child)
	}

	if l.distance < 50 && isOpen(l.Right()) {
		child := Location{Point2D: l.Right(), offset: l.offset, distance: l.distance + 1}
		children = append(children, child)
	}

	return children
}
