package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	N := 40

	var count int
	for r, i := InputToRow(2016, 18), 0; i < N; r, i = r.Next(), i+1 {
		for _, tile := range r {
			if tile == SAFE {
				count++
			}
		}
	}

	fmt.Printf("number of safe tiles: %d\n", count)
}

const (
	SAFE = false
	TRAP = true
)

type Row []bool

func (r Row) Next() Row {
	var next Row
	for x := 0; x < len(r); x++ {
		var left, center, right bool
		if x > 0 {
			left = r[x-1]
		}
		center = r[x]
		if x < len(r)-1 {
			right = r[x+1]
		}

		isTrap := (left == TRAP && center == TRAP && right == SAFE) ||
			(left == SAFE && center == TRAP && right == TRAP) ||
			(left == TRAP && center == SAFE && right == SAFE) ||
			(left == SAFE && center == SAFE && right == TRAP)
		next = append(next, isTrap)
	}

	return next
}

func (r Row) String() string {
	var builder strings.Builder
	for _, b := range r {
		if b == SAFE {
			builder.WriteByte('.')
		} else {
			builder.WriteByte('^')
		}
	}

	return builder.String()
}

func InputToRow(year, day int) Row {
	var row []bool
	for _, c := range aoc.InputToString(year, day) {
		row = append(row, c == '^')
	}

	return row
}
