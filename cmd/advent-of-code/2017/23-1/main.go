package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	registers := map[string]int{}

	get := func(instruction Instruction, arg int) int {
		if value, ok := registers[instruction.Args[arg]]; ok {
			return value
		}

		return instruction.Parsed[arg]
	}

	var count int

	program := InputToProgram()
	pc := 0
	for pc >= 0 && pc < len(program) {
		switch instruction := program[pc]; instruction.OpCode {
		case "set":
			x := instruction.Args[0]
			y := get(instruction, 1)
			registers[x] = y
			pc++

		case "sub":
			x := instruction.Args[0]
			y := get(instruction, 1)
			registers[x] -= y
			pc++

		case "mul":
			x := instruction.Args[0]
			y := get(instruction, 1)
			registers[x] *= y
			pc++
			count++

		case "jnz":
			x := get(instruction, 0)
			y := get(instruction, 1)
			if x != 0 {
				pc += y
			} else {
				pc++
			}
		}
	}

	fmt.Println(count)
}

type Instruction struct {
	OpCode string
	Args   []string
	Parsed []int
}

func InputToProgram() []Instruction {
	return puz.InputLinesTo(2017, 23, func(line string) Instruction {
		fields := strings.Fields(line)
		opcode := fields[0]
		args := fields[1:]
		parsed := make([]int, len(args))

		for i, arg := range args {
			if n, err := strconv.Atoi(arg); err == nil {
				parsed[i] = n
			}
		}

		return Instruction{
			OpCode: opcode,
			Args:   args,
			Parsed: parsed,
		}
	})
}
