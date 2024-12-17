package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	registers, program := InputToRegistersAndProgram()
	A, B, C := registers["A"], registers["B"], registers["C"]

	combo := func(i int) int {
		switch i {
		case 4:
			return A
		case 5:
			return B
		case 6:
			return C
		default:
			return i
		}
	}

	var output []int
	var ip int
	for {
		if ip >= len(program) {
			break
		}

		op, operand := program[ip], program[ip+1]
		switch op {
		case 0: // adv
			A = A / Pow(2, uint(combo(operand)))
		case 1: // bxl
			B = B ^ operand
		case 2: // bst
			B = combo(operand) % 8
		case 3: // jnz
			if A != 0 {
				ip = operand
				continue
			}
		case 4: // bxc
			B = B ^ C
		case 5: // out
			output = append(output, combo(operand)%8)
		case 6: // bdv
			B = A / Pow(2, uint(combo(operand)))
		case 7: // cdv
			C = A / Pow(2, uint(combo(operand)))
		}

		ip += 2
	}

	var ns []string
	for _, n := range output {
		ns = append(ns, fmt.Sprintf("%d", n))
	}
	fmt.Println(strings.Join(ns, ","))
}

func InputToRegistersAndProgram() (map[string]int, []int) {
	registers := make(map[string]int)
	var program []int

	for _, line := range InputToLines() {
		if line == "" {
			continue
		}

		line = strings.ReplaceAll(line, ":", "")
		line = strings.ReplaceAll(line, ",", " ")
		fields := strings.Fields(line)

		switch fields[0] {
		case "Register":
			registers[fields[1]] = ParseInt(fields[2])
		case "Program":
			for _, s := range fields[1:] {
				program = append(program, ParseInt(s))
			}
		}
	}

	return registers, program
}
