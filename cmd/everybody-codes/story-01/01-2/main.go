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
	// This time our exponents are quite large so we want to avoid doing all
	// iterations.  To do that we'll find the cycle that occurs and use that to
	// skip over them.
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

	// Determine how many elements we're going to take from the cycle.  We'd like
	// to take 5, but sometimes we haven't generated 5 elements.
	var take = Min(exp, 5)

	// Determine where in the cycle we should start.  If the cycle is short then
	// we have to start at the end, but if the cycle is long we could start
	// anywhere within it.
	var pos = (exp - 1) % len(cycle)
	if take < 5 {
		pos = take - 1
	}

	var digits []int
	for ; take > 0; take-- {
		digits = append(digits, Digits(cycle[pos])...)
		pos = (pos - 1 + len(cycle)) % len(cycle)
	}
	return JoinDigits(digits)
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
