package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/cpus"
)

func main() {
	var code int
	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(),
		Input:  func() int { return 1 },
		Output: func(value int) { code = Max(code, value) },
	}
	cpu.Execute()

	fmt.Println(code)
}
