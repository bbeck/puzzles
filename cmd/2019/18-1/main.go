package main

import (
	"fmt"
	"log"
	"math/bits"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

type Grid map[aoc.Point2D]bool
type Keys map[aoc.Point2D]string
type Doors map[aoc.Point2D]string
type Graph map[string]map[string]int

var grid Grid
var keys Keys
var doors Doors
var start aoc.Point2D
var graph Graph

var KeyMasks = map[string]uint32{
	"a": 0b10000000000000000000000000,
	"b": 0b01000000000000000000000000,
	"c": 0b00100000000000000000000000,
	"d": 0b00010000000000000000000000,
	"e": 0b00001000000000000000000000,
	"f": 0b00000100000000000000000000,
	"g": 0b00000010000000000000000000,
	"h": 0b00000001000000000000000000,
	"i": 0b00000000100000000000000000,
	"j": 0b00000000010000000000000000,
	"k": 0b00000000001000000000000000,
	"l": 0b00000000000100000000000000,
	"m": 0b00000000000010000000000000,
	"n": 0b00000000000001000000000000,
	"o": 0b00000000000000100000000000,
	"p": 0b00000000000000010000000000,
	"q": 0b00000000000000001000000000,
	"r": 0b00000000000000000100000000,
	"s": 0b00000000000000000010000000,
	"t": 0b00000000000000000001000000,
	"u": 0b00000000000000000000100000,
	"v": 0b00000000000000000000010000,
	"w": 0b00000000000000000000001000,
	"x": 0b00000000000000000000000100,
	"y": 0b00000000000000000000000010,
	"z": 0b00000000000000000000000001,
}

func main() {
	grid, keys, doors, start = InputToGrid(2019, 18)

	// We'll convert the grid into a graph where the nodes represented are the
	// starting point as well as every grid cell that contains a door or a key.
	// The edges will only be the paths in the grid that can be traversed
	// without passing over a key or through a door (unlocked or not).  The
	// edge weights will be the shortest distance between the two nodes.
	GridToGraph()

	// Now that we have a graph, we can use A* to determine the path through
	// the graph that gets all of the keys.

	isGoal := func(node aoc.Node) bool {
		state := node.(State)
		return state.keys == uint32(0b11111111111111111111111111)
	}

	cost := func(fromNode, toNode aoc.Node) int {
		from := fromNode.(State).location
		to := toNode.(State).location
		return graph[from][to]
	}

	heuristic := func(node aoc.Node) int {
		state := node.(State)
		return 26 - bits.OnesCount64(uint64(state.keys))
	}

	// Compute a mask of the keys that aren't used in the input, and we'll
	// pretend that we've already found those keys once we start.  This lets
	// us have a simpler isGoal function.
	mask := uint32(0b11111111111111111111111111)
	for _, key := range keys {
		mask &= ^KeyMasks[key]
	}

	state := State{
		keys:     mask,
		location: "@",
	}

	_, distance, found := aoc.AStarSearch(state, isGoal, cost, heuristic)
	if !found {
		log.Fatal("unable to find path")
	}

	fmt.Printf("distance: %d\n", distance)
}

type State struct {
	// Which keys we've collected
	keys uint32

	// Where we're located
	location string
}

func (s State) ID() string {
	return fmt.Sprintf("%s %d", s.location, s.keys)
}

func (s State) Children() []aoc.Node {
	var children []aoc.Node
	for name := range graph[s.location] {
		if name != strings.ToLower(name) {
			// This is a door, make sure we have the necessary key
			needed := strings.ToLower(name)
			if s.keys&KeyMasks[needed] == 0 {
				// We don't have the key
				continue
			}
		}

		child := State{
			keys:     s.keys | KeyMasks[name],
			location: name,
		}
		children = append(children, child)
	}

	return children
}

func GridToGraph() {
	graph = make(Graph)

	vertices := map[string]aoc.Point2D{
		"@": start,
	}
	for p, key := range keys {
		vertices[key] = p
	}
	for p, door := range doors {
		vertices[door] = p
	}

	for pname, p := range vertices {
		for qname, q := range vertices {
			if p == q {
				continue
			}

			distance := PathDistance(p, q)
			if distance == -1 {
				continue
			}

			m, ok := graph[pname]
			if !ok {
				m = make(map[string]int)
				graph[pname] = m
			}

			m[qname] = distance
		}
	}
}

type Location struct {
	aoc.Point2D
}

func (l Location) ID() string {
	return l.String()
}

func (l Location) Children() []aoc.Node {
	var children []aoc.Node
	if !grid[l.Up()] {
		children = append(children, Location{l.Up()})
	}
	if !grid[l.Down()] {
		children = append(children, Location{l.Down()})
	}
	if !grid[l.Left()] {
		children = append(children, Location{l.Left()})
	}
	if !grid[l.Right()] {
		children = append(children, Location{l.Right()})
	}
	return children
}

// PathDistance returns the distance between two points, as long as it doesn't
// contain any doors, keys or the starting location (other than the endpoints).
func PathDistance(start, end aoc.Point2D) int {
	isGoal := func(n aoc.Node) bool {
		return n.(Location).Point2D == end
	}

	cost := func(from, to aoc.Node) int {
		return 1
	}

	heuristic := func(n aoc.Node) int {
		return end.ManhattanDistance(n.(Location).Point2D)
	}

	path, distance, found := aoc.AStarSearch(Location{start}, isGoal, cost, heuristic)
	if !found {
		return -1
	}

	for _, node := range path[1 : distance-1] {
		p := node.(Location).Point2D
		if keys[p] != "" || doors[p] != "" {
			return -1
		}
	}

	return distance
}

func InputToGrid(year, day int) (Grid, Keys, Doors, aoc.Point2D) {
	var start aoc.Point2D
	grid := make(Grid)
	keys := make(Keys)
	doors := make(Doors)

	for y, line := range aoc.InputToLines(year, day) {
		for x, c := range line {
			p := aoc.Point2D{X: x, Y: y}

			if c == '@' {
				start = p
			} else if c == '#' {
				grid[p] = true
			} else if c >= 'a' && c <= 'z' {
				keys[p] = string(c)
			} else if c >= 'A' && c <= 'Z' {
				doors[p] = string(c)
			}
		}
	}

	return grid, keys, doors, start
}
