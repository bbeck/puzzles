package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/cpus"
	"sync"
)

const N = 5

func main() {
	var best int
	lib.EnumeratePermutations(N, func(perm []int) bool {
		settings := make([]int, N)
		for i := 0; i < len(perm); i++ {
			settings[i] = perm[i] + 5
		}

		best = lib.Max(best, TestSettings(settings))
		return false
	})
	fmt.Println(best)
}

func TestSettings(settings []int) int {
	var chans [N]chan int
	for i := 0; i < len(chans); i++ {
		chans[i] = make(chan int, 2)
	}

	// Send the settings into the inputs
	for i, setting := range settings {
		chans[i] <- setting
	}

	// First amplifier's input is hardcoded to 0
	chans[0] <- 0

	//
	var wg sync.WaitGroup
	wg.Add(N)

	var amps [N]cpus.IntcodeCPU
	for i := 0; i < N; i++ {
		i := i
		amps[i].Memory = cpus.InputToIntcodeMemory()
		amps[i].Input = func() int { return <-chans[i] }
		amps[i].Output = func(value int) { chans[(i+1)%N] <- value }
		amps[i].Halt = func() { wg.Done() }
		go amps[i].Execute()
	}

	wg.Wait()
	return <-chans[0]
}
