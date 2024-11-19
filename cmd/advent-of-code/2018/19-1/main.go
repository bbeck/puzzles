package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
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
		registers[ipr] = ip
		instruction := program[ip]
		operations[instruction.OpCode](instruction.A, instruction.B, instruction.C, registers)
		ip = registers[ipr]
	}

	fmt.Println(registers[0])
}

type Operation func(a, b, c int, reg []int)

type Instruction struct {
	OpCode  string
	A, B, C int
}

func InputToProgram() (int, []Instruction) {
	var ipr int
	var instructions []Instruction

	for _, line := range lib.InputToLines() {
		fields := strings.Fields(line)

		if fields[0] == "#ip" {
			ipr = lib.ParseInt(fields[1])
			continue
		}

		instructions = append(instructions, Instruction{
			OpCode: fields[0],
			A:      lib.ParseInt(fields[1]),
			B:      lib.ParseInt(fields[2]),
			C:      lib.ParseInt(fields[3]),
		})
	}

	return ipr, instructions
}
