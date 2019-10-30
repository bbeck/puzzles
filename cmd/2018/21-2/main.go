package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	// By reading the program, it can be seen that the only time that register 0
	// is read is on instruction 28: eqrr 3 0 1.  We know that normally the
	// program will infinite loop.  What we need to do is determine what the last
	// comparison we encounter on instruction 28 is before we repeat ourselves.
	// To do this we'll remember each value of register 3 that we've checked in
	// the past, as well as the last value that we checked.  Once we check a given
	// value for the second time, we know we've begun to loop, so at that point
	// the last value we checked was the one that register 0 can take on and have
	// the most work performed.
	//
	// For my particular input this value is: 16311888
	ipr, program := InputToProgram(2018, 21)
	registers := make([]int, 6)

	seen := make(map[int]bool)
	last0 := 0

loop:
	for ip := 0; ip >= 0 && ip < len(program); ip++ {
		registers[ipr] = ip
		instruction := program[ip]

		switch instruction.op {
		case "addr":
			registers[instruction.c] = registers[instruction.a] + registers[instruction.b]
		case "addi":
			registers[instruction.c] = registers[instruction.a] + instruction.b
		case "mulr":
			registers[instruction.c] = registers[instruction.a] * registers[instruction.b]
		case "muli":
			registers[instruction.c] = registers[instruction.a] * instruction.b
		case "banr":
			registers[instruction.c] = registers[instruction.a] & registers[instruction.b]
		case "bani":
			registers[instruction.c] = registers[instruction.a] & instruction.b
		case "borr":
			registers[instruction.c] = registers[instruction.a] | registers[instruction.b]
		case "bori":
			registers[instruction.c] = registers[instruction.a] | instruction.b
		case "setr":
			registers[instruction.c] = registers[instruction.a]
		case "seti":
			registers[instruction.c] = instruction.a
		case "gtir":
			if instruction.a > registers[instruction.b] {
				registers[instruction.c] = 1
			} else {
				registers[instruction.c] = 0
			}
		case "gtri":
			if registers[instruction.a] > instruction.b {
				registers[instruction.c] = 1
			} else {
				registers[instruction.c] = 0
			}
		case "gtrr":
			if registers[instruction.a] > registers[instruction.b] {
				registers[instruction.c] = 1
			} else {
				registers[instruction.c] = 0
			}
		case "eqir":
			if instruction.a == registers[instruction.b] {
				registers[instruction.c] = 1
			} else {
				registers[instruction.c] = 0
			}
		case "eqri":
			if registers[instruction.a] == instruction.b {
				registers[instruction.c] = 1
			} else {
				registers[instruction.c] = 0
			}
		case "eqrr":
			// This is when we determine if we're going to loop again.
			if seen[registers[instruction.a]] {
				break loop
			}

			seen[registers[instruction.a]] = true
			last0 = registers[instruction.a]

			if registers[instruction.a] == registers[instruction.b] {
				registers[instruction.c] = 1
			} else {
				registers[instruction.c] = 0
			}
		}

		ip = registers[ipr]
	}

	fmt.Printf("last value before looping: %d\n", last0)
}

type Instruction struct {
	op      string
	a, b, c int
}

func InputToProgram(year, day int) (int, []Instruction) {
	var ip int
	var program []Instruction
	for _, line := range aoc.InputToLines(year, day) {
		if line[0] == '#' {
			if _, err := fmt.Sscanf(line, "#ip %d", &ip); err != nil {
				log.Fatalf("unable to parse ip: %s", line)
			}
			continue
		}

		var op string
		var a, b, c int
		if _, err := fmt.Sscanf(line, "%s %d %d %d", &op, &a, &b, &c); err != nil {
			log.Fatalf("unable to parse instruction: %s", line)
		}

		program = append(program, Instruction{op: op, a: a, b: b, c: c})
	}

	return ip, program
}
