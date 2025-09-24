package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	graph := make(map[string]string)
	for _, o := range InputToOrbits() {
		graph[o.Child] = o.Parent
	}

	var count int
	for n := range graph {
		for {
			n = graph[n]
			if n == "" {
				break
			}
			count++
		}
	}
	fmt.Println(count)
}

type Orbit struct {
	Parent, Child string
}

func InputToOrbits() []Orbit {
	return in.LinesToS(func(in in.Scanner[Orbit]) Orbit {
		parent, child := in.Cut(")")
		return Orbit{Parent: parent, Child: child}
	})
}
