package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

var cave *Cave

func main() {
	cave = InputToCave(2018, 22)

	player := Player{
		location: aoc.Point2D{X: 0, Y: 0},
		equipped: "torch",
	}

	isGoal := func(node aoc.Node) bool {
		p := node.(Player)
		return p.location == cave.target && p.equipped == "torch"
	}

	cost := func(from, to aoc.Node) int {
		a := from.(Player)
		b := to.(Player)

		var cost int
		if a.location != b.location {
			cost += 1
		}
		if a.equipped != b.equipped {
			cost += 7
		}

		return cost
	}

	heuristic := func(node aoc.Node) int {
		p := node.(Player)

		var cost = p.location.ManhattanDistance(cave.target)
		if p.equipped != "torch" {
			cost += 7
		}

		return cost
	}

	path, distance, found := aoc.AStarSearch(player, isGoal, cost, heuristic)
	if !found {
		log.Fatal("unable to find path to target")
	}

	fmt.Printf("location: %10s | equipped: %10s | tm: %4d\n", path[0].(Player).location, path[0].(Player).equipped, 0)
	// fmt.Println(path[0].(Player))
	// fmt.Println()

	var tm int
	for i := 1; i < len(path); i++ {
		previous := path[i-1].(Player)
		current := path[i].(Player)
		tm += cost(previous, current)

		fmt.Printf("location: %10s | equipped: %10s | tm: %4d\n", current.location, current.equipped, tm)
		// fmt.Println(current)
		// fmt.Println()
	}

	fmt.Printf("distance: %v\n", distance)
}

type Player struct {
	location aoc.Point2D
	equipped string
}

func (p Player) ID() string {
	return fmt.Sprintf("%s|%s", p.location, p.equipped)
}

func (p Player) Children() []aoc.Node {
	var children []aoc.Node

	// First, consider changing equipment.
	currentTerrain := cave.GetType(p.location)

	var newEquipped string
	switch {
	case currentTerrain == Rocky && p.equipped == "torch":
		newEquipped = "climbing"
	case currentTerrain == Rocky && p.equipped == "climbing":
		newEquipped = "torch"
	case currentTerrain == Wet && p.equipped == "climbing":
		newEquipped = "neither"
	case currentTerrain == Wet && p.equipped == "neither":
		newEquipped = "climbing"
	case currentTerrain == Narrow && p.equipped == "torch":
		newEquipped = "neither"
	case currentTerrain == Narrow && p.equipped == "neither":
		newEquipped = "torch"
	}
	children = append(children, Player{
		location: p.location,
		equipped: newEquipped,
	})

	// Next, consider moving.
	var ps []aoc.Point2D

	if p.location.X > 0 {
		ps = append(ps, p.location.Left())
	}
	if p.location.Y > 0 {
		ps = append(ps, p.location.Up())
	}
	ps = append(ps, p.location.Right())
	ps = append(ps, p.location.Down())

	for _, location := range ps {
		next := cave.GetType(location)

		if next == Rocky && p.equipped == "neither" {
			continue
		}
		if next == Wet && p.equipped == "torch" {
			continue
		}
		if next == Narrow && p.equipped == "climbing" {
			continue
		}

		children = append(children, Player{
			location: location,
			equipped: p.equipped,
		})
	}

	return children
}

func (p Player) String() string {
	maxX, maxY := cave.target.X, cave.target.Y
	if p.location.X > maxX {
		maxX = p.location.X
	}
	if p.location.Y > maxY {
		maxY = p.location.Y
	}

	var builder strings.Builder
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			loc := aoc.Point2D{X: x, Y: y}
			if loc == p.location {
				builder.WriteString("X")
				continue
			}
			if loc == cave.target {
				builder.WriteString("T")
				continue
			}

			terrain := cave.GetType(loc)
			if terrain == Rocky {
				builder.WriteString(".")
			} else if terrain == Wet {
				builder.WriteString("=")
			} else {
				builder.WriteString("|")
			}
		}

		builder.WriteString("\n")
	}

	return builder.String()
}

type Cave struct {
	depth   int
	target  aoc.Point2D
	indices map[aoc.Point2D]int
}

const (
	Rocky  = 0
	Wet    = 1
	Narrow = 2
)

// Returns 0=rocky, 1=wet, 2=narrow denoting the type of the terrain at the
// given point.
func (c *Cave) GetType(p aoc.Point2D) int {
	return c.el(p) % 3
}

func (c *Cave) el(p aoc.Point2D) int {
	return (c.depth + c.gi(p)) % 20183
}

func (c *Cave) gi(p aoc.Point2D) int {
	if index, ok := c.indices[p]; ok {
		return index
	}

	var index int
	switch {
	case p.X == 0 && p.Y == 0:
		index = 0
	case p == c.target:
		index = 0
	case p.Y == 0:
		index = p.X * 16807
	case p.X == 0:
		index = p.Y * 48271
	default:
		index = c.el(p.Left()) * c.el(p.Up())
	}

	c.indices[p] = index
	return index
}

func InputToCave(year, day int) *Cave {
	depth, target := InputToDepthAndTarget(year, day)
	return &Cave{
		depth:   depth,
		target:  target,
		indices: make(map[aoc.Point2D]int),
	}
}

func InputToDepthAndTarget(year, day int) (int, aoc.Point2D) {
	var depth int
	var target aoc.Point2D
	for _, line := range aoc.InputToLines(year, day) {
		if strings.HasPrefix(line, "depth:") {
			depth = aoc.ParseInt(line[7:])
		} else if strings.HasPrefix(line, "target:") {
			parts := strings.Split(line[8:], ",")
			target.X = aoc.ParseInt(parts[0])
			target.Y = aoc.ParseInt(parts[1])
		} else {
			log.Fatalf("unrecognized line: %s", line)
		}
	}

	return depth, target
}
