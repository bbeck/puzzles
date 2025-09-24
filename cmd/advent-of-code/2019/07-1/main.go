package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/cpus"
)

const N = 5

func main() {
	var memory = cpus.InputToIntcodeMemory()

	var best int
	EnumeratePermutations(N, func(settings []int) bool {
		best = Max(best, TestSettings(memory.Copy(), settings))
		return false
	})
	fmt.Println(best)
}

func TestSettings(memory cpus.Memory, settings []int) int {
	var chans [N + 1]chan int
	for i := range len(chans) {
		chans[i] = make(chan int, 2)
	}

	// Send the settings into the inputs
	for i, setting := range settings {
		chans[i] <- setting
	}

	// First amplifier's input is hardcoded to 0
	chans[0] <- 0

	var amps [N]cpus.IntcodeCPU
	for i := range N {
		i := i
		amps[i].Memory = memory.Copy()
		amps[i].Input = func() int { return <-chans[i] }
		amps[i].Output = func(value int) { chans[i+1] <- value }
		go amps[i].Execute()
	}

	return <-chans[N]
}
