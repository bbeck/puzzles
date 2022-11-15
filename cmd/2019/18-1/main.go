package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid, entrance := InputToGrid()
	graph := GridToGraph(grid, entrance)

	children := func(s State) []State {
		var children []State
		for p := range graph[s.Point2D] {
			value := grid.GetPoint(p)

			// We need to have the key if this child location is a door
			if IsDoor(value) && !s.Keys.Contains(KeyIDForDoorID(value)) {
				continue
			}

			child := State{Point2D: p, Keys: s.Keys}

			// If this location contains a key then we acquire it
			if IsKey(value) {
				child.Keys = child.Keys.Add(value)
			}

			children = append(children, child)
		}
		return children
	}

	goal := func(s State) bool {
		return s.Keys.Size() == 26
	}

	cost := func(from, to State) int {
		return graph[from.Point2D][to.Point2D]
	}

	heuristic := func(s State) int {
		return 26 - s.Keys.Size()
	}

	start := State{Point2D: entrance}
	_, c, _ := aoc.AStarSearch(start, children, goal, cost, heuristic)
	fmt.Println(c)
}

type State struct {
	aoc.Point2D
	Keys aoc.BitSet
}

const (
	Empty = -1
	Wall  = -2
)

func IsKey(v int) bool            { return 0 <= v && v < 26 }
func IsDoor(v int) bool           { return 26 <= v && v < 52 }
func KeyID(s string) int          { return int(s[0] - 'a') }
func DoorID(s string) int         { return int(s[0] - 'A' + 26) }
func KeyIDForDoorID(door int) int { return door - 26 }

func GridToGraph(grid aoc.Grid2D[int], entrance aoc.Point2D) map[aoc.Point2D]map[aoc.Point2D]int {
	// Collect the key/door locations along with the entrance location.
	ps := []aoc.Point2D{entrance}
	grid.ForEachPoint(func(p aoc.Point2D, value int) {
		if value != Empty && value != Wall {
			ps = append(ps, p)
		}
	})

	// Determine the distance between all pairs of points of interest.  Ignore
	// any paths that include a 3rd point of interest.
	graph := make(map[aoc.Point2D]map[aoc.Point2D]int)
	for _, p := range ps {
		graph[p] = make(map[aoc.Point2D]int)
	}

	children := func(p aoc.Point2D) []aoc.Point2D {
		var children []aoc.Point2D
		grid.ForEachOrthogonalNeighbor(p, func(child aoc.Point2D, value int) {
			if value != Wall {
				children = append(children, child)
			}
		})
		return children
	}

	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			path, ok := aoc.BreadthFirstSearch(ps[i], children, func(p aoc.Point2D) bool { return p == ps[j] })
			for _, p := range path[1 : len(path)-2] {
				if grid.GetPoint(p) != Empty {
					ok = false
					break
				}
			}

			if ok {
				graph[ps[i]][ps[j]] = len(path) - 1
				graph[ps[j]][ps[i]] = len(path) - 1
			}
		}
	}

	return graph
}

func InputToGrid() (aoc.Grid2D[int], aoc.Point2D) {
	var entrance aoc.Point2D
	return aoc.InputToGrid2D(2019, 18, func(x int, y int, s string) int {
		if s == "@" {
			entrance = aoc.Point2D{X: x, Y: y}
		} else if s == "#" {
			return Wall
		} else if s >= "a" && s <= "z" {
			return KeyID(s)
		} else if s >= "A" && s <= "Z" {
			return DoorID(s)
		}

		return Empty
	}), entrance
}
