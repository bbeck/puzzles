package main

import (
	"fmt"
	"math"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid, instructions := InputToGridAndInstructions()

	var scores [][]int
	for i := range instructions {
		scores = append(scores, nil)
		for toss := 1; toss <= grid.Width/2+1; toss++ {
			final := Drop(grid, instructions[i], toss)
			scores[i] = append(scores[i], Max(0, 2*final-toss))
		}
	}

	// Search for the optimal scores to use.  This implements a small DP algorithm
	// where a set of span of the best known answers are built up one row of
	// scores at a time.  Each span is indexed by the bitset of columns that have
	// been used by prior rows in order to ensure we don't repeat.
	type span struct{ lo, hi int }
	states := map[BitSet]span{0: {0, 0}}

	for _, row := range scores {
		next := make(map[BitSet]span, len(states)*len(row))
		for used, s := range states {
			for col, score := range row {
				if used.Contains(col) {
					continue
				}

				key := used.Add(col)
				cur, ok := next[key]
				if !ok {
					cur = span{math.MaxInt, math.MinInt}
				}
				cur.lo = Min(cur.lo, s.lo+score)
				cur.hi = Max(cur.hi, s.hi+score)
				next[key] = cur
			}
		}
		states = next
	}

	// Find the best value from each state of the last row
	var lo, hi = math.MaxInt, math.MinInt
	for _, state := range states {
		lo, hi = Min(lo, state.lo), Max(hi, state.hi)
	}
	fmt.Println(lo, hi)
}

func Drop(grid Grid2D[string], instructions string, toss int) int {
	s2p := func(s int) Point2D { return Point2D{X: 2 * (s - 1), Y: -1} }
	p2s := func(p Point2D) int { return p.X/2 + 1 }

	var p = s2p(toss)
	for p.Y < grid.Height-1 {
		switch grid.GetPoint(p.Down()) {
		case ".":
			p = p.Down()

		case "*":
			var op byte
			op, instructions = instructions[0], instructions[1:]

			if op == 'L' {
				if p.X > 0 {
					p = p.Left()
				} else {
					p = p.Right() // Bounce off the left wall
				}
			} else if op == 'R' {
				if p.X < grid.Width-1 {
					p = p.Right()
				} else {
					p = p.Left() // Bounce off the right wall
				}
			}
		}
	}

	return p2s(p)
}

func InputToGridAndInstructions() (Grid2D[string], []string) {
	chunk := in.ChunkS[string]()
	return chunk.StringGrid2D(), in.Lines()
}
