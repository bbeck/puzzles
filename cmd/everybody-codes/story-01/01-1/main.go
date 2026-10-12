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
	var remainders []int

	var product = 1
	for range exp {
		product = product * n % mod
		remainders = append(Digits(product), remainders...)
	}

	return JoinDigits(remainders)
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
