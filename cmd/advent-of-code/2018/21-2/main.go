package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	// #ip 4                     registers: A B C D ip E
	// L00: seti 123 0 3         D = 123
	// L01: bani 3 456 3         D = D & 456
	// L02: eqri 3 72 3          D = (D == 72) ? 1 : 0
	// L03: addr 3 4 4           jump D+1
	// L04: seti 0 0 4           goto L01
	// L05: seti 0 5 3           D = 0
	// L06: bori 3 65536 2       C = D | 65536
	// L07: seti 10736359 9 3    D = 10736359
	// L08:   bani 2 255 1         B = C & 255
	// L09:   addr 3 1 3           D = D + B
	// L10:   bani 3 16777215 3    D = D & 16777215
	// L11:   muli 3 65899 3       D = D * 65899
	// L12:   bani 3 16777215 3    D = D & 16777215
	// L13:   gtir 256 2 1         B = (256 > C) ? 1 : 0
	// L14:   addr 1 4 4           jump B+1
	// L15:   addi 4 1 4           goto L17
	// L16:   seti 27 2 4          goto L28
	// L17:     seti 0 3 1           B = 0
	// L18:     addi 1 1 5           E = B + 1
	// L19:     muli 5 256 5         E = E * 256
	// L20:     gtrr 5 2 5           E = (E > C) ? 1 : 0
	// L21:     addr 5 4 4           jump E+1
	// L22:     addi 4 1 4           goto L24
	// L23:     seti 25 8 4          goto L26
	// L24:     addi 1 1 1           B = B + 1
	// L25:     seti 17 6 4          goto L18
	// L26:   setr 1 5 2           C = B
	// L27:   seti 7 7 4           goto L08
	// L28: eqrr 3 0 1           B = (D == A) ? 1 : 0
	// L29: addr 1 4 4           jump B+1
	// L30: seti 5 1 4           goto L06

	// The above code for my problem has been translated into go for performance.
	// The problem asks us to find the value in register A that will cause the
	// program to run the longest, but still terminate.  Observing the program
	// shows that the only time that register A is read in the program is on line
	// 28 where it is compared to D and a jump is made based on the comparison.
	//
	// Given this, we can instrument the program to observe what the value of D
	// is right before the program goes into an infinite loop.  That will cause
	// the maximum number of iterations of the program before terminating.

	// Visit is called at the end of each of the outermost iterations of the
	// program.  Its return value determines if the program should terminate.
	var seen lib.Set[uint]
	var last uint
	visit := func(d uint) bool {
		if !seen.Add(d) {
			return true
		}

		last = d
		return false
	}

	var b, c, d uint
	for {
		c = d | 65536
		d = 10736359

		for {
			d = d + (c & 255)
			d = d & 16777215
			d = d * 65899
			d = d & 16777215

			if c < 256 {
				break
			}

			b = 0
			for {
				if (b+1)*256 > c {
					break
				}
				b++
			}

			c = b
		}

		// This is line 28 of the program.  Let our visitor decide what happens.
		if visit(d) {
			break
		}
	}

	fmt.Println(last)
}
