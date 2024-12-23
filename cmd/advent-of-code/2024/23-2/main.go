package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"sort"
	"strings"
)

func main() {
	var graph = make(map[string]Set[string])
	for _, link := range InputToLinks() {
		graph[link[0]] = graph[link[0]].UnionElems(link[1])
		graph[link[1]] = graph[link[1]].UnionElems(link[0])
	}

	var clique []string
	EnumerateMaximalCliques(graph, func(c []string) {
		if len(c) > len(clique) {
			clique = c
		}
	})

	sort.Strings(clique)
	fmt.Println(strings.Join(clique, ","))
}

func InputToLinks() [][]string {
	return InputLinesTo(func(s string) []string {
		return strings.Split(s, "-")
	})
}
