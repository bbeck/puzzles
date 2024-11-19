package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	ps := make(map[lib.Point2D]int)
	for _, segment := range InputToSegments() {
		dx, dy := Interpolate(segment.Start, segment.End)

		p := segment.Start
		for {
			ps[p]++
			if p == segment.End {
				break
			}
			p = lib.Point2D{X: p.X + dx, Y: p.Y + dy}
		}
	}

	var count int
	for _, v := range lib.GetMapValues(ps) {
		if v > 1 {
			count++
		}
	}
	fmt.Println(count)
}

func Interpolate(p, q lib.Point2D) (int, int) {
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
	Start, End lib.Point2D
}

func InputToSegments() []Segment {
	return lib.InputLinesTo(func(line string) Segment {
		var s Segment
		fmt.Sscanf(line, "%d,%d -> %d,%d", &s.Start.X, &s.Start.Y, &s.End.X, &s.End.Y)
		return s
	})
}
