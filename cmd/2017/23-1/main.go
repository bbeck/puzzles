package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	program := InputToProgram(2017, 23)
	registers := make(map[string]int)

	// get the immediate value or the value of the register
	get := func(i Instruction) int {
		if i.register == "" {
			return i.immediate
		}

		return registers[i.register]
	}

	// get the test value for a jgz
	test := func(i Instruction) int {
		v, err := strconv.Atoi(i.target)
		if err == nil {
			return v
		}

		return registers[i.target]
	}

	var muls int
	for pc := 0; pc >= 0 && pc < len(program); {
		instruction := program[pc]

		switch instruction.op {
		case "set":
			registers[instruction.target] = get(instruction)
			pc++

		case "sub":
			registers[instruction.target] -= get(instruction)
			pc++

		case "mul":
			registers[instruction.target] *= get(instruction)
			muls++
			pc++

		case "jnz":
			if test(instruction) != 0 {
				pc += instruction.immediate
			} else {
				pc++
			}
		}
	}

	fmt.Printf("mul count: %d\n", muls)
}

type Instruction struct {
	op        string
	target    string
	immediate int
	register  string
}

func (i Instruction) String() string {
	if i.register != "" {
		return fmt.Sprintf("%s %s %s", i.op, i.target, i.register)
	} else {
		return fmt.Sprintf("%s %s %d", i.op, i.target, i.immediate)
	}
}

func InputToProgram(year, day int) []Instruction {
	// parse an argument as either an immediate or the register it refers to
	parse := func(s string) (int, string) {
		immediate, err := strconv.Atoi(s)
		if err == nil {
			return immediate, ""
		}

		return 0, s
	}

	var program []Instruction
	for _, line := range aoc.InputToLines(year, day) {
		tokens := strings.Split(line, " ")

		var target string
		var immediate int
		var register string

		switch len(tokens) {
		case 2:
			immediate, register = parse(tokens[1])

		case 3:
			target = tokens[1]
			immediate, register = parse(tokens[2])
		}

		program = append(program, Instruction{
			op:        tokens[0],
			target:    target,
			immediate: immediate,
			register:  register,
		})
	}

	return program
}
