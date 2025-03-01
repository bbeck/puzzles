package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var length int
	for in.HasNext() {
		L, W, H := in.Int(), in.Int(), in.Int()
		length += 2*(Sum(L, W, H)-Max(L, W, H)) + (L * W * H)
	}
	fmt.Println(length)
}
