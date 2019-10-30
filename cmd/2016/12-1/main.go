package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	program := InputToProgram(2016, 12)
	registers := map[string]int{
		"a": 0, "b": 0, "c": 0, "d": 0,
	}
	pc := 0

	for {
		if pc >= len(program) {
			break
		}

		instruction := program[pc]
		switch instruction.name {
		case "cpy":
			if instruction.source != "" {
				registers[instruction.destination] = registers[instruction.source]
			} else {
				registers[instruction.destination] = instruction.immediate
			}
			pc++

		case "inc":
			registers[instruction.destination] = registers[instruction.destination] + 1
			pc++

		case "dec":
			registers[instruction.destination] = registers[instruction.destination] - 1
			pc++

		case "jnz":
			var test int
			if instruction.source != "" {
				test = registers[instruction.source]
			} else {
				test = instruction.immediate
			}

			if test != 0 {
				pc += instruction.offset
			} else {
				pc++
			}
		}
	}

	fmt.Printf("a: %d\n", registers["a"])
}

type Instruction struct {
	name        string
	source      string
	destination string
	immediate   int
	offset      int
}

func InputToProgram(year, day int) []Instruction {
	var program []Instruction
	for _, line := range aoc.InputToLines(year, day) {
		tokens := strings.Split(line, " ")
		name := tokens[0]

		switch name {
		case "cpy":
			if value, err := strconv.Atoi(tokens[1]); err == nil {
				// we had an immediate value not a register
				program = append(program, Instruction{
					name:        name,
					immediate:   value,
					destination: tokens[2],
				})
			} else {
				program = append(program, Instruction{
					name:        name,
					source:      tokens[1],
					destination: tokens[2],
				})
			}

		case "inc":
			program = append(program, Instruction{
				name:        name,
				destination: tokens[1],
			})

		case "dec":
			program = append(program, Instruction{
				name:        name,
				destination: tokens[1],
			})

		case "jnz":
			if value, err := strconv.Atoi(tokens[1]); err == nil {
				// we had an immediate value not a register
				program = append(program, Instruction{
					name:      name,
					immediate: value,
					offset:    aoc.ParseInt(tokens[2]),
				})
			} else {
				program = append(program, Instruction{
					name:   name,
					source: tokens[1],
					offset: aoc.ParseInt(tokens[2]),
				})
			}

		default:
			log.Fatalf("unrecognized instruction: %s", name)
		}
	}

	return program
}
