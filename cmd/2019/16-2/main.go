package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"log"
)

func main() {
	input := InputToDigits(2019, 16)

	var digits []int8
	for i := 0; i < 10000; i++ {
		digits = append(digits, input...)
	}
	if digits == nil {
		log.Fatal("digits is nil")
	}

	// Working through the example given in the problem statement we can see that
	// the last digit of each phase is the last digit from the previous phase.  We
	// can also see that the next to last digit becomes the sum of the previous
	// phases last two digits, and so on.
	//
	// This works because the coefficients of the pattern this deep in the matrix
	// are all 1's.
	//
	// So with this information and the assumption that the coefficients are
	// always 1's we can write a recurrence to compute any digit we care about.
	//
	// Let N = the length of the input
	// Let P = the number of phases
	//
	//   digit(P, N)     = digit(P-1, N)
	//   digit(P, N-1)   = digit(P-1, N) + digit(P-1, N-1)
	//   digit(P, N-2)   = digit(P-1, N) + digit(P-1, N-1) + digit(P-1, N-2)
	//   ...
	//   digit(P-1, N)   = digit(P-2, N)
	//   digit(P-1, N-1) = digit(P-2, N) + digit(P-2, N-1)
	//   digit(P-1, N-2) = digit(P-2, N) + digit(P-2, N-1) + digit(P-2, N-2)
	//   ...
	//   digit(1, N)     = digit(0, N)
	//   digit(1, N-1)   = digit(0, N) + digit(0, N-1)
	//   digit(1, N-2)   = digit(0, N) + digit(0, N-1) + digit(0, N-2)
	//   ...
	//
	// So armed with this information, we can start with the original digits and
	// then keep summing the tail of the array to compute the result from the
	// rightmost position towards the left.

	s := ""
	for i := 0; i < 7; i++ {
		s = s + fmt.Sprintf("%d", digits[i])
	}
	offset := aoc.ParseInt(s)

	var output []int8
	for i := offset; i < len(digits); i++ {
		output = append(output, digits[i])
	}
	if output == nil {
		log.Fatal("output is nil")
	}

	for phase := 0; phase < 100; phase++ {
		for i := len(output) - 2; i >= 0; i-- {
			output[i] = (output[i] + output[i+1]) % 10
		}
	}

	fmt.Print("output: ")
	for i := 0; i < 8; i++ {
		fmt.Print(output[i])
	}
	fmt.Println()
}

func InputToDigits(year, day int) []int8 {
	var digits []int8
	for _, b := range aoc.InputToString(year, day) {
		digits = append(digits, int8(aoc.ParseInt(string(b))))
	}

	return digits
}
