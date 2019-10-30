package main

import (
	"fmt"
)

func main() {
	N := 10551418

	var sum int
	for i := 1; i <= N; i++ {
		if N%i == 0 {
			sum += i
		}
	}

	fmt.Printf("r0: %d\n", sum)
}

func asm() {
	var r0, r1, r2, r4, r5 int
	r0 = 0
	goto INIT

	// This algorithm is putting into r0 the sum of the factors of r2.
	// For my particular input, when r0 starts as 0 the number it is interested in
	// is 1018.  When r0 starts as 1 (part 2) the number it is interested in is
	// 10551418.

L1:
	r1 = 1

L2:
	r4 = 1

L3:
	if r1*r4 == r2 {
		r0 += r1
	}

	r4 += 1
	if r4 <= r2 {
		goto L3
	}

	r1 += 1
	if r1 > r2 {
		goto END
	}

	goto L2

	// //////////////////////////////
	// Initialization code
	// //////////////////////////////

INIT:
	// 17: addi 2 2 2
	r2 += 2

	// 18: mulr 2 2 2
	r2 *= r2

	// 19: mulr 3 2 2
	r2 *= 19

	// 20: muli 2 11 2
	r2 *= 11

	// 21: addi 5 8 5
	r5 += 8

	// 22: mulr 5 3 5
	r5 *= 22

	// 23: addi 5 6 5
	r5 += 6

	// 24: addr 2 5 2
	r2 += r5

	// 25: addr 3 0 3
	// 26: seti 0 5 3
	if r0 == 0 {
		goto L1
	}

	// 27: setr 3 0 5
	r5 = 27

	// 28: mulr 5 3 5
	r5 *= 28

	// 29: addr 3 5 5
	r5 += 29

	// 30: mulr 3 5 5
	r5 *= 30

	// 31: muli 5 14 5
	r5 *= 14

	// 32: mulr 5 3 5
	r5 *= 32

	// 33: addr 2 5 2
	r2 += r5

	// 34: seti 0 8 0
	r0 = 0

	// 35: seti 0 9 3
	goto L1

END:
	fmt.Printf("r0: %d\n", r0)
	fmt.Printf("r1: %d\n", r1)
	fmt.Printf("r2: %d\n", r2)
	fmt.Printf("r4: %d\n", r4)
	fmt.Printf("r5: %d\n", r5)
}
