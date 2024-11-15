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
	var head, tail puz.Point2D

	seen := puz.SetFrom(tail)
	for _, line := range puz.InputToLines() {
		dir := Headings[line[0]]
		n := puz.ParseInt(line[2:])
		head = head.MoveN(dir, n)

		for {
			next := MoveTowards(head, tail)
			if next == tail {
				break
			}

			tail = next
			seen.Add(tail)
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
