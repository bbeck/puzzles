package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	program := InputToProgram()

	var loop bool
	var acc int
	for i := 0; i < len(program); i++ {
		if program[i].OpCode == "jmp" {
			program[i].OpCode = "nop"
			if loop, acc = IsInfiniteLoop(program); !loop {
				break
			}
			program[i].OpCode = "jmp"
		}

		if program[i].OpCode == "nop" {
			program[i].OpCode = "jmp"
			if loop, acc = IsInfiniteLoop(program); !loop {
				break
			}
			program[i].OpCode = "nop"
		}
	}

	fmt.Println(acc)
}

func IsInfiniteLoop(program []Instruction) (bool, int) {
	var pc, acc int
	var seen puz.Set[int]
	for {
		if !seen.Add(pc) {
			return true, acc
		}

		if pc >= len(program) {
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

	return false, acc
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
