package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	L := 4
	U := 200

	outcomes := make(map[int]int) // mapping from successful AP to the outcome
	for L <= U {
		ap := (L + U) / 2

		result, outcome := simulate(ap)
		if result {
			U = ap - 1
			outcomes[ap] = outcome
		} else {
			L = ap + 1
		}
	}

	for ap := 4; ; ap++ {
		if outcomes[ap] > 0 {
			fmt.Printf("ap: %d, outcome: %d\n", ap, outcomes[ap])
			break
		}
	}
}

func simulate(ap int) (bool, int) {
	cavern := InputToCavern(2018, 15, ap)
	// fmt.Printf("=== Initial =======================================\n")
	// fmt.Println(cavern)

	var round int
	for round = 1; ; round++ {
		// fmt.Printf("=== Round %2d =======================================\n", round)

		var done bool
		for _, unit := range cavern.TurnOrder() {
			// Check if the game is over.  If so then we have a partial round.
			if len(cavern.elves) == 0 || len(cavern.goblins) == 0 {
				done = true
				round--

				// fmt.Printf("  Ended early\n")
				// fmt.Println()
				break
			}

			if unit.hp <= 0 {
				// The unit died earlier in the round
				continue
			}

			// Move (if necessary)
			cavern.Move(unit)

			// Attack
			if cavern.Attack(unit) {
				return false, 0
			}
		}

		if done {
			break
		}

		// fmt.Printf("At end of round %d, board is:\n", round)
		// fmt.Println(cavern)
	}

	// fmt.Printf("=== Final =======================================\n")
	// fmt.Println(cavern)

	var hps int
	for _, elf := range cavern.elves {
		hps += elf.hp
	}
	for _, goblin := range cavern.goblins {
		hps += goblin.hp
	}

	// fmt.Printf("full rounds: %d, hps: %d, outcome: %d\n", round, hps, round*hps)
	return true, round * hps
}

const (
	WALL  string = "#"
	EMPTY        = "."
)

type Cavern struct {
	width   int
	height  int
	cells   map[aoc.Point2D]string
	elves   map[aoc.Point2D]*Unit
	goblins map[aoc.Point2D]*Unit
}

func (c *Cavern) TurnOrder() []*Unit {
	var units []*Unit
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			p := aoc.Point2D{X: x, Y: y}
			if elf, ok := c.elves[p]; ok {
				units = append(units, elf)
			}
			if goblin, ok := c.goblins[p]; ok {
				units = append(units, goblin)
			}
		}
	}

	return units
}

func (c *Cavern) Move(u *Unit) {
	enemies := c.goblins
	if u.side == "G" {
		enemies = c.elves
	}

	// First, determine the distance and path to each cell that's adjacent to an
	// enemy.
	targets := make([]aoc.Point2D, 0)
	distances := make(map[aoc.Point2D]int)
	for _, enemy := range enemies {
		neighbors := []aoc.Point2D{
			enemy.location.Up(),
			enemy.location.Left(),
			enemy.location.Right(),
			enemy.location.Down(),
		}
		for _, neighbor := range neighbors {
			if neighbor == u.location {
				targets = append(targets, neighbor)
				distances[neighbor] = 0
				continue
			}

			if c.cells[neighbor] == WALL {
				continue
			}

			if c.elves[neighbor] != nil {
				continue
			}

			if c.goblins[neighbor] != nil {
				continue
			}

			heuristic := func(node aoc.Node) int {
				p := node.(*Unit).location
				return p.ManhattanDistance(neighbor)
			}

			isGoal := func(node aoc.Node) bool {
				p := node.(*Unit).location
				return p == neighbor
			}

			cost := func(from, to aoc.Node) int {
				return 1
			}

			// We can't use the path discovered by the A* search because it doesn't
			// always obey reading order -- we don't have control over the tie breaks
			// within the priority queue.
			_, distance, found := aoc.AStarSearch(u, isGoal, cost, heuristic)
			if !found {
				// There wasn't a path, nothing to remember.
				continue
			}

			targets = append(targets, neighbor)
			distances[neighbor] = distance
		}
	}

	if len(targets) == 0 {
		// There wasn't a valid target, we can't move.
		return
	}

	// Order the targets in reading order
	sort.Slice(targets, func(i, j int) bool {
		return targets[i].Y < targets[j].Y ||
			(targets[i].Y == targets[j].Y && targets[i].X < targets[j].X)
	})

	// Now that we know all of the targets, choose the first one with the shortest
	// distance.
	var target aoc.Point2D
	var distance = math.MaxInt64
	for _, t := range targets {
		if distances[t] < distance {
			target = t
			distance = distances[t]
		}
	}

	// If the distance is 0 then we're already in the place we want to be.  No
	// need to move.
	if distance == 0 {
		return
	}

	// Now determine the reading order path to our target.  We know a path exists
	// due to the A* search we already performed.
	path, _ := aoc.BreadthFirstSearch(u, func(node aoc.Node) bool {
		return node.(*Unit).location == target
	})

	next := path[1].(*Unit)
	if u.side == "E" {
		delete(c.elves, u.location)
		c.elves[next.location] = u
		u.location = next.location
	} else {
		delete(c.goblins, u.location)
		c.goblins[next.location] = u
		u.location = next.location
	}
}

func (c *Cavern) Attack(u *Unit) bool {
	// These are the possible locations of a target in reading order.
	ps := []aoc.Point2D{
		u.location.Up(),
		u.location.Left(),
		u.location.Right(),
		u.location.Down(),
	}

	var target *Unit
	for _, p := range ps {
		switch u.side {
		case "E":
			if goblin, ok := c.goblins[p]; ok {
				if target == nil || goblin.hp < target.hp {
					target = goblin
				}
			}
		case "G":
			if elf, ok := c.elves[p]; ok {
				if target == nil || elf.hp < target.hp {
					target = elf
				}
			}
		}
	}

	// There wasn't a target to attack, skip the attacking.
	if target == nil {
		return false
	}

	target.hp -= u.ap
	if target.hp <= 0 {
		switch target.side {
		case "E":
			delete(c.elves, target.location)
			return true

		case "G":
			delete(c.goblins, target.location)
		}
	}

	return false
}

func (c *Cavern) String() string {
	var builder strings.Builder
	for y := 0; y < c.height; y++ {
		var units []string

		builder.WriteString("  ")
		for x := 0; x < c.width; x++ {
			p := aoc.Point2D{X: x, Y: y}

			if c.cells[p] == WALL {
				builder.WriteString("█")
			} else if elf, ok := c.elves[p]; ok {
				builder.WriteString(elf.side)
				units = append(units, fmt.Sprintf("E(%d)", elf.hp))
			} else if goblin, ok := c.goblins[p]; ok {
				builder.WriteString(goblin.side)
				units = append(units, fmt.Sprintf("G(%d)", goblin.hp))
			} else {
				builder.WriteString(EMPTY)
			}
		}

		builder.WriteString("  ")
		builder.WriteString(strings.Join(units, ", "))
		if y < c.height-1 {
			builder.WriteString("\n")
		}
	}

	return builder.String()
}

type Unit struct {
	side     string
	location aoc.Point2D
	ap       int
	hp       int
	cavern   *Cavern
}

func (u *Unit) ID() string {
	return u.location.String()
}

func (u *Unit) Children() []aoc.Node {
	ps := []aoc.Point2D{
		u.location.Up(),
		u.location.Left(),
		u.location.Right(),
		u.location.Down(),
	}

	var children []aoc.Node
	for _, p := range ps {
		if cell, ok := u.cavern.cells[p]; !ok || cell == WALL {
			continue
		}

		if _, ok := u.cavern.elves[p]; ok {
			continue
		}

		if _, ok := u.cavern.goblins[p]; ok {
			continue
		}

		children = append(children, &Unit{
			side:     u.side,
			location: p,
			hp:       u.hp,
			cavern:   u.cavern,
		})
	}

	return children
}

func InputToCavern(year, day int, elfAP int) *Cavern {
	cavern := new(Cavern)
	width := 0
	height := 0
	cells := make(map[aoc.Point2D]string)
	elves := make(map[aoc.Point2D]*Unit)
	goblins := make(map[aoc.Point2D]*Unit)
	for y, line := range aoc.InputToLines(year, day) {
		for x, c := range line {
			p := aoc.Point2D{X: x, Y: y}
			switch c {
			case '#':
				cells[p] = WALL
			case '.':
				cells[p] = EMPTY
			case 'E':
				cells[p] = EMPTY
				elves[p] = &Unit{side: "E", location: p, ap: elfAP, hp: 200, cavern: cavern}
			case 'G':
				cells[p] = EMPTY
				goblins[p] = &Unit{side: "G", location: p, ap: 3, hp: 200, cavern: cavern}
			}
			width = x + 1
		}
		height = y + 1
	}

	cavern.width = width
	cavern.height = height
	cavern.cells = cells
	cavern.elves = elves
	cavern.goblins = goblins
	return cavern
}
