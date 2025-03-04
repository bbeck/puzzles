package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib/in"
	"strconv"
)

func main() {
	chans := []chan int{make(chan int, 1024), make(chan int, 1024)}

	var count int
	send0 := func(value int) { count++; chans[0] <- value }
	send1 := func(value int) { chans[1] <- value }

	program := InputToProgram()
	cpu0 := CPU(0, program, send1, chans[0])
	cpu1 := CPU(1, program, send0, chans[1])

	for {
		if cpu0() && cpu1() {
			break
		}
	}
	fmt.Println(count)
}

type StepFunc func() bool
type SendFunc func(int)

func CPU(id int, program []Instruction, send SendFunc, recv <-chan int) StepFunc {
	registers := map[string]int{"p": id}

	get := func(instruction Instruction, index int) int {
		arg := instruction.Args[index]
		if 'a' <= arg[0] && arg[0] <= 'z' {
			return registers[arg]
		}
		return instruction.Parsed[index]
	}

	pc := 0

	return func() bool {
		if pc < 0 || pc >= len(program) {
			return true
		}

		switch instruction := program[pc]; instruction.OpCode {
		case "snd":
			x := get(instruction, 0)
			send(x)
			pc++

		case "set":
			x, y := instruction.Args[0], get(instruction, 1)
			registers[x] = y
			pc++

		case "add":
			x, y := instruction.Args[0], get(instruction, 1)
			registers[x] += y
			pc++

		case "mul":
			x, y := instruction.Args[0], get(instruction, 1)
			registers[x] *= y
			pc++

		case "mod":
			x, y := instruction.Args[0], get(instruction, 1)
			registers[x] %= y
			pc++

		case "rcv":
			select {
			case value := <-recv:
				x := instruction.Args[0]
				registers[x] = value
				pc++
			default:
				// Value isn't ready yet, indicate we're blocked.  Intentionally don't
				// increment pc in order to try this instruction again next time.
				return true
			}

		case "jgz":
			x, y := get(instruction, 0), get(instruction, 1)
			if x > 0 {
				pc += y
			} else {
				pc++
			}
		}

		return false
	}
}

type Instruction struct {
	OpCode string
	Args   []string
	Parsed []int
}

func InputToProgram() []Instruction {
	return in.LinesToS(func(in in.Scanner[Instruction]) Instruction {
		var opcode = in.String()
		var args = in.Fields()

		var parsed = make([]int, len(args))
		for i, arg := range args {
			if n, err := strconv.Atoi(arg); err == nil {
				parsed[i] = n
			}
		}

		return Instruction{OpCode: opcode, Args: args, Parsed: parsed}
	})
}
