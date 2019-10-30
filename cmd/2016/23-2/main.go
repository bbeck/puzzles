package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

const debug = false

func main() {
	program := InputToProgram(2016, 23)
	registers := map[string]int{
		"a": 12, "b": 0, "c": 0, "d": 0,
	}
	pc := 0

	for n := 0; ; n++ {
		if pc >= len(program) {
			break
		}

		instruction := program[pc]
		if debug {
			fmt.Println("registers:")
			fmt.Printf("  a: %d, b: %d, c: %d, d: %d\n", registers["a"], registers["b"], registers["c"], registers["d"])
			fmt.Println()

			fmt.Println("program:")
			for i := 0; i < len(program); i++ {
				if pc == i {
					fmt.Printf("> %+v\n", program[i])
				} else {
					fmt.Printf("  %+v\n", program[i])
				}
			}
			fmt.Println()
			fmt.Println("=========================================")
			fmt.Println()
		}

		switch instruction.name {
		case "cpy":
			var value int
			if immediate, err := strconv.Atoi(instruction.args[0]); err == nil {
				value = immediate
			} else {
				value = registers[instruction.args[0]]
			}

			if _, ok := registers[instruction.args[1]]; ok {
				registers[instruction.args[1]] = value
			}
			pc++

		case "inc":
			registers[instruction.args[0]]++
			pc++

		case "dec":
			registers[instruction.args[0]]--
			pc++

		case "jnz":
			var test int
			if immediate, err := strconv.Atoi(instruction.args[0]); err == nil {
				test = immediate
			} else {
				test = registers[instruction.args[0]]
			}

			var offset int
			if immediate, err := strconv.Atoi(instruction.args[1]); err == nil {
				offset = immediate
			} else {
				offset = registers[instruction.args[1]]
			}

			if test != 0 {
				pc += offset
			} else {
				pc++
			}

		case "tgl":
			// toggle the instruction <source> away
			// cpy -> jnz
			// dec -> inc
			// inc -> dec
			// jnz -> cpy
			// tgl -> inc
			index := pc + registers[instruction.args[0]]
			if index >= 0 && index < len(program) {
				instr := program[index]
				switch instr.name {
				case "cpy":
					instr.name = "jnz"

				case "dec":
					instr.name = "inc"

				case "inc":
					instr.name = "dec"

				case "jnz":
					instr.name = "cpy"

				case "tgl":
					instr.name = "inc"
				}
			}

			pc++
		}
	}

	fmt.Printf("a: %d\n", registers["a"])
}

type Instruction struct {
	name string
	args []string
}

func (instr *Instruction) String() string {
	var builder strings.Builder
	builder.WriteString(instr.name)
	builder.WriteString(" ")
	builder.WriteString(strings.Join(instr.args, " "))
	return builder.String()
}

func InputToProgram(year, day int) []*Instruction {
	var program []*Instruction
	for _, line := range aoc.InputToLines(year, day) {
		tokens := strings.Split(line, " ")
		name := tokens[0]

		program = append(program, &Instruction{
			name: name,
			args: tokens[1:],
		})
	}

	return program
}
