package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	program := []string{
		// If there's a hole right in front of me, jump
		"NOT A T",
		"OR T J",

		// If there's a hole 3 steps in front and land 4 steps
		"NOT C T",
		"AND D T",
		"OR T J",

		"WALK",
	}
	inputs := make(chan int, 1024)
	for _, line := range program {
		for _, c := range line {
			inputs <- int(c)
		}
		inputs <- '\n'
	}
	close(inputs)

	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(2019, 21),
		Input:  func(addr int) int { return <-inputs },
		Output: func(value int) {
			if value <= 255 {
				fmt.Printf("%c", value)
			} else {
				fmt.Printf("output: %d\n", value)
			}
		},
	}
	cpu.Execute()
}
