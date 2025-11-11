package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

var UPPER = map[byte]byte{'a': 'A', 'b': 'B', 'c': 'C'}

const REPEAT = 1000
const DIST = 1000

func main() {
	var s = in.String()
	get := func(i int) byte {
		if i < 0 || i >= REPEAT*len(s) {
			return 0
		}
		return s[i%len(s)]
	}

	var sum int
	var window = make(map[byte]int)
	for idx := -DIST; idx < len(s)*REPEAT; idx++ {
		window[get(idx-DIST-1)]--
		window[get(idx+DIST)]++

		switch ch := get(idx); ch {
		case 'a', 'b', 'c':
			sum += window[UPPER[ch]]
		}
	}
	fmt.Println(sum)
}
