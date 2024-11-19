package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	program := InputToProgram()

	var pc, acc int
	var seen lib.Set[int]
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
	return lib.InputLinesTo(func(line string) Instruction {
		fields := strings.Fields(line)

		return Instruction{
			OpCode: fields[0],
			Arg:    lib.ParseInt(fields[1]),
		}
	})
}
