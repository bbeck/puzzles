package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	banks := InputToBanks()
	N := len(banks)

	var seen lib.Set[string]
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
	}

	fmt.Println(cycle)
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

func InputToBanks() []int {
	var banks []int
	for _, field := range strings.Fields(lib.InputToString()) {
		banks = append(banks, lib.ParseInt(field))
	}
	return banks
}
