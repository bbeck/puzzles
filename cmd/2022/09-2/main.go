package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

var Headings = map[byte]aoc.Heading{
	'U': aoc.Up,
	'D': aoc.Down,
	'L': aoc.Left,
	'R': aoc.Right,
}

func main() {
	var knots [10]aoc.Point2D

	seen := aoc.SetFrom(knots[9])
	for _, line := range aoc.InputToLines(2022, 9) {
		dir := Headings[line[0]]
		n := aoc.ParseInt(line[2:])
		knots[0] = knots[0].MoveN(dir, n)

		for {
			var changed bool
			for i := 1; i < len(knots); i++ {
				next := MoveTowards(knots[i-1], knots[i])
				changed = changed || knots[i] != next
				knots[i] = next
			}

			seen.Add(knots[9])

			if !changed {
				break
			}
		}
	}

	fmt.Println(len(seen))
}

func MoveTowards(head, tail aoc.Point2D) aoc.Point2D {
	neighbors := aoc.SetFrom(head.Neighbors()...)
	if neighbors.Contains(tail) {
		return tail
	}

	dx, dy := aoc.Sign(head.X-tail.X), aoc.Sign(head.Y-tail.Y)
	return aoc.Point2D{X: tail.X + dx, Y: tail.Y + dy}
}
