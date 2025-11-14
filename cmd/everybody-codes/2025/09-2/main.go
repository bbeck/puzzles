package main

import (
	"fmt"

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

	var parents = make(map[int][2]int)
	for p1 := range dna {
		for p2 := p1 + 1; p2 < len(dna); p2++ {
			for c := range dna {
				if c != p1 && c != p2 && isChild(c, p1, p2) {
					parents[c] = [2]int{p1, p2}
				}
			}
		}
	}

	var sum int
	for c, ps := range parents {
		var count1, count2 int
		for n := range N {
			if dna[c][n] == dna[ps[0]][n] {
				count1++
			}
			if dna[c][n] == dna[ps[1]][n] {
				count2++
			}
		}

		sum += count1 * count2
	}
	fmt.Println(sum)
}
