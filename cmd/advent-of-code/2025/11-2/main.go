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

	c1 := Count("svr", "fft", nodes) *
		Count("fft", "dac", nodes) *
		Count("dac", "out", nodes)
	c2 := Count("svr", "dac", nodes) *
		Count("dac", "fft", nodes) *
		Count("fft", "out", nodes)
	fmt.Println(c1 + c2)
}

var memo = make(map[string]int)

func Count(current, goal string, nodes map[string][]string) int {
	if current == goal {
		return 1
	}

	key := fmt.Sprintf("%s->%s", current, goal)
	if v, ok := memo[key]; ok {
		return v
	}

	var count int
	for _, next := range nodes[current] {
		count += Count(next, goal, nodes)
	}

	memo[key] = count
	return count
}
