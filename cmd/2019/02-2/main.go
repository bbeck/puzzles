package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var noun, verb int

loop:
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			output := Execute(noun, verb)
			if output == 19690720 {
				break loop
			}
		}
	}

	fmt.Printf("noun: %d, verb: %d, output: %d\n", noun, verb, 100*noun+verb)
}

func Execute(noun, verb int) int {
	memory := InputToProgram(2019, 2)
	memory[1] = noun
	memory[2] = verb

	get := func(i int) int {
		return memory[memory[i]]
	}

	set := func(i int, value int) {
		memory[memory[i]] = value
	}

loop:
	for ip := 0; ; ip += 4 {
		switch memory[ip] {
		case 1:
			set(ip+3, get(ip+1)+get(ip+2))
		case 2:
			set(ip+3, get(ip+1)*get(ip+2))
		case 99:
			break loop
		}
	}

	return memory[0]
}

func InputToProgram(year, day int) []int {
	var opcodes []int
	for _, s := range strings.Split(aoc.InputToString(year, day), ",") {
		opcodes = append(opcodes, aoc.ParseInt(s))
	}

	return opcodes
}
