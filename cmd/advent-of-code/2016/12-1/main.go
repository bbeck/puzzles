package main

import (
	"fmt"
	"strconv"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	program := InputToProgram()
	registers := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}
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
	return in.LinesTo(func(in *in.Scanner[Instruction]) Instruction {
		switch fields := in.Fields(); fields[0] {
		case "cpy":
			if n, err := strconv.Atoi(fields[1]); err == nil {
				return Instruction{OpCode: fields[0], Immediate: n, TargetRegister: fields[2]}
			}
			return Instruction{OpCode: fields[0], SourceRegister: fields[1], TargetRegister: fields[2]}

		case "inc":
			return Instruction{OpCode: fields[0], TargetRegister: fields[1]}

		case "dec":
			return Instruction{OpCode: fields[0], TargetRegister: fields[1]}

		case "jnz":
			if n, err := strconv.Atoi(fields[1]); err == nil {
				return Instruction{OpCode: fields[0], Immediate: n, Offset: ParseInt(fields[2])}
			}
			return Instruction{OpCode: fields[0], SourceRegister: fields[1], Offset: ParseInt(fields[2])}

		default:
			panic(fmt.Sprintf("unknown opcode: %s", fields[0]))
		}
	})
}
