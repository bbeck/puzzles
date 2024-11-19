package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/cpus"
)

func main() {
	program := []string{
		// Jump if there's a hole @1
		"NOT A T",
		"OR T J",

		// Jump if there's a hole @3 and land @4
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
		Memory: cpus.InputToIntcodeMemory(),
		Input:  func() int { return <-inputs },
		Output: func(value int) {
			if value > 255 {
				fmt.Println(value)
			}
		},
	}
	cpu.Execute()
}
