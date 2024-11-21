package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	grid, entrance := InputToGrid()
	grid.SetPoint(entrance, Wall)
	grid.SetPoint(entrance.Up(), Wall)
	grid.SetPoint(entrance.Down(), Wall)
	grid.SetPoint(entrance.Left(), Wall)
	grid.SetPoint(entrance.Right(), Wall)

	entrances := [4]lib.Point2D{
		entrance.Up().Left(),
		entrance.Up().Right(),
		entrance.Down().Left(),
		entrance.Down().Right(),
	}

	graph := GridToGraph(grid, entrances)

	children := func(s State) []State {
		var children []State
		for i := 0; i < len(s.Locations); i++ {
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
		for i := 0; i < len(from.Locations); i++ {
			sum += graph[from.Locations[i]][to.Locations[i]]
		}
		return sum
	}

	heuristic := func(s State) int {
		return 26 - s.Keys.Size()
	}

	start := State{Locations: entrances}
	_, c, _ := lib.AStarSearch(start, children, goal, cost, heuristic)
	fmt.Println(c)
}

type State struct {
	Locations [4]lib.Point2D
	Keys      lib.BitSet
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

func GridToGraph(grid lib.Grid2D[int], entrances [4]lib.Point2D) map[lib.Point2D]map[lib.Point2D]int {
	// Collect the key/door locations along with the entrance location.
	ps := entrances[:]
	grid.ForEachPoint(func(p lib.Point2D, value int) {
		if value != Empty && value != Wall {
			ps = append(ps, p)
		}
	})

	// Determine the distance between all pairs of points of interest.  Ignore
	// any paths that include a 3rd point of interest.
	graph := make(map[lib.Point2D]map[lib.Point2D]int)
	for _, p := range ps {
		graph[p] = make(map[lib.Point2D]int)
	}

	children := func(p lib.Point2D) []lib.Point2D {
		var children []lib.Point2D
		grid.ForEachOrthogonalNeighborPoint(p, func(child lib.Point2D, value int) {
			if value != Wall {
				children = append(children, child)
			}
		})
		return children
	}

	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			path, ok := lib.BreadthFirstSearch(ps[i], children, func(p lib.Point2D) bool { return p == ps[j] })

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

func InputToGrid() (lib.Grid2D[int], lib.Point2D) {
	var entrance lib.Point2D
	return lib.InputToGrid2D(func(x int, y int, s string) int {
		if s == "@" {
			entrance = lib.Point2D{X: x, Y: y}
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
