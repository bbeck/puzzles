package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"

	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	var code int
	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(2019, 5),
		Input:  func() int { return 1 },
		Output: func(value int) { code = aoc.Max(code, value) },
	}
	cpu.Execute()

	fmt.Println(code)
}
