package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	plants := InputToPlants()

	var sum int
	for _, plant := range plants {
		sum += Grow(plant)
	}
	fmt.Println(sum)
}

func Grow(p Plant) int {
	var branches = []int{0}
	var forks int

	for tm := 1; ; tm++ {
		n := len(branches) // Snapshot length so we don't see appends

		for i, b := range branches[:n] {
			if b+1 == p.Growth {
				// Splitting yields b/2, b/2 and 1
				forks++
				branches[i] = 1
				branches = append(branches, b/2, b/2)
			} else {
				branches[i]++
			}
		}

		if Sum(branches...)+forks >= p.Limit {
			return tm
		}
	}
}

type Plant struct {
	ID, Growth, Limit int
}

func InputToPlants() []Plant {
	return in.LinesToS[Plant](func(in in.Scanner[Plant]) Plant {
		return Plant{ID: in.Int(), Growth: in.Int(), Limit: in.Int()}
	})
}
