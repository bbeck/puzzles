package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz/cpus"
)

func main() {
	program := []string{
		// Jump if there's a hole @3 and land @4 and land @
		"NOT C T",
		"AND D T",
		"AND H T",
		"OR T J",

		// Jump if there's a hole @2 and land @4
		"NOT B T",
		"AND D T",
		"OR T J",

		// Jump if there's a hole @1
		"NOT A T",
		"OR T J",

		"RUN",
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
		Input:  func() int { return <-inputs },
		Output: func(value int) {
			if value > 255 {
				fmt.Println(value)
			}
		},
	}
	cpu.Execute()
}
