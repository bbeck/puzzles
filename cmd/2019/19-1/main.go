package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	var sum int
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if Affected(x, y) {
				sum++
			}
		}
	}

	fmt.Println("num affected:", sum)
}

func Affected(x, y int) bool {
	inputs := make(chan int, 2)
	inputs <- x
	inputs <- y
	close(inputs)

	var output int
	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(2019, 19),
		Input:  func(addr int) int { return <-inputs },
		Output: func(value int) { output = value },
	}
	cpu.Execute()

	return output == 1
}
