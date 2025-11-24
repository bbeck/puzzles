package main

import (
	"fmt"
	"sort"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	spell := InputToSpell()

	count := func(length int) int {
		var count int
		for i := range spell {
			count += length / spell[i]
		}
		return count
	}

	length := sort.Search(1e15, func(i int) bool {
		return count(i) >= 202520252025000
	})
	fmt.Println(length - 1)
}

func InputToSpell() []int {
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

	return spell
}
