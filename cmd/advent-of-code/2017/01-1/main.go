package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	s := in.Line()
	N := len(s)

	var sum int
	for i := range N {
		j := (i + 1) % N
		if s[i] == s[j] {
			sum += ParseInt(string(s[i]))
		}
	}
	fmt.Println(sum)
}
