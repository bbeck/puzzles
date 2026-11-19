package main

import (
	"fmt"
	"math"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	ps := InputToPoints()

	var maxY int
	var index = make(map[int][]Point2D)
	for _, p := range ps {
		index[p.Y] = append(index[p.Y], p)
		maxY = Max(maxY, p.Y)
	}

	var count int
	var left, right = math.MaxInt64, 0
	for y := maxY; y > 0; y-- {
		for _, p := range index[y] {
			left, right = Min(left, p.X), Max(right, p.X)
		}
		count += right - left + 1
	}
	fmt.Println(count)
}

func InputToPoints() []Point2D {
	return in.LinesToS(func(in in.Scanner[Point2D]) Point2D {
		return Point2D{X: in.Int(), Y: in.Int()}
	})
}
