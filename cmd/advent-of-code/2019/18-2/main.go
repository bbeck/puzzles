package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid, entrance := InputToGrid()
	grid.SetPoint(entrance, Wall)
	grid.SetPoint(entrance.Up(), Wall)
	grid.SetPoint(entrance.Down(), Wall)
	grid.SetPoint(entrance.Left(), Wall)
	grid.SetPoint(entrance.Right(), Wall)

	entrances := [4]Point2D{
		entrance.Up().Left(),
		entrance.Up().Right(),
		entrance.Down().Left(),
		entrance.Down().Right(),
	}

	graph := GridToGraph(grid, entrances)

	children := func(s State) []State {
		var children []State
		for i := range len(s.Locations) {
			for p := range graph[s.Locations[i]] {
				value := grid.GetPoint(p)

				// We need to have the key if this child location is a door
				if IsDoor(value) && !s.Keys.Contains(KeyIDForDoorID(value)) {
					continue
				}

				locations := s.Locations
				locations[i] = p
				child := State{Locations: locations, Keys: s.Keys}

				// If this location contains a key then we acquire it
				if IsKey(value) {
					child.Keys = child.Keys.Add(value)
				}

				children = append(children, child)
			}
		}

		return children
	}

	goal := func(s State) bool {
		return s.Keys.Size() == 26
	}

	cost := func(from, to State) int {
		var sum int
		for i := range len(from.Locations) {
			sum += graph[from.Locations[i]][to.Locations[i]]
		}
		return sum
	}

	heuristic := func(s State) int {
		return 26 - s.Keys.Size()
	}

	start := State{Locations: entrances}
	_, c, _ := AStarSearch(start, children, goal, cost, heuristic)
	fmt.Println(c)
}

type State struct {
	Locations [4]Point2D
	Keys      BitSet
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

func GridToGraph(grid Grid2D[int], entrances [4]Point2D) map[Point2D]map[Point2D]int {
	// Collect the key/door locations along with the entrance location.
	ps := entrances[:]
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

			if ok {
				for _, p := range path[1 : len(path)-2] {
					if grid.GetPoint(p) != Empty {
						ok = false
						break
					}
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
