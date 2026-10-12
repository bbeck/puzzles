package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var best int
	for _, vars := range InputToVariables() {
		r1 := eni(vars["A"], vars["X"], vars["M"])
		r2 := eni(vars["B"], vars["Y"], vars["M"])
		r3 := eni(vars["C"], vars["Z"], vars["M"])
		best = Max(best, r1+r2+r3)
	}
	fmt.Println(best)
}

func eni(n, exp, mod int) int {
	pre, post := FindCycle(1, func(score int) int { return score * n % mod })

	// Build the cycle.  If there's no pre cycle elements that means our first
	// element (a 1 by construction) needs to be moved to the end of the cycle
	// because it participates.  If there's a pre cycle then no need to adjust
	// at all.
	var cycle []int
	for _, score := range post {
		cycle = append(cycle, score)
	}
	if len(pre) == 0 {
		cycle = append(cycle[1:], cycle[0])
	}

	// If our pre list starts with a 1 that's the synthetic 1 we started with, it
	// doesn't belong as it's not part of the sequence or the cycle.
	if len(pre) > 0 && pre[0] == 1 {
		pre = pre[1:]
	}

	// We now need to take exp remainders and sum them up.  Start with the tail
	// from before the cycle (skipping the 1).
	var take = exp

	var sum int
	for _, elem := range pre {
		sum += elem
		take--
	}

	// Next figure out how many times we're going to consume the full cycle.
	full, take := take/len(cycle), take%len(cycle)
	sum += full * Sum(cycle...)

	// Lastly, consume the last few elements.
	sum += Sum(cycle[:take]...)

	return sum
}

func InputToVariables() []map[string]int {
	return in.LinesToS(func(in in.Scanner[map[string]int]) map[string]int {
		vars := make(map[string]int)
		for _, field := range in.FieldsS() {
			lhs, rhs := field.CutS("=")
			vars[lhs.String()] = rhs.Int()
		}
		return vars
	})
}
