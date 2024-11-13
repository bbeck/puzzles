package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	program := InputToProgram()
	registers := map[string]int{"a": 7, "b": 0, "c": 0, "d": 0}
	pc := 0

	reg := func(instruction Instruction, arg int) (string, error) {
		if _, ok := registers[instruction.Args[arg]]; ok {
			return instruction.Args[arg], nil
		}
		return "", fmt.Errorf("not a register: %s", instruction.Args[arg])
	}

	get := func(instruction Instruction, arg int) int {
		if value, ok := registers[instruction.Args[arg]]; ok {
			return value
		}

		return instruction.Parsed[arg]
	}

	for pc >= 0 && pc < len(program) {
		switch instruction := program[pc]; instruction.OpCode {
		case "cpy":
			if target, err := reg(instruction, 1); err == nil {
				registers[target] = get(instruction, 0)
			}
			pc++

		case "inc":
			if target, err := reg(instruction, 0); err == nil {
				registers[target]++
			}
			pc++

		case "dec":
			if target, err := reg(instruction, 0); err == nil {
				registers[target]--
			}
			pc++

		case "jnz":
			if get(instruction, 0) != 0 {
				pc += get(instruction, 1)
			} else {
				pc++
			}

		case "tgl":
			address := pc + get(instruction, 0)
			if address >= 0 && address < len(program) {
				switch target := &program[address]; target.OpCode {
				// Single argument instructions
				case "inc":
					target.OpCode = "dec"
				case "dec":
					target.OpCode = "inc"
				case "tgl":
					target.OpCode = "inc"

				// Two argument instructions
				case "cpy":
					target.OpCode = "jnz"
				case "jnz":
					target.OpCode = "cpy"
				}
			}
			pc++
		}
	}

	fmt.Println(registers["a"])
}

type Instruction struct {
	OpCode string
	Args   []string
	Parsed []int
}

func InputToProgram() []Instruction {
	return puz.InputLinesTo(2016, 23, func(line string) Instruction {
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
