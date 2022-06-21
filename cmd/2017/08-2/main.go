package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	registers := make(map[string]int)

	var max int
	for _, instruction := range InputToProgram() {
		if instruction.Condition(registers) {
			switch instruction.Op {
			case "inc":
				registers[instruction.Register] += instruction.Amount
			case "dec":
				registers[instruction.Register] -= instruction.Amount
			}
			
			for _, value := range registers {
				max = aoc.Max(max, value)
			}
		}
	}

	fmt.Println(max)
}

type Instruction struct {
	Register  string
	Op        string
	Amount    int
	Condition func(map[string]int) bool
}

func InputToProgram() []Instruction {
	conditions := map[string]func(int, int) bool{
		">":  func(a int, b int) bool { return a > b },
		">=": func(a int, b int) bool { return a >= b },
		"<":  func(a int, b int) bool { return a < b },
		"<=": func(a int, b int) bool { return a <= b },
		"==": func(a int, b int) bool { return a == b },
		"!=": func(a int, b int) bool { return a != b },
	}

	condition := func(register string, op string, lim string) func(map[string]int) bool {
		limit := aoc.ParseInt(lim)

		return func(registers map[string]int) bool {
			return conditions[op](registers[register], limit)
		}
	}

	return aoc.InputLinesTo(2017, 8, func(line string) (Instruction, error) {
		fields := strings.Fields(line)

		return Instruction{
			Register:  fields[0],
			Op:        fields[1],
			Amount:    aoc.ParseInt(fields[2]),
			Condition: condition(fields[4], fields[5], fields[6]),
		}, nil
	})
}
