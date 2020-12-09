package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(2019, 9),
		Input:  func(addr int) int { return 2 },
		Output: func(value int) { fmt.Println(value) },
	}

	cpu.Execute()
}
