package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	energies, connections := InputToClouds()
	N := len(energies) + 1

	var start int
	for id := range N {
		if len(connections[id]) > len(connections[start]) {
			start = id
		}
	}

	var counts = make([][11]int, N)
	counts[start][10] = 1

	var received = make(map[int]int)
	for active := true; active; {
		active = false

		var next = make([][11]int, N)
		for id := range N {
			for amt := 1; amt <= 10; amt++ {
				count := counts[id][amt]
				received[id] += count * amt

				if count > 0 && amt > 1 {
					for _, other := range connections[id] {
						next[other][amt-1] += count
						active = true
					}
				}
			}
		}

		counts = next
	}

	var total int
	for id := range N {
		total += energies[id] + received[id]
	}
	fmt.Println(total)
}

func InputToClouds() (map[int]int, map[int][]int) {
	var energies, connections = make(map[int]int), make(map[int][]int)
	for in.HasNext() {
		in := in.Scanner[any](in.Line())

		id := in.Int()
		energies[id] = in.Int()
		connections[id] = in.Ints()
	}

	return energies, connections
}
