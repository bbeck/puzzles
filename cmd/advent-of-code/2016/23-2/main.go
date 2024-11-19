package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	program := InputToProgram()

	// In my particular input, lines 3 through 9 form a nested loop that performs
	// a multiplication operation using repeated incrementing.
	//
	//        cpy a d        # d = a
	//        cpy 0 a        # a = 0
	//
	//                       # outer loop iterates a times
	// outer: cpy b c        # c = b
	//
	//                       # inner loop adds b to a
	// inner: inc a          # a++
	//        dec c          # c--
	//        jnz c -2       # if c != 0: goto inner
	//
	//        dec d          # d = d - 1
	//        jnz d -5       # if d != 0: goto outer
	//
	// To optimize this we'll replace these nested loops with a multiplication
	// instruction.  It's however important to be sure to keep the overall number
	// of instructions the same so that relative jumps still function as
	// expected.  To accomplish this we'll also introduce a no-op instruction.
	program[3] = Instruction{OpCode: "mul", Args: []string{"b", "a"}}
	program[4] = Instruction{OpCode: "cpy", Args: []string{"0", "c"}, Parsed: []int{0, 0}}
	program[5] = Instruction{OpCode: "cpy", Args: []string{"0", "d"}, Parsed: []int{0, 0}}
	program[6] = Instruction{OpCode: "nop"}
	program[7] = Instruction{OpCode: "nop"}
	program[8] = Instruction{OpCode: "nop"}
	program[9] = Instruction{OpCode: "nop"}

	registers := map[string]int{"a": 12, "b": 0, "c": 0, "d": 0}
	pc := 0

	reg := func(instruction Instruction, arg int) (string, error) {
		if _, ok := registers[instruction.Args[arg]]; ok {
			return instruction.Args[arg], nil
		}
		return "", fmt.Errorf("not a register: %s", instruction.Args[arg])
	}

	get := func(instruction Instruction, arg int) int {
		if value, ok := registers[instruction.Args[arg]]; ok {
			return value
		}

		return instruction.Parsed[arg]
	}

	for pc >= 0 && pc < len(program) {
		switch instruction := program[pc]; instruction.OpCode {
		case "nop":
			pc++

		case "cpy":
			if target, err := reg(instruction, 1); err == nil {
				registers[target] = get(instruction, 0)
			}
			pc++

		case "inc":
			if target, err := reg(instruction, 0); err == nil {
				registers[target]++
			}
			pc++

		case "dec":
			if target, err := reg(instruction, 0); err == nil {
				registers[target]--
			}
			pc++

		case "mul":
			if target, err := reg(instruction, 1); err == nil {
				registers[target] = registers[target] * get(instruction, 0)
			}
			pc++

		case "jnz":
			if get(instruction, 0) != 0 {
				pc += get(instruction, 1)
			} else {
				pc++
			}

		case "tgl":
			address := pc + get(instruction, 0)
			if address >= 0 && address < len(program) {
				switch target := &program[address]; target.OpCode {
				// Single argument instructions
				case "inc":
					target.OpCode = "dec"
				case "dec":
					target.OpCode = "inc"
				case "tgl":
					target.OpCode = "inc"

				// Two argument instructions
				case "cpy":
					target.OpCode = "jnz"
				case "jnz":
					target.OpCode = "cpy"
				}
			}
			pc++
		}
	}

	fmt.Println(registers["a"])
}

type Instruction struct {
	OpCode string
	Args   []string
	Parsed []int
}

func InputToProgram() []Instruction {
	return lib.InputLinesTo(func(line string) Instruction {
		fields := strings.Fields(line)
		opcode := fields[0]
		args := fields[1:]
		parsed := make([]int, len(args))

		for i, arg := range args {
			if n, err := strconv.Atoi(arg); err == nil {
				parsed[i] = n
			}
		}

		return Instruction{
			OpCode: opcode,
			Args:   args,
			Parsed: parsed,
		}
	})
}
