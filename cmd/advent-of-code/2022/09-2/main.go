package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

var Headings = map[byte]lib.Heading{
	'U': lib.Up,
	'D': lib.Down,
	'L': lib.Left,
	'R': lib.Right,
}

func main() {
	var knots [10]lib.Point2D

	seen := lib.SetFrom(knots[9])
	for _, line := range lib.InputToLines() {
		dir := Headings[line[0]]
		n := lib.ParseInt(line[2:])
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

func MoveTowards(head, tail lib.Point2D) lib.Point2D {
	neighbors := lib.SetFrom(head.Neighbors()...)
	if neighbors.Contains(tail) {
		return tail
	}

	dx, dy := lib.Sign(head.X-tail.X), lib.Sign(head.Y-tail.Y)
	return lib.Point2D{X: tail.X + dx, Y: tail.Y + dy}
}
