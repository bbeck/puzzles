package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib"
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
	return lib.InputLinesTo(func(line string) Instruction {
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "+", "")
		opcode, rest, _ := strings.Cut(line, " ")
		args := strings.Fields(rest)

		instruction := Instruction{OpCode: opcode}
		switch opcode {
		case "jmp":
			instruction.Offset = lib.ParseInt(args[0])
		case "jie":
			instruction.Register = args[0]
			instruction.Offset = lib.ParseInt(args[1])
		case "jio":
			instruction.Register = args[0]
			instruction.Offset = lib.ParseInt(args[1])
		default:
			instruction.Register = args[0]
		}

		return instruction
	})
}
