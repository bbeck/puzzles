package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib/in"
	"strconv"
)

func main() {
	registers := make(map[string]int)

	get := func(instruction Instruction, index int) int {
		arg := instruction.Args[index]
		if 'a' <= arg[0] && arg[0] <= 'z' {
			return registers[arg]
		}
		return instruction.Parsed[index]
	}

	program := InputToProgram()
	var pc, frequency int

loop:
	for pc >= 0 && pc < len(program) {
		switch instruction := program[pc]; instruction.OpCode {
		case "snd":
			frequency = get(instruction, 0)
			pc++

		case "set":
			x, y := instruction.Args[0], get(instruction, 1)
			registers[x] = y
			pc++

		case "add":
			x, y := instruction.Args[0], get(instruction, 1)
			registers[x] += y
			pc++

		case "mul":
			x, y := instruction.Args[0], get(instruction, 1)
			registers[x] *= y
			pc++

		case "mod":
			x, y := instruction.Args[0], get(instruction, 1)
			registers[x] %= y
			pc++

		case "rcv":
			x := get(instruction, 0)
			if x != 0 {
				break loop
			}
			pc++

		case "jgz":
			x, y := get(instruction, 0), get(instruction, 1)
			if x > 0 {
				pc += y
			} else {
				pc++
			}
		}
	}

	fmt.Println(frequency)
}

type Instruction struct {
	OpCode string
	Args   []string
	Parsed []int
}

func InputToProgram() []Instruction {
	return in.LinesToS(func(in in.Scanner[Instruction]) Instruction {
		var opcode = in.String()
		var args = in.Fields()

		var parsed = make([]int, len(args))
		for i, arg := range args {
			if n, err := strconv.Atoi(arg); err == nil {
				parsed[i] = n
			}
		}

		return Instruction{OpCode: opcode, Args: args, Parsed: parsed}
	})
}
