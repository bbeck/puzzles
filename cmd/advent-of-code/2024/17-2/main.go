package main

import (
	"fmt"
	"slices"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	registers, program := InputToRegistersAndProgram()
	B, C := registers["B"], registers["C"]

	// It's not completely clear to me what the program is doing, but one thing
	// that it is doing for sure is producing each output digit based on the
	// lower 3 bits of A.  Because of this we ought to be able to do perform a
	// DFS trying to brute force one digit of the output at a time.

	var solve func(int, int) (int, bool)
	solve = func(A int, n int) (int, bool) {
		if n < 0 {
			// This should be a solution, double check
			return A, slices.Equal(program, Run(A, B, C))
		}

		for da := 0; da < 8; da++ {
			out := Run(A<<3+da, B, C)

			if slices.Equal(program[n:], out) {
				if a, ok := solve(A<<3+da, n-1); ok {
					return a, true
				}
			}
		}

		return 0, false
	}

	A, _ := solve(0, len(program)-1)
	fmt.Println(A)
}

func Run(A, B, C int) []int {
	var output []int

	// This function implements my specific program as native code for better
	// performance.
	for {
		B = A % 8                    // bst 4
		B = B ^ 2                    // bxl 2
		C = A / (1 << B)             // cdv 5
		B = B ^ 3                    // bxl 3
		B = B ^ C                    // bxc 4
		output = append(output, B%8) // out 5
		A = A / (1 << 3)             // adv 3
		if A == 0 {                  // jnz 0
			break
		}
	}

	return output
}

func InputToRegistersAndProgram() (map[string]int, []int) {
	registers := make(map[string]int)
	var program []int

	for _, line := range InputToLines() {
		if line == "" {
			continue
		}

		line = strings.ReplaceAll(line, ":", "")
		line = strings.ReplaceAll(line, ",", " ")
		fields := strings.Fields(line)

		switch fields[0] {
		case "Register":
			registers[fields[1]] = ParseInt(fields[2])
		case "Program":
			for _, s := range fields[1:] {
				program = append(program, ParseInt(s))
			}
		}
	}

	return registers, program
}
