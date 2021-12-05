package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	counts := make(map[aoc.Point2D]int)
	for _, s := range InputToSegments() {
		dx, dy := Slope(s.Start, s.End)

		for p := s.Start; p != s.End; {
			counts[p]++
			p = aoc.Point2D{X: p.X + dx, Y: p.Y + dy}
		}
		counts[s.End]++
	}

	var count int
	for _, c := range counts {
		if c > 1 {
			count++
		}
	}
	fmt.Println(count)
}

type Segment struct {
	Start aoc.Point2D
	End   aoc.Point2D
}

func InputToSegments() []Segment {
	var segments []Segment
	for _, line := range aoc.InputToLines(2021, 5) {
		var a, b, c, d int
		if _, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &a, &b, &c, &d); err != nil {
			log.Fatal(err)
		}

		segments = append(segments, Segment{
			Start: aoc.Point2D{X: a, Y: b},
			End:   aoc.Point2D{X: c, Y: d},
		})
	}

	return segments
}

func Slope(p, q aoc.Point2D) (int, int) {
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
