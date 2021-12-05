package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	segments := InputToSegments()
	counts := make(map[aoc.Point2D]int)

	for _, s := range segments {
		if s.start.X == s.end.X {
			for y := aoc.MinInt(s.start.Y, s.end.Y); y <= aoc.MaxInt(s.start.Y, s.end.Y); y++ {
				counts[aoc.Point2D{X: s.start.X, Y: y}]++
			}
		}
		if s.start.Y == s.end.Y {
			for x := aoc.MinInt(s.start.X, s.end.X); x <= aoc.MaxInt(s.start.X, s.end.X); x++ {
				counts[aoc.Point2D{X: x, Y: s.start.Y}]++
			}
		}
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
	start, end aoc.Point2D
}

func InputToSegments() []Segment {
	var segments []Segment
	for _, line := range aoc.InputToLines(2021, 5) {
		var a, b, c, d int
		if _, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &a, &b, &c, &d); err != nil {
			log.Fatal(err)
		}

		segments = append(segments, Segment{
			start: aoc.Point2D{X: a, Y: b},
			end:   aoc.Point2D{X: c, Y: d},
		})
	}

	return segments
}
