package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
	"math"
	"sort"
)

func main() {
	cavern, units := InputToCavernAndUnits()

	var round int
	for round = 1; ; round++ {
		sort.Slice(units, TurnOrder(units))

		var done bool
		for i := 0; !done && i < len(units); i++ {
			TakeTurn(&units[i], cavern, units)

			if IsGameOver(units) {
				done = true
				if i != len(units)-1 {
					round--
				}
			}
		}

		// Remove the dead
		var next []Unit
		for _, unit := range units {
			if unit.HP > 0 {
				next = append(next, unit)
			}
		}
		units = next

		if done {
			break
		}
	}

	var hps int
	for _, unit := range units {
		hps += unit.HP
	}
	fmt.Println(round * hps)
}

func TurnOrder(units []Unit) func(int, int) bool {
	return func(i int, j int) bool {
		return units[i].Y < units[j].Y ||
			(units[i].Y == units[j].Y && units[i].X < units[j].X)
	}
}

func ReadingOrder(ps []Point2D) func(int, int) bool {
	return func(i int, j int) bool {
		return ps[i].Y < ps[j].Y ||
			(ps[i].Y == ps[j].Y && ps[i].X < ps[j].X)
	}
}

func IsGameOver(units []Unit) bool {
	var foundE, foundG bool
	for i := 0; i < len(units) && (!foundE || !foundG); i++ {
		if units[i].HP <= 0 {
			continue
		}

		switch units[i].Kind {
		case "E":
			foundE = true
		case "G":
			foundG = true
		}
	}

	return !foundE || !foundG
}

func TakeTurn(unit *Unit, cavern Grid2D[bool], units []Unit) {
	if unit.HP <= 0 {
		return
	}

	var occupied Set[Point2D]
	for _, u := range units {
		if u.HP > 0 && u.Point2D != unit.Point2D {
			occupied.Add(u.Point2D)
		}
	}

	enemies := make(map[Point2D]int)
	for index, u := range units {
		if u.HP > 0 && u.Kind != unit.Kind {
			enemies[u.Point2D] = index
		}
	}

	// Attempt to move.  Start by computing all possible targets for this unit.
	// A target is an open cell adjacent to an enemy.
	var candidates Set[Point2D]
	for target := range enemies {
		for _, p := range target.OrthogonalNeighbors() {
			if cavern.GetPoint(p) && !occupied.Contains(p) {
				candidates.Add(p)
			}
		}
	}

	// If we're already at one of the candidate positions then no move is
	// necessary.  We'll use an empty targets slice in this situation.
	var targets []Point2D
	if !candidates.Contains(unit.Point2D) {
		targets = candidates.Entries()
		sort.Slice(targets, ReadingOrder(targets))
	}

	// This unit can move to one of its neighboring cells.  Choose the neighboring
	// cell that's closest to a target cell.
	var choice Point2D
	best := math.MaxInt
	for _, end := range targets {
		for _, start := range []Point2D{unit.Up(), unit.Left(), unit.Right(), unit.Down()} {
			if !cavern.GetPoint(start) || occupied.Contains(start) {
				continue
			}

			distance := Distance(start, end, cavern, occupied)
			if distance < best {
				best = distance
				choice = start
			}
		}
	}

	if best < math.MaxInt {
		unit.Point2D = choice
	}

	// Now determine if this unit is in range of an enemy to attack.  If multiple
	// enemies are in range the one with the lowest hit points is chosen.
	attack := -1
	for _, p := range []Point2D{unit.Up(), unit.Left(), unit.Right(), unit.Down()} {
		if index, found := enemies[p]; found && (attack == -1 || units[index].HP < units[attack].HP) {
			attack = index
		}
	}
	if attack != -1 {
		units[attack].HP -= unit.AP
	}
}

func Distance(start, end Point2D, cavern Grid2D[bool], occupied Set[Point2D]) int {
	children := func(p Point2D) []Point2D {
		var children []Point2D
		for _, neighbor := range []Point2D{p.Up(), p.Left(), p.Right(), p.Down()} {
			if cavern.InBoundsPoint(neighbor) && cavern.GetPoint(neighbor) && !occupied.Contains(neighbor) {
				children = append(children, neighbor)
			}
		}
		return children
	}

	isGoal := func(p Point2D) bool {
		return p == end
	}

	path, found := BreadthFirstSearch(start, children, isGoal)
	if !found {
		return math.MaxInt
	}
	return len(path)
}

type Unit struct {
	Point2D
	Kind string
	HP   int
	AP   int
}

func InputToCavernAndUnits() (Grid2D[bool], []Unit) {
	var grid = in.ToGrid2D(func(_, _ int, s string) string { return s })

	var cavern = NewGrid2D[bool](grid.Width, grid.Height)

	var units []Unit
	grid.ForEachPoint(func(p Point2D, s string) {
		cavern.SetPoint(p, s != "#")

		switch s {
		case "E":
			units = append(units, Unit{Kind: s, Point2D: p, HP: 200, AP: 3})

		case "G":
			units = append(units, Unit{Kind: s, Point2D: p, HP: 200, AP: 3})
		}
	})

	return cavern, units
}
