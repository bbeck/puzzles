package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	rules, updates := InputToRulesAndUpdates()

	var sum int
	for _, update := range updates {
		if IsCorrect(rules, update) {
			continue
		}

		// The complete rule graph has cycles in it involving nodes that aren't
		// used together in updates.  So build a separate graph for each update
		// involving only the nodes that are used in that update.
		used := SetFrom(update...)

		var graph Graph[string]
		for parent, children := range rules {
			for _, child := range children {
				if used.Contains(parent) && used.Contains(child) {
					graph.AddEdge(parent, child)
				}
			}
		}

		ts, _ := graph.TopologicalSort()
		sum += ParseInt(ts[len(ts)/2])
	}
	fmt.Println(sum)
}

func IsCorrect(rules map[string][]string, update []string) bool {
	var seen Set[string]
	for _, parent := range update {
		for _, child := range rules[parent] {
			if seen.Contains(child) {
				return false
			}
		}
		seen.Add(parent)
	}
	return true
}

func InputToRulesAndUpdates() (map[string][]string, [][]string) {
	rules := make(map[string][]string)
	updates := make([][]string, 0)

	for _, line := range InputToLines() {
		switch {
		case strings.Contains(line, "|"):
			lhs, rhs, _ := strings.Cut(line, "|")
			rules[lhs] = append(rules[lhs], rhs)
		case strings.Contains(line, ","):
			updates = append(updates, strings.Split(line, ","))
		}
	}

	return rules, updates
}

type Graph[N comparable] struct {
	Children map[N][]N
}

func (g *Graph[N]) AddEdge(child, parent N) {
	if g.Children == nil {
		g.Children = make(map[N][]N)
	}

	g.Children[parent] = append(g.Children[parent], child)
}

func (g *Graph[N]) TopologicalSort() ([]N, error) {
	var ts []N
	mark := make(map[N]string)

	var visit func(N) error
	visit = func(n N) error {
		if mark[n] == "permanent" {
			return nil
		}
		if mark[n] == "temporary" {
			return fmt.Errorf("graph contains cycle involving: %v", n)
		}

		mark[n] = "temporary"
		for _, m := range g.Children[n] {
			if err := visit(m); err != nil {
				return err
			}
		}
		mark[n] = "permanent"
		ts = append([]N{n}, ts...)
		return nil
	}

	for n := range g.Children {
		if mark[n] == "" {
			if err := visit(n); err != nil {
				return nil, err
			}
		}
	}

	return ts, nil
}
