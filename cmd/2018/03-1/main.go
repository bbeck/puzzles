package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	counts := make(map[aoc.Point2D]int)
	for _, claim := range InputToClaims(2018, 3) {
		for x := claim.left; x < claim.left+claim.width; x++ {
			for y := claim.top; y < claim.top+claim.height; y++ {
				counts[aoc.Point2D{X: x, Y: y}]++
			}
		}
	}

	var overlap int
	for _, count := range counts {
		if count > 1 {
			overlap++
		}
	}

	fmt.Printf("overlap: %d\n", overlap)
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
