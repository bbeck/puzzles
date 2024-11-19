package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"

	"github.com/bbeck/advent-of-code/lib/cpus"
)

const N = 5

func main() {
	var best int
	lib.EnumeratePermutations(N, func(settings []int) bool {
		best = lib.Max(best, TestSettings(settings))
		return false
	})
	fmt.Println(best)
}

func TestSettings(settings []int) int {
	var chans [N + 1]chan int
	for i := 0; i < len(chans); i++ {
		chans[i] = make(chan int, 2)
	}

	// Send the settings into the inputs
	for i, setting := range settings {
		chans[i] <- setting
	}

	// First amplifier's input is hardcoded to 0
	chans[0] <- 0

	var amps [N]cpus.IntcodeCPU
	for i := 0; i < N; i++ {
		i := i
		amps[i].Memory = cpus.InputToIntcodeMemory()
		amps[i].Input = func() int { return <-chans[i] }
		amps[i].Output = func(value int) { chans[i+1] <- value }
		go amps[i].Execute()
	}

	return <-chans[N]
}
