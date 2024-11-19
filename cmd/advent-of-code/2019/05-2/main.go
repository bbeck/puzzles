package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/cpus"
)

func main() {
	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(),
		Input:  func() int { return 5 },
		Output: func(value int) { fmt.Println(value) },
	}
	cpu.Execute()
}
