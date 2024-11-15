package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var sum int
	for _, pattern := range InputToPatterns() {
		if col, found := FindVerticalReflection(pattern); found {
			sum += col
		}
		if row, found := FindHorizontalReflection(pattern); found {
			sum += 100 * row
		}
	}

	fmt.Println(sum)
}

func FindVerticalReflection(p puz.Grid2D[string]) (int, bool) {
outer:
	for rhs := 1; rhs < p.Width; rhs++ {
		N := puz.Min(rhs, p.Width-rhs)

		for y := 0; y < p.Height; y++ {
			for n := 1; n <= N; n++ {
				x1, x2 := rhs-n, rhs+n-1

				if p.Get(x1, y) != p.Get(x2, y) {
					continue outer
				}
			}
		}

		return rhs, true
	}

	return 0, false
}

func FindHorizontalReflection(p puz.Grid2D[string]) (int, bool) {
	return FindVerticalReflection(p.RotateLeft())
}

func InputToPatterns() []puz.Grid2D[string] {
	chunks := puz.Split(puz.InputToLines(), func(s string) bool {
		return s != ""
	})

	var patterns []puz.Grid2D[string]
	for _, chunk := range chunks {
		pattern := puz.NewGrid2D[string](len(chunk[0]), len(chunk))
		for y := 0; y < len(chunk); y++ {
			for x := 0; x < len(chunk[y]); x++ {
				pattern.Set(x, y, string(chunk[y][x]))
			}
		}

		patterns = append(patterns, pattern)
	}

	return patterns
}
