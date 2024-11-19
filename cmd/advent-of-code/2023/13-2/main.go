package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

var Smudge = map[string]string{
	"#": ".",
	".": "#",
}

func main() {
	var sum int
	for _, pattern := range InputToPatterns() {
		oldCol, oldRow, _ := FindReflection(pattern, 0, 0)

		// Try smudging each point until we find a new reflection
		var col, row int
		var found bool
		for y := 0; y < pattern.Height && !found; y++ {
			for x := 0; x < pattern.Width && !found; x++ {
				pattern.Set(x, y, Smudge[pattern.Get(x, y)])
				col, row, found = FindReflection(pattern, oldCol, oldRow)
				pattern.Set(x, y, Smudge[pattern.Get(x, y)])
			}
		}

		sum += 100*row + col
	}

	fmt.Println(sum)
}

func FindReflection(p lib.Grid2D[string], skipCol, skipRow int) (int, int, bool) {
	if col, found := FindVerticalReflection(p, skipCol); found {
		return col, 0, true
	}
	if row, found := FindHorizontalReflection(p, skipRow); found {
		return 0, row, true
	}
	return 0, 0, false
}

func FindVerticalReflection(p lib.Grid2D[string], skip int) (int, bool) {
outer:
	for rhs := 1; rhs < p.Width; rhs++ {
		if rhs == skip {
			continue
		}

		N := lib.Min(rhs, p.Width-rhs)

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

func FindHorizontalReflection(p lib.Grid2D[string], skip int) (int, bool) {
	return FindVerticalReflection(p.RotateLeft(), skip)
}

func InputToPatterns() []lib.Grid2D[string] {
	chunks := lib.Split(lib.InputToLines(), func(s string) bool {
		return s != ""
	})

	var patterns []lib.Grid2D[string]
	for _, chunk := range chunks {
		pattern := lib.NewGrid2D[string](len(chunk[0]), len(chunk))
		for y := 0; y < len(chunk); y++ {
			for x := 0; x < len(chunk[y]); x++ {
				pattern.Set(x, y, string(chunk[y][x]))
			}
		}

		patterns = append(patterns, pattern)
	}

	return patterns
}
