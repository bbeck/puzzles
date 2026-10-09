package main

import (
	"fmt"
	"math"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	buttons, joltages := InputToButtonsAndJoltages()

	var sum int
	for i := range buttons {
		index := BuildParityIndex(buttons[i], len(joltages[i]))
		best, found := Solve(joltages[i], index, make(map[string]int))
		if found {
			sum += best
		}
	}
	fmt.Println(sum)
}

func Solve(joltages Joltages, index Index, memo map[string]int) (int, bool) {
	// The idea here is that we're going to divide and conquer.  Since we need to
	// make all joltages zero and a given button push can only ever impact a
	// joltage by 1, we can consider the parity of the joltages.  Any odd joltage
	// values will need to be made even to get closer to a solution.
	//
	// Once all joltage values are even, then we can halve their values and find
	// the solution to this now simpler problem.  The solution to the larger
	// problem is just doubling the simpler solution.

	// Base case 1, check if we've processed this case before.
	key := fmt.Sprintf("%v", joltages)
	if v, ok := memo[key]; ok {
		return v, true
	}

	// Base case 2, check if we've found a solution.
	var done = true
	for i := range joltages {
		if joltages[i] != 0 {
			done = false
			break
		}
	}
	if done {
		return 0, true
	}

	// Determine the mask of joltages that need to change parity
	var mask uint64
	for i, joltage := range joltages {
		if joltage%2 == 1 {
			mask |= 1 << (len(joltages) - i - 1)
		}
	}

	// Now consider each different way we have to reach the state where all
	// joltages are even.
	var best = math.MaxInt64
	var found bool

outer:
	for _, entry := range index[mask] {
		var subproblem = make(Joltages, len(joltages))
		for i := range joltages {
			// It's possible that using this solution would result in reducing our
			// joltages too much.  If this happens skip it and move on to the next
			// candidate.
			if entry.JoltageDelta[i] > joltages[i] {
				continue outer
			}

			subproblem[i] = (joltages[i] - entry.JoltageDelta[i]) / 2
		}

		sub, ok := Solve(subproblem, index, memo)
		if !ok {
			continue
		}

		best = Min(best, 2*sub+entry.NumButtons)
		found = true
	}

	if found {
		memo[key] = best
	}
	return best, found
}

type Entry struct {
	NumButtons   int
	JoltageDelta Joltages
}

type Index = map[uint64][]Entry

func BuildParityIndex(buttons Buttons, J int) Index {
	// Build an index of which buttons can be pushed to cause a specific parity
	// change to a joltage.  For example if we had these buttons available:
	//   b0: (2) | b1: (1,2) | b2: (0,1) | b3: (0,1,2)
	//
	// Then the resulting index would look like:
	//   Mask -> Buttons              -> Entries
	//   000  -> [(b0,b2,b3)]         -> [(NB:0, JD:[0,0,0]) (NB:3, JD:[2,2,2])]
	//   001  -> [(b0) (b2,b3)]       -> [(NB:1, JD:[0,0,1]) (NB:2, JD:[2,2,1])]
	//   010  -> [(b0,b1) (b1,b2,b3)] -> [(NB:2, JD:[0,1,2]) (NB:3, JD:[2,3,2])]
	//   011  -> [(b1) (b0,b1,b2,b3)] -> [(NB:1, JD:[0,1,1]) (NB:4, JD:[2,3,3])]
	//   100  -> [(b1,b3) (b0,b1,b2)] -> [(NB:2, JD:[1,2,2]) (NB:3, JD:[1,2,2])]
	//   101  -> [(b1,b2) (b0,b1,b3)] -> [(NB:2, JD:[1,2,1]) (NB:3, JD:[1,2,3])]
	//   110  -> [(b2) (b0,b3)]       -> [(NB:1, JD:[1,1,0]) (NB:2, JD:[1,1,2])]
	//   111  -> [(b3) (b0,b2)]       -> [(NB:1, JD:[1,1,1]) (NB:2, JD:[1,1,1])]
	var index = Index{
		// Make sure to include the option of pushing no buttons to make no change.
		uint64(0): {{NumButtons: 0, JoltageDelta: make([]int, J)}},
	}

	for k := 1; k <= len(buttons); k++ {
		EnumerateCombinations(len(buttons), k, func(indices []int) bool {
			var mask uint64
			var entry = Entry{
				NumButtons:   len(indices),
				JoltageDelta: make(Joltages, J),
			}
			for _, idx := range indices {
				for _, button := range buttons[idx] {
					mask ^= 1 << uint(J-button-1)
					entry.JoltageDelta[button]++
				}
			}

			index[mask] = append(index[mask], entry)
			return false
		})
	}

	return index
}

type Buttons = [][]int
type Joltages []int

func InputToButtonsAndJoltages() ([]Buttons, []Joltages) {
	in.Remove("[", "]", "(", ")", "{", "}")

	var allButtons []Buttons
	var allJoltages []Joltages
	for in.HasNext() {
		fields := in.FieldsS[any]()

		var buttons Buttons
		for i := 1; i < len(fields)-1; i++ {
			buttons = append(buttons, fields[i].Ints())
		}
		allButtons = append(allButtons, buttons)

		var joltages = fields[len(fields)-1].Ints()
		allJoltages = append(allJoltages, joltages)
	}

	return allButtons, allJoltages
}
