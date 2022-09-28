package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	claims := InputToClaims()

	var set aoc.DisjointSet[string]
	area := aoc.Make2D[string](1001, 1001)
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
	TL            aoc.Point2D
	Width, Height int
}

func InputToClaims() []Claim {
	return aoc.InputLinesTo(2018, 3, func(line string) (Claim, error) {
		var claim Claim
		fmt.Sscanf(line, "#%s @ %d,%d: %dx%d", &claim.ID, &claim.TL.X, &claim.TL.Y, &claim.Width, &claim.Height)

		return claim, nil
	})
}
