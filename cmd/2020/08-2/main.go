package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	program := InputToInstructions(2020, 8)
	for i, instruction := range program {
		if instruction.opcode == "acc" {
			continue
		}

		var isInfinite bool
		var acc int

		if instruction.opcode == "jmp" {
			program[i].opcode = "nop"
			isInfinite, acc = IsInfiniteLoop(program)
			program[i].opcode = "jmp"
		} else if instruction.opcode == "nop" {
			program[i].opcode = "jmp"
			isInfinite, acc = IsInfiniteLoop(program)
			program[i].opcode = "nop"
		}

		if !isInfinite {
			fmt.Println(acc)
			break
		}
	}
}

func IsInfiniteLoop(program []Instruction) (bool, int) {
	pc := 0
	acc := 0
	seen := make(map[int]bool)

	for pc < len(program) {
		if _, found := seen[pc]; found {
			return true, 0
		}

		seen[pc] = true
		instruction := program[pc]

		switch instruction.opcode {
		case "acc":
			acc += instruction.argument
			pc++
		case "jmp":
			pc += instruction.argument
		case "nop":
			pc++
		}
	}

	return false, acc
}

type Instruction struct {
	opcode   string
	argument int
}

func InputToInstructions(year, day int) []Instruction {
	var instructions []Instruction
	for _, line := range aoc.InputToLines(year, day) {
		var opcode string
		var argument int

		_, err := fmt.Sscanf(line, "%s %d", &opcode, &argument)
		if err != nil {
			log.Fatalf("unable to parse line: %s", line)
		}

		instructions = append(instructions, Instruction{
			opcode:   opcode,
			argument: argument,
		})
	}

	return instructions
}
