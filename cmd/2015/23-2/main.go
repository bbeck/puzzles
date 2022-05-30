package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	program := InputToProgram()
	registers := map[string]uint{"a": 1}

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
	return aoc.InputLinesTo(2015, 23, func(line string) (Instruction, error) {
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "+", "")
		opcode, rest, _ := strings.Cut(line, " ")
		args := strings.Fields(rest)

		instruction := Instruction{OpCode: opcode}
		switch opcode {
		case "jmp":
			instruction.Offset = aoc.ParseInt(args[0])
		case "jie":
			instruction.Register = args[0]
			instruction.Offset = aoc.ParseInt(args[1])
		case "jio":
			instruction.Register = args[0]
			instruction.Offset = aoc.ParseInt(args[1])
		default:
			instruction.Register = args[0]
		}

		return instruction, nil
	})
}
