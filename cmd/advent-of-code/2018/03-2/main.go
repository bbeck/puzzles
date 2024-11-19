package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	claims := InputToClaims()

	var set lib.DisjointSet[string]
	area := lib.Make2D[string](1001, 1001)
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
	TL            lib.Point2D
	Width, Height int
}

func InputToClaims() []Claim {
	return lib.InputLinesTo(func(line string) Claim {
		var claim Claim
		fmt.Sscanf(line, "#%s @ %d,%d: %dx%d", &claim.ID, &claim.TL.X, &claim.TL.Y, &claim.Width, &claim.Height)
		return claim
	})
}
