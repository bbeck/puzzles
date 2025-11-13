package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

const N = 256

func main() {
	nails := in.Ints()

	var starts, ends []int
	for a, b := range Zip(nails, nails[1:]) {
		starts = append(starts, Min(a, b))
		ends = append(ends, Max(a, b))
	}

	intersects := func(a, b, c, d int) bool {
		return (a < c && c < b && b < d) || (c < a && a < d && d < b)
	}

	var best int
	for s := 1; s <= N; s++ {
		for e := s + 1; e <= N; e++ {
			var count int
			for i := range starts {
				if intersects(s, e, starts[i], ends[i]) {
					count++
				}
			}

			best = Max(best, count)
		}
	}
	fmt.Println(best)
}
