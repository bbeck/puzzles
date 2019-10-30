package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	program := InputToProgram(2015, 23)
	pc := 0
	registers := map[string]uint{
		"a": 1,
	}

	for {
		if pc >= len(program) {
			break
		}

		instruction := program[pc]
		switch instruction.name {
		case "hlf":
			registers[instruction.register] /= 2
			pc++

		case "tpl":
			registers[instruction.register] *= 3
			pc++

		case "inc":
			registers[instruction.register]++
			pc++

		case "jmp":
			pc += instruction.offset

		case "jie":
			if registers[instruction.register]%2 == 0 {
				pc += instruction.offset
			} else {
				pc++
			}

		case "jio":
			if registers[instruction.register] == 1 {
				pc += instruction.offset
			} else {
				pc++
			}
		}
	}

	fmt.Printf("a: %d, b: %d\n", registers["a"], registers["b"])
}

type Instruction struct {
	name     string
	register string
	offset   int
}

func InputToProgram(year, day int) []Instruction {
	program := make([]Instruction, 0)
	for _, line := range aoc.InputToLines(year, day) {
		tokens := strings.Split(line, " ")
		name := tokens[0]

		switch name {
		case "hlf":
			program = append(program, Instruction{
				name:     name,
				register: ParseRegister(tokens[1]),
			})

		case "tpl":
			program = append(program, Instruction{
				name:     name,
				register: ParseRegister(tokens[1]),
			})

		case "inc":
			program = append(program, Instruction{
				name:     name,
				register: ParseRegister(tokens[1]),
			})

		case "jmp":
			program = append(program, Instruction{
				name:   name,
				offset: ParseOffset(tokens[1]),
			})

		case "jie":
			program = append(program, Instruction{
				name:     name,
				register: ParseRegister(tokens[1]),
				offset:   ParseOffset(tokens[2]),
			})

		case "jio":
			program = append(program, Instruction{
				name:     name,
				register: ParseRegister(tokens[1]),
				offset:   ParseOffset(tokens[2]),
			})
		}
	}

	return program
}

func ParseRegister(s string) string {
	return strings.ReplaceAll(s, ",", "")
}

func ParseOffset(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("unable to parse offset: %s", s)
	}

	return n
}
