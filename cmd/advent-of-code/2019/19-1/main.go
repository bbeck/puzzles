package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/cpus"
)

func main() {
	var memory = cpus.InputToIntcodeMemory()

	var count int
	for x := range 50 {
		for y := range 50 {
			inputs := make(chan int, 2)
			inputs <- x
			inputs <- y

			cpu := cpus.IntcodeCPU{
				Memory: memory.Copy(),
				Input:  func() int { return <-inputs },
				Output: func(value int) { count += value },
			}
			cpu.Execute()
		}
	}

	fmt.Println(count)
}
