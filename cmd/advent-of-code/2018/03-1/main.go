package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	area := Make2D[int](1001, 1001) // TODO: Grid2D instead?
	for _, claim := range InputToClaims() {
		for dx := 0; dx < claim.Width; dx++ {
			for dy := 0; dy < claim.Height; dy++ {
				area[claim.TL.Y+dy][claim.TL.X+dx]++
			}
		}
	}

	var count int
	for y := range 1000 {
		for x := range 1000 {
			if area[y][x] > 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}

type Claim struct {
	ID            string
	TL            Point2D
	Width, Height int
}

func InputToClaims() []Claim {
	return in.LinesToS(func(in in.Scanner[Claim]) Claim {
		var c Claim
		in.Scanf("#%d @ %d,%d: %dx%d", &c.ID, &c.TL.X, &c.TL.Y, &c.Width, &c.Height)
		return c
	})
}
