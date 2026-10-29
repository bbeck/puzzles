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
		base, coeff := RDigits(bases[i]), RDigits(coeffs[i])
		target := ISO(base, coeff)

		// Now that we know the target, see if we can find a smaller coefficient
		// that yields the same value.
		for j := range coeffs[i] {
			if ISO(base, RDigits(j)) == target {
				target = 0
				break
			}
		}
		sum += target
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

var seen = make(map[int][]int)

func RDigits(n int) []int {
	if value, ok := seen[n]; ok {
		return value
	}
	value := Reversed(Digits(n))
	seen[n] = value
	return value
}

func InputToBasesAndCoeffs() ([]int, []int) {
	in.Remove("ISO2-")

	var bases, coeffs []int
	for in.HasNextLine() {
		bases = append(bases, in.Int())
		coeffs = append(coeffs, in.Int())
	}
	return bases, coeffs
}
