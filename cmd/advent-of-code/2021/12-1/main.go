package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	caves := InputToCaves()
	count := CountPaths("start", "end", caves, lib.SetFrom("start"))
	fmt.Println(count)
}

func CountPaths(current, goal string, caves map[string][]string, seen lib.Set[string]) int {
	if current == goal {
		return 1
	}

	var count int
	for _, n := range caves[current] {
		if n == strings.ToLower(n) && seen.Contains(n) {
			continue
		}

		count += CountPaths(n, goal, caves, seen.UnionElems(n))
	}

	return count
}

func InputToCaves() map[string][]string {
	caves := make(map[string][]string)
	for _, line := range lib.InputToLines() {
		lhs, rhs, _ := strings.Cut(line, "-")
		caves[lhs] = append(caves[lhs], rhs)
		caves[rhs] = append(caves[rhs], lhs)
	}
	return caves
}
