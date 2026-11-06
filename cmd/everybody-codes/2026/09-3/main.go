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
	if p.Limit == 1 {
		return 1
	}

	// The idea here is that each branch keeps growing until it reaches the growth
	// time for it.  Then it will split and cause more leaves.  Only at these
	// split times do the number of leaves change.  So instead of stepping over
	// time one second at a time, we'll keep track of the times when splits happen
	// in a heap and fast-forward.

	// Number of branches splitting at a particular time
	var counts = map[int]int{p.Growth: 1}

	// Priority queue storing the times of the next split
	var pq PriorityQueue[int]
	pq.Push(p.Growth, p.Growth)

	var leaves = 1
	for {
		tm := pq.Pop()
		n := counts[tm]
		delete(counts, tm)

		// Each split adds a branch that has a leaf.
		leaves += n
		if leaves >= p.Limit {
			return tm
		}

		// The split results in two halves of the original long branch that will
		// continue to grow, as well as a new size 1 branch.  The longer branches
		// are already half grown, whereas the size 1 branch has a lot of growing
		// to do.
		tmHalf, tmFull := tm+p.Growth/2+1, tm+p.Growth-1
		if _, ok := counts[tmHalf]; !ok {
			pq.Push(tmHalf, tmHalf)
		}
		if _, ok := counts[tmFull]; !ok {
			pq.Push(tmFull, tmFull)
		}
		counts[tmHalf] += 2 * n
		counts[tmFull] += n
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
