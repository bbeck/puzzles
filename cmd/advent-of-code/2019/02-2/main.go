package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz/cpus"
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

	fmt.Println(100*noun + verb)
}

func Execute(noun, verb int) int {
	memory := cpus.InputToIntcodeMemory(2019, 2)
	memory[1] = noun
	memory[2] = verb

	cpu := cpus.IntcodeCPU{
		Memory: memory,
	}
	cpu.Execute()

	return memory[0]
}
