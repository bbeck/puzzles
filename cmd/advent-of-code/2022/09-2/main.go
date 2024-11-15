package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

var Headings = map[byte]puz.Heading{
	'U': puz.Up,
	'D': puz.Down,
	'L': puz.Left,
	'R': puz.Right,
}

func main() {
	var knots [10]puz.Point2D

	seen := puz.SetFrom(knots[9])
	for _, line := range puz.InputToLines() {
		dir := Headings[line[0]]
		n := puz.ParseInt(line[2:])
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

func MoveTowards(head, tail puz.Point2D) puz.Point2D {
	neighbors := puz.SetFrom(head.Neighbors()...)
	if neighbors.Contains(tail) {
		return tail
	}

	dx, dy := puz.Sign(head.X-tail.X), puz.Sign(head.Y-tail.Y)
	return puz.Point2D{X: tail.X + dx, Y: tail.Y + dy}
}
