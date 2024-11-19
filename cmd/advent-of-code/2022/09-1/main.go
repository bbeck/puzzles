package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

var Headings = map[byte]lib.Heading{
	'U': lib.Up,
	'D': lib.Down,
	'L': lib.Left,
	'R': lib.Right,
}

func main() {
	var head, tail lib.Point2D

	seen := lib.SetFrom(tail)
	for _, line := range lib.InputToLines() {
		dir := Headings[line[0]]
		n := lib.ParseInt(line[2:])
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

func MoveTowards(head, tail lib.Point2D) lib.Point2D {
	neighbors := lib.SetFrom(head.Neighbors()...)
	if neighbors.Contains(tail) {
		return tail
	}

	dx, dy := lib.Sign(head.X-tail.X), lib.Sign(head.Y-tail.Y)
	return lib.Point2D{X: tail.X + dx, Y: tail.Y + dy}
}
