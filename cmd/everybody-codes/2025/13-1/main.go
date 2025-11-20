package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var ring = InputToRing()
	fmt.Println(ring[2025%len(ring)])
}

func InputToRing() []int {
	var front, back []int
	for in.HasNext() {
		if in.HasNext() {
			front = append(front, in.Int())
		}
		if in.HasNext() {
			back = append(back, in.Int())
		}
	}

	return append([]int{1}, append(front, Reversed(back)...)...)
}
