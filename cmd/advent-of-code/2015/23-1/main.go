package main

import (
	"fmt"
	"strconv"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	program := InputToProgram()

	registers := make(map[string]uint)

	var pc int
	for {
		if pc >= len(program) {
			break
		}

		switch instruction := program[pc]; instruction.OpCode {
		case "hlf":
			registers[instruction.Register] /= 2
			pc++
		case "tpl":
			registers[instruction.Register] *= 3
			pc++
		case "inc":
			registers[instruction.Register] += 1
			pc++
		case "jmp":
			pc += instruction.Offset
		case "jie":
			if registers[instruction.Register]%2 == 0 {
				pc += instruction.Offset
			} else {
				pc++
			}
		case "jio":
			if registers[instruction.Register] == 1 {
				pc += instruction.Offset
			} else {
				pc++
			}
		}
	}

	fmt.Println(registers["b"])
}

type Instruction struct {
	OpCode   string
	Register string
	Offset   int
}

func InputToProgram() []Instruction {
	return in.LinesTo(func(in *in.Scanner[Instruction]) Instruction {
		var opcode = in.String()
		arg1, arg2 := in.Cut(", ")

		var register string
		var offset int
		if n, err := strconv.Atoi(arg1); err == nil {
			offset = n
		} else {
			register = arg1
		}
		if n, err := strconv.Atoi(arg2); err == nil {
			offset = n
		}

		return Instruction{OpCode: opcode, Register: register, Offset: offset}
	})
}
