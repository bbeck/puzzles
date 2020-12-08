package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	program := cpus.InputToHandheldProgram(2020, 8)

	// Loop through each instruction of the program considering what happens if it
	// were changed from a jmp -> nop or nop -> jmp.
	var acc int
	for i := 0; i < len(program); i++ {
		var infinite bool

		switch program[i].Op {
		case "jmp":
			program[i].Op = "nop"
			infinite, acc = IsInfiniteLoop(program)
			program[i].Op = "jmp"

		case "nop":
			program[i].Op = "jmp"
			infinite, acc = IsInfiniteLoop(program)
			program[i].Op = "nop"

		default:
			continue
		}

		if !infinite {
			break
		}
	}

	fmt.Println(acc)
}

func IsInfiniteLoop(program []cpus.HandheldInstruction) (bool, int) {
	seen := make(map[int]bool)

	var cpu cpus.HandheldCPU
	for !cpu.Halted {
		if seen[cpu.IP] {
			return true, 0
		}

		seen[cpu.IP] = true
		cpu.Step(program)
	}

	return false, cpu.ACC
}
