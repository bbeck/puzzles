package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"

	"github.com/bbeck/advent-of-code/puz/cpus"
)

func main() {
	var code int
	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(),
		Input:  func() int { return 1 },
		Output: func(value int) { code = puz.Max(code, value) },
	}
	cpu.Execute()

	fmt.Println(code)
}
