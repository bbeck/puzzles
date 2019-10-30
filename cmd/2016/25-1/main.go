package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	program := InputToProgram(2016, 25)
	for i, instruction := range program {
		fmt.Printf("%2d: %+v\n", i, instruction)
	}

	var found bool
	var a int
	for a = 1; !found; a++ {
		fmt.Printf("testing a=%d...", a)

		count := 0
		expected := 0
		check := func(value int) bool {
			count++

			if count > 1000 {
				found = true
				return true
			}

			if value != expected {
				return true
			}

			expected = 1 - value
			return false
		}

		run(program, map[string]int{"a": a}, check)

		if found {
			fmt.Println("FOUND!")
		} else {
			fmt.Println("failed")
		}
	}

	fmt.Printf("a: %d\n", a)
}

func run(program []Instruction, registers map[string]int, done func(int) bool) {
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

		case "out":
			var value int
			if instruction.source == "" {
				value = instruction.immediate
			} else {
				value = registers[instruction.source]
			}

			if done(value) {
				return
			}

			pc++
		}
	}
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

		case "out":
			if value, err := strconv.Atoi(tokens[1]); err == nil {
				// we had an immediate value not a register
				program = append(program, Instruction{
					name:      name,
					immediate: value,
				})
			} else {
				program = append(program, Instruction{
					name:   name,
					source: tokens[1],
				})
			}

		default:
			log.Fatalf("unrecognized instruction: %s", name)
		}
	}

	return program
}
