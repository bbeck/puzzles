package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	area := lib.Make2D[int](1001, 1001)
	for _, claim := range InputToClaims() {
		for dx := 0; dx < claim.Width; dx++ {
			for dy := 0; dy < claim.Height; dy++ {
				area[claim.TL.Y+dy][claim.TL.X+dx]++
			}
		}
	}

	var count int
	for y := 0; y <= 1000; y++ {
		for x := 0; x <= 1000; x++ {
			if area[y][x] > 1 {
				count++
			}
		}
	}
	fmt.Println(count)
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
