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

	var points = Make3D[int](nids, nquests, 3)
	for quest := range nquests {
		for part := range 3 {
			var times = make([]int, nids)
			for id := range nids {
				if results[id][quest][part] != 0 {
					times[id] = Sum(results[id][quest][:part+1]...)
				}
			}

			var unique = slices.DeleteFunc(
				slices.Clone(times), // must clone because DeleteFunc modifies in place
				func(i int) bool { return i == 0 },
			)
			slices.Sort(unique)

			for id := range nids {
				if results[id][quest][part] != 0 {
					idx, _ := slices.BinarySearch(unique, times[id])
					points[id][quest][part] = Max(0, 50*(part+1)-idx)
				}
			}
		}
	}

	var totals = make([]int, nids)
	for id := range nids {
		for quest := range nquests {
			totals[id] += Sum(points[id][quest]...)
		}
	}
	slices.SortFunc(totals, func(a, b int) int { return cmp.Compare(b, a) })

	var sum int
	for idx, total := range totals {
		for idx > 0 && totals[idx-1] == total {
			idx--
		}
		sum += (idx + 1) * total
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
