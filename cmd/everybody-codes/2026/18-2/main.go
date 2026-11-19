package main

import (
	"fmt"
	"math"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	segments := InputToSegments()

	var maxY int
	var index = make(map[int][]Segment)
	for _, s := range segments {
		index[s.Y] = append(index[s.Y], s)
		maxY = Max(maxY, s.Y)
	}

	var count int
	var left, right = math.MaxInt64, 0
	for y := maxY; y > 0; y-- {
		for _, s := range index[y] {
			left, right = Min(left, s.X0), Max(right, s.X1)
		}
		count += right - left
	}
	fmt.Println(count)
}

type Segment struct {
	X0, X1, Y int
}

func InputToSegments() []Segment {
	return in.LinesToS(func(in in.Scanner[Segment]) Segment {
		x, y, w := in.Int(), in.Int(), in.Int()
		return Segment{X0: x, X1: x + w, Y: y}
	})
}
