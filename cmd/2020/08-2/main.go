package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	program := InputToHandheldProgram(2020, 8)

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

func IsInfiniteLoop(program []HandheldInstruction) (bool, int) {
	seen := make(map[int]bool)

	var cpu HandheldCPU
	for !cpu.Halted {
		if seen[cpu.IP] {
			return true, 0
		}

		seen[cpu.IP] = true
		cpu.Step(program)
	}

	return false, cpu.ACC
}

type HandheldCPU struct {
	IP     int
	ACC    int
	Halted bool
}

func (cpu *HandheldCPU) RunProgram(program []HandheldInstruction) {
	for !cpu.Halted {
		cpu.Step(program)
	}
}

func (cpu *HandheldCPU) Step(program []HandheldInstruction) {
	if cpu.IP < 0 || cpu.IP >= len(program) {
		cpu.Halted = true
		return
	}

	instruction := program[cpu.IP]
	switch instruction.Op {
	case "acc":
		cpu.ACC += instruction.Argument
		cpu.IP++
	case "jmp":
		cpu.IP += instruction.Argument
	case "nop":
		cpu.IP++
	}
}

type HandheldInstruction struct {
	Op       string
	Argument int
}

func InputToHandheldProgram(year, day int) []HandheldInstruction {
	var instructions []HandheldInstruction
	for _, line := range aoc.InputToLines(year, day) {
		var op string
		var argument int

		_, err := fmt.Sscanf(line, "%s %d", &op, &argument)
		if err != nil {
			log.Fatalf("unable to parse line: %s", line)
		}

		instructions = append(instructions, HandheldInstruction{
			Op:       op,
			Argument: argument,
		})
	}

	return instructions
}
