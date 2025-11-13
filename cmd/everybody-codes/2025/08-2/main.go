package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

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

	var count int
	for i := 0; i < len(starts); i++ {
		for j := i + 1; j < len(starts); j++ {
			if intersects(starts[i], ends[i], starts[j], ends[j]) {
				count++
			}
		}
	}
	fmt.Println(count)
}
