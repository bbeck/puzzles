package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	wall := in.Ints()

	find := func() (int, bool) {
		for i := range wall {
			if wall[i] > 0 {
				return i + 1, true
			}
		}
		return 0, false
	}

	var spell []int
	for {
		n, ok := find()
		if !ok {
			break
		}

		spell = append(spell, n)

		for i := n; i <= len(wall); i += n {
			wall[i-1]--
		}
	}
	fmt.Println(Product(spell...))
}
