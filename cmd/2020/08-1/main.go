package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	program := InputToProgram()

	var pc, acc int
	var seen aoc.Set[int]
	for {
		if !seen.Add(pc) {
			break
		}

		switch instruction := program[pc]; instruction.OpCode {
		case "acc":
			acc += instruction.Arg
			pc++

		case "jmp":
			pc += instruction.Arg

		case "nop":
			pc++
		}
	}

	fmt.Println(acc)
}

type Instruction struct {
	OpCode string
	Arg    int
}

func InputToProgram() []Instruction {
	return aoc.InputLinesTo(2020, 8, func(line string) (Instruction, error) {
		fields := strings.Fields(line)

		return Instruction{
			OpCode: fields[0],
			Arg:    aoc.ParseInt(fields[1]),
		}, nil
	})
}
