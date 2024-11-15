package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	// L28: eqrr 3 0 1           B = (D == A) ? 1 : 0
	// L29: addr 1 4 4           IP = IP + B + 1
	// L30: seti 5 1 4           goto L06
	// -- END OF PROGRAM --
	//
	// By reading the program, it can be seen that the only time that register 0
	// is read is by the instruction eqrr 3 0 1.  Immediately after this read, we
	// add register 1 to the instruction pointer, exiting the program if the
	// comparison returned true.  So we just have to run the program until we
	// reach this instruction for the first time, see what value register 0 is
	// being checked against, and initialize register 0 to that value.
	//
	// For my particular input this value is: 16311888
	gt := func(a, b int) int {
		if a > b {
			return 1
		}
		return 0
	}

	eq := func(a, b int) int {
		if a == b {
			return 1
		}
		return 0
	}

	operations := map[string]Operation{
		"addr": func(a, b, c int, reg []int) { reg[c] = reg[a] + reg[b] },
		"addi": func(a, b, c int, reg []int) { reg[c] = reg[a] + b },
		"mulr": func(a, b, c int, reg []int) { reg[c] = reg[a] * reg[b] },
		"muli": func(a, b, c int, reg []int) { reg[c] = reg[a] * b },
		"banr": func(a, b, c int, reg []int) { reg[c] = reg[a] & reg[b] },
		"bani": func(a, b, c int, reg []int) { reg[c] = reg[a] & b },
		"borr": func(a, b, c int, reg []int) { reg[c] = reg[a] | reg[b] },
		"bori": func(a, b, c int, reg []int) { reg[c] = reg[a] | b },
		"setr": func(a, b, c int, reg []int) { reg[c] = reg[a] },
		"seti": func(a, b, c int, reg []int) { reg[c] = a },
		"gtir": func(a, b, c int, reg []int) { reg[c] = gt(a, reg[b]) },
		"gtri": func(a, b, c int, reg []int) { reg[c] = gt(reg[a], b) },
		"gtrr": func(a, b, c int, reg []int) { reg[c] = gt(reg[a], reg[b]) },
		"eqir": func(a, b, c int, reg []int) { reg[c] = eq(a, reg[b]) },
		"eqri": func(a, b, c int, reg []int) { reg[c] = eq(reg[a], b) },
		"eqrr": func(a, b, c int, reg []int) { reg[c] = eq(reg[a], reg[b]) },
	}

	ipr, program := InputToProgram()
	registers := make([]int, 6)

	for ip := 0; ip >= 0 && ip < len(program); ip++ {
		if ip == 28 {
			break
		}
		registers[ipr] = ip
		instruction := program[ip]
		operations[instruction.OpCode](instruction.A, instruction.B, instruction.C, registers)
		ip = registers[ipr]
	}

	fmt.Println(registers[3])
}

type Operation func(a, b, c int, reg []int)

type Instruction struct {
	OpCode  string
	A, B, C int
}

func InputToProgram() (int, []Instruction) {
	var ipr int
	var instructions []Instruction

	for _, line := range puz.InputToLines() {
		fields := strings.Fields(line)

		if fields[0] == "#ip" {
			ipr = puz.ParseInt(fields[1])
			continue
		}

		instructions = append(instructions, Instruction{
			OpCode: fields[0],
			A:      puz.ParseInt(fields[1]),
			B:      puz.ParseInt(fields[2]),
			C:      puz.ParseInt(fields[3]),
		})
	}

	return ipr, instructions
}
