package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/lib/cpus"
)

func main() {
	memory := cpus.InputToIntcodeMemory()
	memory[1] = 12
	memory[2] = 2

	cpu := cpus.IntcodeCPU{
		Memory: memory,
	}
	cpu.Execute()
	fmt.Println(memory[0])
}
