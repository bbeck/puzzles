package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	opcodes := InputToProgram(2019, 2)
	opcodes[1] = 12
	opcodes[2] = 2

	get := func(i int) int {
		return opcodes[opcodes[i]]
	}

	set := func(i int, value int) {
		opcodes[opcodes[i]] = value
	}

loop:
	for ip := 0; ; ip += 4 {
		switch opcodes[ip] {
		case 1:
			set(ip+3, get(ip+1)+get(ip+2))
		case 2:
			set(ip+3, get(ip+1)*get(ip+2))
		case 99:
			break loop
		}
	}

	fmt.Printf("position 0: %d\n", opcodes[0])
}

func InputToProgram(year, day int) []int {
	var opcodes []int
	for _, s := range strings.Split(aoc.InputToString(year, day), ",") {
		opcodes = append(opcodes, aoc.ParseInt(s))
	}

	return opcodes
}
