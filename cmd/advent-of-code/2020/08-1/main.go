package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	program := InputToProgram()

	var pc, acc int
	var seen puz.Set[int]
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
	return puz.InputLinesTo(func(line string) Instruction {
		fields := strings.Fields(line)

		return Instruction{
			OpCode: fields[0],
			Arg:    puz.ParseInt(fields[1]),
		}
	})
}
