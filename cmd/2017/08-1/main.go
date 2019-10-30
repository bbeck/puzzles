package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	program := InputToProgram(2017, 8)
	registers := make(map[string]int)

	for pc := 0; pc < len(program); pc++ {
		instruction := program[pc]

		if instruction.relation(registers[instruction.check], instruction.limit) {
			switch instruction.op {
			case "inc":
				registers[instruction.target] += instruction.offset

			case "dec":
				registers[instruction.target] -= instruction.offset
			}
		}
	}

	largest := func(registers map[string]int) int {
		max := math.MinInt64
		for _, value := range registers {
			if value > max {
				max = value
			}
		}

		return max
	}

	fmt.Printf("largest register value: %d\n", largest(registers))
}

type Instruction struct {
	target   string
	op       string
	offset   int
	check    string
	relation func(int, int) bool
	limit    int
}

func (i Instruction) String() string {
	return fmt.Sprintf("%s %s %d if %s %s %d", i.target, i.op, i.offset, i.check, i.relation, i.limit)
}

func InputToProgram(year, day int) []Instruction {
	relations := map[string]func(int, int) bool{
		">":  func(a int, b int) bool { return a > b },
		">=": func(a int, b int) bool { return a >= b },
		"<":  func(a int, b int) bool { return a < b },
		"<=": func(a int, b int) bool { return a <= b },
		"==": func(a int, b int) bool { return a == b },
		"!=": func(a int, b int) bool { return a != b },
	}

	var program []Instruction
	for _, line := range aoc.InputToLines(year, day) {
		tokens := strings.Split(line, " ")

		program = append(program, Instruction{
			target:   tokens[0],
			op:       tokens[1],
			offset:   aoc.ParseInt(tokens[2]),
			check:    tokens[4],
			relation: relations[tokens[5]],
			limit:    aoc.ParseInt(tokens[6]),
		})
	}

	return program
}
