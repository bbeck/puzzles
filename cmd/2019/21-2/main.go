package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	program := []string{
		// If there's a hole right in front of me, jump, there's no other choice
		"NOT A T",
		"OR T J",

		// If there's a hole @2 and land @4, jump
		"NOT B T",
		"AND D T",
		"OR T J",

		// If there's a hole @3 and land @4, jump
		"NOT C T",
		"AND D T",
		"OR T J",

		// We're only allowed to jump if we have a space to move when we land
		// meaning there's land @5 or land @8
		"OR J T",
		"AND E T",
		"OR H T",
		"AND T J",

		"RUN",
	}
	if len(program) > 15 {
		log.Fatal("program too long")
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
				fmt.Println()
				fmt.Printf("output: %d\n", value)
			}
		},
	}
	cpu.Execute()
}
