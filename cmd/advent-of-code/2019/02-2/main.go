package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/cpus"
)

func main() {
	memory := cpus.InputToIntcodeMemory()
	var noun, verb int

loop:
	for noun = range 100 {
		for verb = range 100 {
			output := Execute(memory.Copy(), noun, verb)
			if output == 19690720 {
				break loop
			}
		}
	}

	fmt.Println(100*noun + verb)
}

func Execute(memory cpus.Memory, noun, verb int) int {
	memory[1] = noun
	memory[2] = verb

	cpu := cpus.IntcodeCPU{Memory: memory}
	cpu.Execute()
	return memory[0]
}
