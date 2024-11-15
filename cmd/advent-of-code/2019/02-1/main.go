package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz/cpus"
)

func main() {
	memory := cpus.InputToIntcodeMemory(2019, 2)
	memory[1] = 12
	memory[2] = 2

	cpu := cpus.IntcodeCPU{
		Memory: memory,
	}
	cpu.Execute()
	fmt.Println(memory[0])
}
