package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	dna := in.LinesToS(func(in in.Scanner[string]) string {
		_, rhs := in.Cut(":")
		return rhs
	})
	N := len(dna[0])

	isChild := func(c, p1, p2 int) bool {
		for n := range N {
			if dna[c][n] != dna[p1][n] && dna[c][n] != dna[p2][n] {
				return false
			}
		}
		return true
	}

	var c, p1, p2 int
	EnumeratePermutations(3, func(indices []int) bool {
		c, p1, p2 = indices[0], indices[1], indices[2]
		return isChild(c, p1, p2)
	})

	var count1, count2 int
	for n := range N {
		if dna[c][n] == dna[p1][n] {
			count1++
		}
		if dna[c][n] == dna[p2][n] {
			count2++
		}
	}
	fmt.Println(count1 * count2)
}
