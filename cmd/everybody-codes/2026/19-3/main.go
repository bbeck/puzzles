package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var counts FrequencyCounter[rune]
	for in.HasNext() {
		for _, c := range in.Line() {
			counts.Add(c)
		}
	}

	// There's two worlds to consider here.  There are S = nC + nM + nY non-black
	// cells in the grid.
	S := counts.GetCount('C') + counts.GetCount('M') + counts.GetCount('Y')

	// World #1 is when the 3 largest output groups don't include the K.  In this
	// case we know a+b+c=S, and we're trying to maximize a*b*c.
	a, b, c := (S+2)/3, (S+1)/3, S/3
	choice1 := a * b * c

	// World #2 is when the 3 largest output groups does include the K.  In this
	// case we know that a+b=S, and we're trying to maximize a*b*nK.
	a, b = (S+1)/2, S/2
	choice2 := a * b * counts.GetCount('K')

	fmt.Println(Max(choice1, choice2))
}
