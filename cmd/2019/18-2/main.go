package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid, entrance := InputToGrid()
	grid.Add(entrance, Wall)
	grid.Add(entrance.Up(), Wall)
	grid.Add(entrance.Down(), Wall)
	grid.Add(entrance.Left(), Wall)
	grid.Add(entrance.Right(), Wall)

	entrances := [4]aoc.Point2D{
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
				value := grid.Get(p)

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
	_, c, _ := aoc.AStarSearch(start, children, goal, cost, heuristic)
	fmt.Println(c)
}

type State struct {
	Locations [4]aoc.Point2D
	Keys      aoc.BitSet
}

const (
	Empty = -1
	Wall  = -2
)

func IsKey(v int) bool            { return 0 <= v && v < 26 }
func IsDoor(v int) bool           { return 26 <= v && v < 52 }
func KeyID(c rune) int            { return int(c - 'a') }
func DoorID(c rune) int           { return int(c - 'A' + 26) }
func KeyIDForDoorID(door int) int { return door - 26 }

func GridToGraph(grid aoc.Grid2D[int], entrances [4]aoc.Point2D) map[aoc.Point2D]map[aoc.Point2D]int {
	// Collect the key/door locations along with the entrance location.
	ps := entrances[:]
	grid.ForEach(func(p aoc.Point2D, value int) {
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
		for _, child := range p.OrthogonalNeighbors() {
			if grid.Get(child) != Wall {
				children = append(children, child)
			}
		}
		return children
	}

	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			path, ok := aoc.BreadthFirstSearch(ps[i], children, func(p aoc.Point2D) bool { return p == ps[j] })

			if ok {
				for _, p := range path[1 : len(path)-2] {
					if grid.Get(p) != Empty {
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

func InputToGrid() (aoc.Grid2D[int], aoc.Point2D) {
	lines := aoc.InputToLines(2019, 18)

	grid := aoc.NewGrid2D[int](len(lines[0]), len(lines))
	entrance := aoc.Origin2D
	for y, line := range lines {
		for x, c := range line {
			value := Empty
			if c == '@' {
				entrance = aoc.Point2D{X: x, Y: y}
			} else if c == '#' {
				value = Wall
			} else if c >= 'a' && c <= 'z' {
				value = KeyID(c)
			} else if c >= 'A' && c <= 'Z' {
				value = DoorID(c)
			}
			grid.AddXY(x, y, value)
		}
	}

	return grid, entrance
}
