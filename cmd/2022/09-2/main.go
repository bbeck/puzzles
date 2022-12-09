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
	var head aoc.Point2D
	var tails [9]aoc.Point2D

	seen := aoc.SetFrom(tails[8])
	for _, line := range aoc.InputToLines(2022, 9) {
		dir := Headings[line[0]]
		n := aoc.ParseInt(line[2:])
		head = head.MoveN(dir, n)

		for {
			var changed bool
			for i, tail := range tails {
				target := head
				if i != 0 {
					target = tails[i-1]
				}

				next := MoveTowards(target, tail)
				changed = changed || tail != next
				tails[i] = next
			}

			seen.Add(tails[8])

			if !changed {
				break
			}
		}
	}

	fmt.Println(len(seen))
}

func MoveTowards(head, tail aoc.Point2D) aoc.Point2D {
	neighbors := aoc.SetFrom(head.Neighbors()...)
	if head == tail || neighbors.Contains(tail) {
		return tail
	}

	dx, dy := head.X-tail.X, head.Y-tail.Y
	if dx != 0 {
		dx = aoc.Abs(dx) / dx
	}
	if dy != 0 {
		dy = aoc.Abs(dy) / dy
	}

	return aoc.Point2D{X: tail.X + dx, Y: tail.Y + dy}
}
