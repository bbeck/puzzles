package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	program := InputToProgram()
	registers := map[string]int{"a": 0, "b": 0, "c": 1, "d": 0}
	pc := 0

	get := func(register string, immediate int) int {
		if register != "" {
			return registers[register]
		}
		return immediate
	}

	for pc >= 0 && pc < len(program) {
		switch instruction := program[pc]; instruction.OpCode {
		case "cpy":
			registers[instruction.TargetRegister] = get(instruction.SourceRegister, instruction.Immediate)
			pc++

		case "inc":
			registers[instruction.TargetRegister]++
			pc++

		case "dec":
			registers[instruction.TargetRegister]--
			pc++

		case "jnz":
			if get(instruction.SourceRegister, instruction.Immediate) != 0 {
				pc += instruction.Offset
			} else {
				pc++
			}
		}
	}

	fmt.Println(registers["a"])
}

type Instruction struct {
	OpCode         string
	TargetRegister string
	SourceRegister string
	Immediate      int
	Offset         int
}

func InputToProgram() []Instruction {
	return aoc.InputLinesTo(2016, 12, func(line string) (Instruction, error) {
		opcode, rest, _ := strings.Cut(line, " ")
		args := strings.Fields(rest)

		instruction := Instruction{OpCode: opcode}
		switch opcode {
		case "cpy":
			if n, err := strconv.Atoi(args[0]); err == nil {
				instruction.Immediate = n
			} else {
				instruction.SourceRegister = args[0]
			}
			instruction.TargetRegister = args[1]
		case "inc":
			instruction.TargetRegister = args[0]
		case "dec":
			instruction.TargetRegister = args[0]
		case "jnz":
			if n, err := strconv.Atoi(args[0]); err == nil {
				instruction.Immediate = n
			} else {
				instruction.SourceRegister = args[0]
			}
			instruction.Offset = aoc.ParseInt(args[1])
		}

		return instruction, nil
	})
}
