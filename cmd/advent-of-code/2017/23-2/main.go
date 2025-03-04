package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib/in"
	"strconv"
)

func main() {
	// The input is an assembly language program that effectively counts how many
	// composite integers there are between two numbers (b and c with a step size
	// of 17).  The algorithm looks something like this:
	//
	// for b, c := 106700, 123700; b != c; b += 17 {
	//   isComposite = false  // f register
	//
	//   for d := 2; d != b; c++ {
	//     for e := 2; e != b; e++ {
	//       if d*e == b {
	//         isComposite = true
	//       }
	//     }
	//   }
	//
	//   if isComposite {
	//     h++
	//   }
	// }
	//
	// This implementation is quite inefficient because the nested loops on d and
	// e.  One way to drastically optimize this would be to provide a mod
	// instruction.  That would enable removing the innermost loop on e entirely.
	// Another optimization is to provide a div instruction that divides a
	// register by a value.  This would allow the remaining inner loop to only
	// check for divisors up to b/2.
	//
	// Unfortunately even with these optimizations the interpreter is still fairly
	// slow, so we'll just implement the algorithm directly and not use the
	// interpreter.

	// Read the b value from the program.
	b := InputToProgram()[0].Parsed[1]
	b = b*100 + 100000
	c := b + 17000

	h := 0
	for ; b <= c; b += 17 {
		isComposite := false

		for d := 2; d < b; d++ {
			if b%d == 0 {
				isComposite = true
				break
			}
		}

		if isComposite {
			h++
		}
	}

	fmt.Println(h)
}

type Instruction struct {
	OpCode string
	Args   []string
	Parsed []int
}

func InputToProgram() []Instruction {
	return in.LinesToS(func(in in.Scanner[Instruction]) Instruction {
		var opcode = in.String()
		var args = in.Fields()

		var parsed = make([]int, len(args))
		for i, arg := range args {
			if n, err := strconv.Atoi(arg); err == nil {
				parsed[i] = n
			}
		}

		return Instruction{OpCode: opcode, Args: args, Parsed: parsed}
	})
}
