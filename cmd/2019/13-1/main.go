package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	sequence := make(chan int)
	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(2019, 13),
		Output: func(value int) { sequence <- value },
		Halt:   func() { close(sequence) },
	}

	go cpu.Execute()

	screen := make(map[aoc.Point2D]int)
	for {
		x, ok := <-sequence
		if !ok {
			break
		}

		y := <-sequence
		id := <-sequence
		screen[aoc.Point2D{X: x, Y: y}] = id
	}

	var count int
	for _, id := range screen {
		if id == 2 {
			count++
		}
	}

	fmt.Printf("number of block tiles: %d\n", count)
}
