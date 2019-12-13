package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	sequence := make(chan int)
	cpu := &CPU{
		memory: InputToMemory(2019, 13),
		output: func(value int) { sequence <- value },
		halt:   func() { close(sequence) },
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
		screen[aoc.Point2D{x, y}] = id
	}

	var count int
	for _, id := range screen {
		if id == 2 {
			count++
		}
	}

	fmt.Printf("number of block tiles: %d\n", count)
}
