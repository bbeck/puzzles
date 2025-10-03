package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	rotations, grid := in.Line(), in.ToGrid2D(func(_, _ int, s string) string {
		return s
	})

	// Perform a single round on a grid of integers so that we can track where
	// each integer goes.
	tmp := NewGrid2D[int](grid.Width, grid.Height)
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			tmp.Set(x, y, y*grid.Width+x)
		}
	}
	DoRound(tmp, rotations)

	// Convert the grid into a permutation cycle represented as an array.
	var mapping = make([]int, grid.Width*grid.Height)
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			mapping[tmp.Get(x, y)] = y*grid.Width + x
		}
	}

	// Now exponentiate the permutation by the number of rounds we want to
	// perform.
	multiply := func(p1, p2 []int) []int {
		var result []int
		for i := range p1 {
			result = append(result, p1[p2[i]])
		}
		return result
	}

	var exponentiate func(p []int, n int) []int
	exponentiate = func(p []int, n int) []int {
		switch {
		case n == 0:
			var identity []int
			for i := range p {
				identity = append(identity, i)
			}
			return identity

		case n == 1:
			return p

		case n%2 == 0:
			half := exponentiate(p, n/2)
			return multiply(half, half)

		default: // n%2 == 1
			return multiply(p, exponentiate(p, n-1))
		}
	}

	mapping = exponentiate(mapping, 1048576000)

	// Apply the mapping to the original grid to get the final result.
	result := NewGrid2D[string](grid.Width, grid.Height)
	for p, q := range mapping {
		x1, y1 := p%grid.Width, p/grid.Width
		x2, y2 := q%grid.Width, q/grid.Width
		result.Set(x2, y2, grid.Get(x1, y1))
	}

	// Find the answer within the grid.
	var sb strings.Builder
	var inside bool
	for y := 0; y < result.Height; y++ {
		for x := 0; x < result.Width; x++ {
			switch s := result.Get(x, y); {
			case s == ">":
				inside = true
			case s == "<":
				inside = false
			case inside:
				sb.WriteString(s)
			}
		}
	}
	fmt.Println(sb.String())
}

func DoRound[T any](grid Grid2D[T], rotations string) {
	var n int
	for y := 1; y < grid.Height-1; y++ {
		for x := 1; x < grid.Width-1; x++ {
			Rotate(&grid, x, y, rotations[n%len(rotations)])
			n++
		}
	}
}

var offsets = []Point2D{
	{X: -1, Y: -1}, {X: +0, Y: -1}, {X: +1, Y: -1}, {X: +1, Y: +0},
	{X: +1, Y: +1}, {X: +0, Y: +1}, {X: -1, Y: +1}, {X: -1, Y: +0},
}

func Rotate[T any](grid *Grid2D[T], x, y int, rotation byte) {
	var s []T
	for _, delta := range offsets {
		s = append(s, grid.Get(x+delta.X, y+delta.Y))
	}

	switch rotation {
	case 'R':
		s = append([]T{s[7]}, s[:8]...)
	case 'L':
		s = append(s[1:], s[0])
	}

	for i, delta := range offsets {
		grid.Set(x+delta.X, y+delta.Y, s[i])
	}
}
