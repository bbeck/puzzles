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

	var ds DisjointSet[int]
	var largestId, largest int
	for p1 := range dna {
		for p2 := p1 + 1; p2 < len(dna); p2++ {
			for c := range dna {
				if c != p1 && c != p2 && isChild(c, p1, p2) {
					ds.UnionWithAdd(c, p1)
					ds.UnionWithAdd(c, p2)

					id, _ := ds.Find(c)
					if size := ds.Size(id); size > largest {
						largest = size
						largestId = id
					}
				}
			}
		}
	}

	var sum int
	for i := range dna {
		if id, ok := ds.Find(i); ok && id == largestId {
			sum += i + 1
		}
	}
	fmt.Println(sum)
}
