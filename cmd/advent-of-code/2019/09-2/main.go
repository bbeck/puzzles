package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz/cpus"
)

func main() {
	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(2019, 9),
		Input:  func() int { return 2 },
		Output: func(value int) { fmt.Println(value) },
	}
	cpu.Execute()
}
