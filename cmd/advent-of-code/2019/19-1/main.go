package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz/cpus"
)

func main() {
	var count int

	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			inputs := make(chan int, 2)
			inputs <- x
			inputs <- y

			cpu := cpus.IntcodeCPU{
				Memory: cpus.InputToIntcodeMemory(),
				Input:  func() int { return <-inputs },
				Output: func(value int) { count += value },
			}
			cpu.Execute()
		}
	}

	fmt.Println(count)
}
