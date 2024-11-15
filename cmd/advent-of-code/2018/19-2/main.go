package main

import (
	"fmt"
)

func main() {
	// The input is an assembly language program of a computationally inefficient
	// algorithm.  In pseudocode the core of the algorithm looks like this.
	//
	// for r1 := 1; r1 <= r2; r1++ {
	//   for r4 := 1; r4 <= r2; r4++ {
	//     if r1*r4 == r2 {
	//       r0 += r1
	//     }
	//   }
	// }
	//
	// This treats r2 as an input, and r0 as an accumulator.  The algorithm
	// attempts to find all factors of r2 and sums them in r0.
	//
	// The input, r2, is computed at the beginning of the program by the
	// following pseudocode.
	//
	// r2 = 2*2*19*11 + 8*22+6
	// if r0 == 1 {
	//   r2 += (27*28 + 29)*30*14*32
	// }
	//
	// So when r0 == 0, the input is 1018, and when r0 == 1 the input is
	// 10551418.
	//
	// Given ths we can implement the equivalent algorithm a bit more
	// efficiently.
	N := 10551418

	var sum int
	for i := 1; i <= N; i++ {
		if N%i == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
