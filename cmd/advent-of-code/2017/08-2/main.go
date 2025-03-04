package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
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
				max = lib.Max(max, value)
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

	condition := func(register string, op string, limit int) func(map[string]int) bool {
		return func(registers map[string]int) bool {
			return conditions[op](registers[register], limit)
		}
	}

	return in.LinesToS(func(in in.Scanner[Instruction]) Instruction {
		var register, op, cregister, cop string
		var amount, camount int
		in.Scanf("%s %s %d if %s %s %d", &register, &op, &amount, &cregister, &cop, &camount)

		return Instruction{
			Register:  register,
			Op:        op,
			Amount:    amount,
			Condition: condition(cregister, cop, camount),
		}
	})
}
