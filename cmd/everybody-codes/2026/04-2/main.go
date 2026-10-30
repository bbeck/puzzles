package main

import (
	"cmp"
	"fmt"
	"slices"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var nids, nquests, results = InputToResults()

	var ids, times, parts = make([]int, nids), make([]int, nids), make([]int, nids)
	for id := range nids {
		ids[id] = id
		for quest := range nquests {
			for part := range 3 {
				if tm := results[id][quest][part]; tm != 0 {
					times[id] += tm
					parts[id]++
				}
			}
		}
	}
	slices.SortFunc(ids, func(a, b int) int {
		return cmp.Or(
			cmp.Compare(parts[b], parts[a]),
			cmp.Compare(times[a], times[b]),
		)
	})

	var sum int
	for idx, id := range ids {
		sum += (idx + 1) * times[id]
	}
	fmt.Println(sum)
}

func InputToResults() (int, int, [][][]int) {
	var ids, quests, parts, times []int
	for in.HasNextLine() {
		ids = append(ids, in.Int()-1)
		quests = append(quests, in.Int()-1)
		parts = append(parts, in.Int()-1)
		times = append(times, in.Int())
	}

	var numIds, numQuests = len(SetFrom(ids...)), len(SetFrom(quests...))
	var results = Make3D[int](numIds, numQuests, 3)
	for i := range ids {
		results[ids[i]][quests[i]][parts[i]] = times[i]
	}
	return numIds, numQuests, results
}
