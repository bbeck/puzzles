package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	bases, coeffs := InputToBasesAndCoeffs()

	var sum int
	for i := range bases {
		base, coeff := Reversed(Digits(bases[i])), Reversed(Digits(coeffs[i]))
		sum += ISO(base, coeff)
	}
	fmt.Println(sum)
}

func ISO(base, coeff []int) int {
	var value int
	for i := range coeff {
		value += coeff[i] * Pow(base[i%len(base)], uint(i))
	}
	return value
}

func InputToBasesAndCoeffs() ([]int, []int) {
	in.Remove("ISO-")

	var bases, coeffs []int
	for in.HasNextLine() {
		bases = append(bases, in.Int())
		coeffs = append(coeffs, in.Int())
	}
	return bases, coeffs
}
