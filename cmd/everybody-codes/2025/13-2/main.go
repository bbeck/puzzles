package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var ring = InputToRing()
	fmt.Println(ring[20252025%len(ring)])
}

func InputToRing() []int {
	next := func() []int {
		var ns []int
		var start, end int
		in.Scanf("%d-%d", &start, &end)
		for i := start; i <= end; i++ {
			ns = append(ns, i)
		}
		return ns
	}

	var front, back []int
	for in.HasNext() {
		if in.HasNext() {
			front = append(front, next()...)
		}
		if in.HasNext() {
			back = append(back, next()...)
		}
	}

	return append([]int{1}, append(front, Reversed(back)...)...)
}
