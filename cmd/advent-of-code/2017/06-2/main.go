package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	banks := in.Ints()
	N := len(banks)

	var seen Set[string]
	last := make(map[string]int)

	var cycle int
	for cycle = 1; ; cycle++ {
		remaining, index := Choose(banks)
		banks[index] = 0

		// Each bank will get div blocks added to it, and the first mod banks will
		// get an extra block.
		div, mod := remaining/N, remaining%N

		for i := 1; i <= N; i++ {
			var extra int
			if mod > 0 {
				extra = 1
				mod--
			}
			banks[(index+i+N)%N] += div + extra
		}

		if !seen.Add(ID(banks)) {
			break
		}

		last[ID(banks)] = cycle
	}

	fmt.Println(cycle - last[ID(banks)])
}

func ID(banks []int) string {
	return fmt.Sprintf("%v", banks)
}

func Choose(banks []int) (int, int) {
	var max, index int
	for i, bank := range banks {
		if bank > max {
			max = bank
			index = i
		}
	}

	return max, index
}
