package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var ps []Point2D
	for in.HasNext() {
		ps = append(ps, Point2D{X: in.Int(), Y: in.Int()})
	}

	var best int
	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			var dx, dy = Abs(ps[i].X-ps[j].X) + 1, Abs(ps[i].Y-ps[j].Y) + 1
			best = Max(best, dx*dy)
		}
	}
	fmt.Println(best)
}
