package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
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
	_, c, _ := AStarSearch(start, children, goal, cost, heuristic)
	fmt.Println(c)
}

type State struct {
	Point2D
	Keys BitSet
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

func GridToGraph(grid Grid2D[int], entrance Point2D) map[Point2D]map[Point2D]int {
	// Collect the key/door locations along with the entrance location.
	ps := []Point2D{entrance}
	grid.ForEachPoint(func(p Point2D, value int) {
		if value != Empty && value != Wall {
			ps = append(ps, p)
		}
	})

	// Determine the distance between all pairs of points of interest.  Ignore
	// any paths that include a 3rd point of interest.
	graph := make(map[Point2D]map[Point2D]int)
	for _, p := range ps {
		graph[p] = make(map[Point2D]int)
	}

	children := func(p Point2D) []Point2D {
		var children []Point2D
		grid.ForEachOrthogonalNeighborPoint(p, func(child Point2D, value int) {
			if value != Wall {
				children = append(children, child)
			}
		})
		return children
	}

	for i := range ps {
		for j := i + 1; j < len(ps); j++ {
			path, ok := BreadthFirstSearch(ps[i], children, func(p Point2D) bool { return p == ps[j] })
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

func InputToGrid() (Grid2D[int], Point2D) {
	var entrance Point2D
	return in.ToGrid2D[int](func(x, y int, s string) int {
		if s == "@" {
			entrance = Point2D{X: x, Y: y}
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
