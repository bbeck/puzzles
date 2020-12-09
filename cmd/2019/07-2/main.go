package main

import (
	"fmt"
	"sync"

	"github.com/bbeck/advent-of-code/aoc"
	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	settings := []int{5, 6, 7, 8, 9}

	var best int
	aoc.EnumeratePermutations(5, func(perm []int) {
		phases := []int{
			settings[perm[0]],
			settings[perm[1]],
			settings[perm[2]],
			settings[perm[3]],
			settings[perm[4]],
		}

		EtoA := make(chan int, 2)
		AtoB := make(chan int, 2)
		BtoC := make(chan int, 2)
		CtoD := make(chan int, 2)
		DtoE := make(chan int, 2)

		// Initialize the phase settings
		EtoA <- phases[0]
		AtoB <- phases[1]
		BtoC <- phases[2]
		CtoD <- phases[3]
		DtoE <- phases[4]

		var wg sync.WaitGroup
		wg.Add(5)

		EtoA <- 0 // First amplifier's first input is hardcoded to zero
		A := cpus.IntcodeCPU{
			Memory: cpus.InputToIntcodeMemory(2019, 7),
			Input:  func(int) int { return <-EtoA },
			Output: func(value int) { AtoB <- value },
		}
		go func() { A.Execute(); wg.Done() }()

		B := cpus.IntcodeCPU{
			Memory: cpus.InputToIntcodeMemory(2019, 7),
			Input:  func(int) int { return <-AtoB },
			Output: func(value int) { BtoC <- value },
		}
		go func() { B.Execute(); wg.Done() }()

		C := cpus.IntcodeCPU{
			Memory: cpus.InputToIntcodeMemory(2019, 7),
			Input:  func(int) int { return <-BtoC },
			Output: func(value int) { CtoD <- value },
		}
		go func() { C.Execute(); wg.Done() }()

		D := cpus.IntcodeCPU{
			Memory: cpus.InputToIntcodeMemory(2019, 7),
			Input:  func(int) int { return <-CtoD },
			Output: func(value int) { DtoE <- value },
		}
		go func() { D.Execute(); wg.Done() }()

		E := cpus.IntcodeCPU{
			Memory: cpus.InputToIntcodeMemory(2019, 7),
			Input:  func(int) int { return <-DtoE },
			Output: func(value int) { EtoA <- value },
		}
		go func() { E.Execute(); wg.Done() }()

		wg.Wait()

		signal := <-EtoA
		if signal > best {
			best = signal
		}
	})

	fmt.Printf("largest signal: %d\n", best)
}
