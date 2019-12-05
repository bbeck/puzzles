package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	memory := InputToProgram(2019, 5)
	Execute(memory)
}

func Execute(memory []int) {
	get := func(n int, mode int) int {
		switch mode {
		case 0: // position mode
			return memory[n]
		case 1: // immediate mode
			return n
		}

		log.Fatalf("don't know how to get n: %d, in mode: %d\n", n, mode)
		return -1
	}

	set := func(n int, value int, mode int) {
		switch mode {
		case 0: // position mode
			memory[n] = value
		default:
			log.Fatalf("don't know how to set n: %d, in mode: %d\n", n, mode)
		}
	}

loop:
	for ip := 0; ; {
		op := memory[ip] % 100
		aMode := (memory[ip] / 100) % 10
		bMode := (memory[ip] / 1000) % 10
		cMode := (memory[ip] / 10000) % 10

		switch op {
		case 1: // add
			a := memory[ip+1]
			b := memory[ip+2]
			c := memory[ip+3]
			set(c, get(a, aMode)+get(b, bMode), cMode)
			ip += 4

		case 2: // mul
			a := memory[ip+1]
			b := memory[ip+2]
			c := memory[ip+3]
			set(c, get(a, aMode)*get(b, bMode), cMode)
			ip += 4

		case 3: // input
			a := memory[ip+1]
			set(a, 1, aMode) // hardcoded to 1 for now
			ip += 2

		case 4: // output
			a := memory[ip+1]
			fmt.Printf("output: %d\n", get(a, aMode))
			ip += 2

		case 99:
			break loop
		}
	}
}

func InputToProgram(year, day int) []int {
	var opcodes []int
	for _, s := range strings.Split(aoc.InputToString(year, day), ",") {
		opcodes = append(opcodes, aoc.ParseInt(s))
	}

	return opcodes
}
