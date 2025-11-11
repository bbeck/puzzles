package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

var UPPER = map[byte]byte{'a': 'A', 'b': 'B', 'c': 'C'}

func main() {
	var sum int
	var mentors = make(map[byte]int)
	for in.HasNext() {
		switch ch := in.Byte(); ch {
		case 'A', 'B', 'C':
			mentors[ch]++
		case 'a', 'b', 'c':
			sum += mentors[UPPER[ch]]
		}
	}
	fmt.Println(sum)
}
