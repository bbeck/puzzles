package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	claims := InputToClaims()

	var set DisjointSet[string]
	area := Make2D[string](1001, 1001)
	for _, claim := range claims {
		set.Add(claim.ID)

		for y := claim.TL.Y; y <= claim.TL.Y+claim.Height; y++ {
			for x := claim.TL.X; x <= claim.TL.X+claim.Width; x++ {
				if area[y][x] != "" {
					set.Union(area[y][x], claim.ID)
					continue
				}

				area[y][x] = claim.ID
			}
		}
	}

	for _, claim := range claims {
		if set.Size(claim.ID) == 1 {
			fmt.Println(claim.ID)
		}
	}
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
