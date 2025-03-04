package main

import (
	"fmt"
	"strconv"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	program := InputToProgram()

	var a int
	var found bool
	for !found {
		a++
		registers := map[string]int{"a": a, "b": 0, "c": 0, "d": 0}

		var count int
		run(program, registers, func(x int) bool {
			if x != (count % 2) {
				found = false
				return true
			}

			count++
			if count > 100 {
				found = true
				return true
			}

			return false
		})
	}

	fmt.Println(a)
}

func run(program []Instruction, registers map[string]int, out func(int) bool) {
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

	pc := 0
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

		case "out":
			if out(get(instruction, 0)) {
				return
			}
			pc++
		}
	}
}

type Instruction struct {
	OpCode string
	Args   []string
	Parsed []int
}

func InputToProgram() []Instruction {
	return in.LinesTo(func(in *in.Scanner[Instruction]) Instruction {
		var opcode = in.String()
		var args = in.Fields()

		var parsed = make([]int, len(args))
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
