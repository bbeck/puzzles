package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	phases := []int{0, 1, 2, 3, 4}

	var best int
	aoc.EnumeratePermutations(5, func(perm []int) {
		ZtoA := make(chan int, 2)
		AtoB := make(chan int, 2)
		BtoC := make(chan int, 2)
		CtoD := make(chan int, 2)
		DtoE := make(chan int, 2)
		EtoT := make(chan int)

		// Initialize the phase settings
		ZtoA <- phases[perm[0]]
		AtoB <- phases[perm[1]]
		BtoC <- phases[perm[2]]
		CtoD <- phases[perm[3]]
		DtoE <- phases[perm[4]]

		ZtoA <- 0 // First amplifier's input is hardcoded to zero
		A := cpus.IntcodeCPU{
			Memory: cpus.InputToIntcodeMemory(2019, 7),
			Input:  func(int) int { return <-ZtoA },
			Output: func(value int) { AtoB <- value },
		}
		go A.Execute()

		B := cpus.IntcodeCPU{
			Memory: cpus.InputToIntcodeMemory(2019, 7),
			Input:  func(int) int { return <-AtoB },
			Output: func(value int) { BtoC <- value },
		}
		go B.Execute()

		C := cpus.IntcodeCPU{
			Memory: cpus.InputToIntcodeMemory(2019, 7),
			Input:  func(int) int { return <-BtoC },
			Output: func(value int) { CtoD <- value },
		}
		go C.Execute()

		D := cpus.IntcodeCPU{
			Memory: cpus.InputToIntcodeMemory(2019, 7),
			Input:  func(int) int { return <-CtoD },
			Output: func(value int) { DtoE <- value },
		}
		go D.Execute()

		E := cpus.IntcodeCPU{
			Memory: cpus.InputToIntcodeMemory(2019, 7),
			Input:  func(int) int { return <-DtoE },
			Output: func(value int) { EtoT <- value },
		}
		go E.Execute()

		signal := <-EtoT
		if signal > best {
			best = signal
		}
	})

	fmt.Printf("largest signal: %d\n", best)
}
