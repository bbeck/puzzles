package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/cpus"
)

func main() {
	var count, blocks int
	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(),
		Output: func(value int) {
			count++
			if count%3 == 0 && value == 2 {
				blocks++
			}
		},
	}
	cpu.Execute()

	fmt.Println(blocks)
}
