package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	in.Remove(":")

	var nodes = make(map[string][]string)
	for in.HasNext() {
		nodes[in.String()] = in.Fields()
	}

	fmt.Println(Count("you", "out", nodes))
}

func Count(n, end string, nodes map[string][]string) int {
	if n == end {
		return 1
	}

	var count int
	for _, s := range nodes[n] {
		count += Count(s, end, nodes)
	}
	return count
}
