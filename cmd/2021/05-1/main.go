package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ps := make(map[aoc.Point2D]int)
	for _, segment := range InputToSegments() {
		dx, dy := Interpolate(segment.Start, segment.End)
		if dy != 0 && dx != 0 {
			continue
		}

		p := segment.Start
		for {
			ps[p]++
			if p == segment.End {
				break
			}
			p = aoc.Point2D{X: p.X + dx, Y: p.Y + dy}
		}
	}

	var count int
	for _, v := range aoc.GetMapValues(ps) {
		if v > 1 {
			count++
		}
	}
	fmt.Println(count)
}

func Interpolate(p, q aoc.Point2D) (int, int) {
	var dx, dy int
	if p.X < q.X {
		dx = 1
	} else if p.X > q.X {
		dx = -1
	}

	if p.Y < q.Y {
		dy = 1
	} else if p.Y > q.Y {
		dy = -1
	}

	return dx, dy
}

type Segment struct {
	Start, End aoc.Point2D
}

func InputToSegments() []Segment {
	return aoc.InputLinesTo(2021, 5, func(line string) Segment {
		var s Segment
		fmt.Sscanf(line, "%d,%d -> %d,%d", &s.Start.X, &s.Start.Y, &s.End.X, &s.End.Y)
		return s
	})
}
