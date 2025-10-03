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

	for range 100 {
		DoRound(grid, rotations)
	}

	// Find the answer within the grid.
	var sb strings.Builder
	var inside bool
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			switch s := grid.Get(x, y); {
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

func DoRound(grid Grid2D[string], rotations string) {
	var n int
	for y := 1; y < grid.Height-1; y++ {
		for x := 1; x < grid.Width-1; x++ {
			Rotate(grid, x, y, rotations[n%len(rotations)])
			n++
		}
	}
}

var offsets = []Point2D{
	{X: -1, Y: -1}, {X: +0, Y: -1}, {X: +1, Y: -1}, {X: +1, Y: +0},
	{X: +1, Y: +1}, {X: +0, Y: +1}, {X: -1, Y: +1}, {X: -1, Y: +0},
}

func Rotate(grid Grid2D[string], x, y int, rotation byte) {
	var s []string
	for _, delta := range offsets {
		s = append(s, grid.Get(x+delta.X, y+delta.Y))
	}

	switch rotation {
	case 'R':
		s = append([]string{s[7]}, s[:8]...)
	case 'L':
		s = append(s[1:], s[0])
	}

	for i, delta := range offsets {
		grid.Set(x+delta.X, y+delta.Y, s[i])
	}
}
