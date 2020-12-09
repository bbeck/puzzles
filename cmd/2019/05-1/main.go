package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	memory := cpus.InputToIntcodeMemory(2019, 5)
	cpu := cpus.IntcodeCPU{
		Memory: memory,
		Input:  func(addr int) int { return 1 },
		Output: func(value int) {
			if value != 0 {
				fmt.Println(value)
			}
		},
	}
	cpu.Execute()
}
