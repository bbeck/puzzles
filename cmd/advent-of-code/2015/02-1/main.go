package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var area int
	for in.HasNext() {
		L, W, H := in.Int(), in.Int(), in.Int()
		area += (2*L*W + 2*L*H + 2*W*H) + Min(L*W, L*H, W*H)
	}

	fmt.Println(area)
}
