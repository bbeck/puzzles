package main

import (
	"fmt"
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

	cpu := &CPU{
		memory: InputToMemory(2019, 21),
		input:  func(addr int) int { return <-inputs },
		output: func(value int) {
			if value <= 255 {
				fmt.Printf("%c", value)
			} else {
				fmt.Println()
				fmt.Printf("output: %d\n", value)
			}
		},
		halt: nil,
	}
	cpu.Execute()
}
