package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ipr, program := InputToProgram(2018, 19)
	registers := make([]int, 6)

	for ip := 0; ip >= 0 && ip < len(program); ip++ {
		registers[ipr] = ip
		instruction := program[ip]

		fmt.Printf("ip=%02d: %s %2d %2d %2d  |  [%s] -> ", ip, instruction.op, instruction.a, instruction.b, instruction.c, RegisterString(registers))

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
			if registers[instruction.a] == registers[instruction.b] {
				registers[instruction.c] = 1
			} else {
				registers[instruction.c] = 0
			}
		}

		fmt.Printf("[%s]\n", RegisterString(registers))

		ip = registers[ipr]
	}

	fmt.Printf("register 0: %d\n", registers[0])
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

func RegisterString(rs []int) string {
	widths := []int{1, 1, 8, 2, 3, 8}
	var builder strings.Builder
	for i, r := range rs {
		if i == 3 {
			continue
		}
		s := fmt.Sprintf("%%%dd ", widths[i])
		builder.WriteString(fmt.Sprintf(s, r))
	}
	return builder.String()
}
