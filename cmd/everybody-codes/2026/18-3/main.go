package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	segments := InputToSegments()

	var ys = []int{0}
	var index = make(map[int][]Segment)
	for _, s := range segments {
		ys = append(ys, s.Y)
		index[s.Y] = append(index[s.Y], s)
	}
	slices.SortFunc(ys, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	var count int
	var left, right, last_y = math.MaxInt64, 0, ys[0]
	for _, y := range ys {
		count += (right - left) * (last_y - y)
		last_y = y

		for _, s := range index[y] {
			left, right = Min(left, s.X0), Max(right, s.X1)
		}
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
