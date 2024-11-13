package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	// This is asking us to find a tm that satisfies a system of congruences:
	//   0 = tm + offsets[0] (mod ids[0])  ->  tm = -offsets[0] (mod ids[0])
	//   0 = tm + offsets[1] (mod ids[1])  ->  tm = -offsets[1] (mod ids[1])
	//   ...
	//
	// To do this we can use the chinese remainder theorem.
	ids, offsets := InputToBuses()
	for i := 0; i < len(offsets); i++ {
		offsets[i] *= -1
	}
	fmt.Println(puz.ChineseRemainderTheorem(offsets, ids))
}

func InputToBuses() ([]int, []int) {
	lines := puz.InputToLines(2020, 13)

	var ids, offsets []int
	for i, s := range strings.Split(lines[1], ",") {
		if s != "x" {
			ids = append(ids, puz.ParseInt(s))
			offsets = append(offsets, i)
		}
	}

	return ids, offsets
}
