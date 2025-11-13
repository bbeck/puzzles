package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

const N = 32

func main() {
	nails := in.Ints()

	var count int
	for i := 0; i < len(nails)-1; i++ {
		n1, n2 := Min(nails[i:i+2]...), Max(nails[i:i+2]...)
		if n2-n1 == N/2 {
			count++
		}
	}
	fmt.Println(count)
}
