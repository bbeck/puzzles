package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

const LIMIT = 1000000

func main() {
	energies, connections := InputToClouds()
	N := len(energies) + 1

	do := func(start int) int {
		var counts = make([][11]int, N)
		counts[start][10] = 1

		var exploded Set[int]
		var received = make(map[int]int)

		for active := true; active; {
			active = false

			var next = make([][11]int, N)
			for id := range N {
				if exploded.Contains(id) {
					continue
				}

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

			for id := range N {
				if exploded.Contains(id) {
					continue
				}

				if energies[id]+received[id] >= LIMIT {
					exploded.Add(id)
					for _, other := range connections[id] {
						next[other][10]++
						active = true
					}
				}
			}

			counts = next
		}

		var total int
		for id := range N {
			if !exploded.Contains(id) {
				total += energies[id] + received[id]
			}
		}
		return total
	}

	var best int
	for id := 1; id < N; id++ {
		best = Max(best, do(id))
	}
	fmt.Println(best)
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
