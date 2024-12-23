package main

import (
	"fmt"
	"sort"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	var graph = make(map[string]Set[string])
	for _, link := range InputToLinks() {
		graph[link[0]] = graph[link[0]].UnionElems(link[1])
		graph[link[1]] = graph[link[1]].UnionElems(link[0])
	}

	var cliques Set[[3]string]
	for c := range graph {
		cs := graph[c]
		if c[0] != 't' {
			continue
		}

		for d := range cs {
			ds := graph[d]
			for e := range ds {
				es := graph[e]

				if cs.Contains(d) && cs.Contains(e) &&
					ds.Contains(c) && ds.Contains(e) &&
					es.Contains(c) && es.Contains(d) {
					clique := []string{c, d, e}
					sort.Strings(clique)
					cliques.Add([3]string(clique))
				}
			}
		}
	}

	fmt.Println(len(cliques))
}

func InputToLinks() [][]string {
	return InputLinesTo(func(s string) []string {
		return strings.Split(s, "-")
	})
}
