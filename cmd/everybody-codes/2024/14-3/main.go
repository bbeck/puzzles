package main

import (
	"fmt"
	"math"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var tree, leaves Set[Point3D]
	for in.HasNext() {
		var p Point3D
		for _, instruction := range strings.Split(in.Line(), ",") {
			var dx, dy, dz int
			switch instruction[0] {
			case 'U':
				dy = 1
			case 'D':
				dy = -1
			case 'L':
				dx = 1
			case 'R':
				dx = -1
			case 'F':
				dz = 1
			case 'B':
				dz = -1
			}

			N := ParseInt(instruction[1:])
			for n := 0; n < N; n++ {
				p = Point3D{X: p.X + dx, Y: p.Y + dy, Z: p.Z + dz}
				tree.Add(p)
			}
		}

		leaves.Add(p)
	}

	best := math.MaxInt

	trunk := Point3D{Y: 1}
	for tree.Contains(trunk) {
		distances := Distances(tree, trunk)

		var murkiness int
		for leaf := range leaves {
			murkiness += distances[leaf]
		}

		best = Min(best, murkiness)
		trunk.Y++
	}
	fmt.Println(best)
}

func Distances(tree Set[Point3D], start Point3D) map[Point3D]int {
	type State struct {
		P     Point3D
		Depth int
	}

	seen := make(map[Point3D]int)
	children := func(s State) []State {
		seen[s.P] = s.Depth

		var children []State
		for _, n := range s.P.OrthogonalNeighbors() {
			if _, found := seen[n]; found || !tree.Contains(n) {
				continue
			}

			children = append(children, State{P: n, Depth: s.Depth + 1})
		}
		return children
	}

	goal := func(State) bool { return false }
	BreadthFirstSearch(State{P: start}, children, goal)
	return seen
}
