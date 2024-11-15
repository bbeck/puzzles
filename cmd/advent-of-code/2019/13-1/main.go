package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz/cpus"
)

func main() {
	var count, blocks int
	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(2019, 13),
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
