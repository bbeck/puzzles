package main

import (
	"fmt"
	"slices"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var nids, nquests, results = InputToResults()

	var tms []int
	for id := range nids {
		for quest := range nquests {
			tms = append(tms, Sum(results[id][quest]...))
		}
	}
	slices.Sort(tms)

	var sum int
	for i, tm := range tms {
		sum += (i + 1) * tm
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
