package cpus

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

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
