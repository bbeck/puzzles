package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	claims := InputToClaims(2018, 3)

	// index the claims into the grid and keep track of the ids of the overlapping
	// claims.
	grid := make(map[aoc.Point2D][]int)
	overlaps := make(map[int]bool) // claim id -> true if it has overlapped

	for _, claim := range claims {
		for x := claim.left; x < claim.left+claim.width; x++ {
			for y := claim.top; y < claim.top+claim.height; y++ {
				p := aoc.Point2D{X: x, Y: y}
				grid[p] = append(grid[p], claim.id)

				if len(grid[p]) > 1 {
					for _, id := range grid[p] {
						overlaps[id] = true
					}
				}
			}
		}
	}

	// Find the id that doesn't ever overlap with another id.
	for _, claim := range claims {
		if !overlaps[claim.id] {
			fmt.Printf("claim %d does not overlap\n", claim.id)
		}
	}
}

type Claim struct {
	id            int
	left, top     int
	width, height int
}

func InputToClaims(year, day int) []Claim {
	var claims []Claim
	for _, line := range aoc.InputToLines(year, day) {
		var id, left, top, width, height int
		if _, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &left, &top, &width, &height); err != nil {
			log.Fatalf("unable to parse claim: %s", line)
		}

		claims = append(claims, Claim{id, left, top, width, height})
	}

	return claims
}
