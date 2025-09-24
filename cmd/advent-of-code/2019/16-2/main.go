package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	ds := InputToDigits()
	offset := JoinDigits(ds[:7])

	// The specified offset is >91% of the way through the list of repeated
	// digits, implying that we're focusing on the tail of the transform.
	//
	// By observing the examples in the problem statement we can see that
	// within the tail the n-th digit of a phase is just the sum of the
	// n-1 digits (mod 10) from the previous phase.  This ends up being true
	// because the coefficients of the tail end up all being 1.
	//
	// This observation leads to a dynamic programming based solution
	// for directly computing the value of a specific tail digit.

	// Build the tail of the digits starting at our offset.
	digits := make([]int, 0, len(ds)*10000)
	for range 10000 {
		digits = append(digits, ds...)
	}
	digits = digits[offset:]

	for range 100 {
		for i := len(digits) - 2; i >= 0; i-- {
			digits[i] = (digits[i] + digits[i+1]) % 10
		}
	}

	fmt.Println(JoinDigits(digits[:8]))
}

func InputToDigits() []int {
	var digits []int
	for in.HasNext() {
		digits = append(digits, ParseInt(string(in.Byte())))
	}
	return digits
}
