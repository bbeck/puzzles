package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	program := cpus.InputToHandheldProgram(2020, 8)
	seen := make(map[int]bool)

	var cpu cpus.HandheldCPU
	for {
		if seen[cpu.IP] {
			break
		}

		seen[cpu.IP] = true
		cpu.Step(program)
	}

	fmt.Println(cpu.ACC)
}
