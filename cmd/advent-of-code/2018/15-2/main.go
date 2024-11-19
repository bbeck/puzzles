package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"math"
	"sort"
)

func main() {
	// A binary search on AP doesn't work here because it's not always true that
	// a success at one AP means success at all higher APs.  Consider the case
	// where a higher AP elf finishes off a gnome more quickly and gets into a
	// new fight.  That elf may have just blocked a higher health one that was
	// getting into position to attack.  Thus, we have a lower health elf fighting
	// and could take a loss.
	//
	// Considering the number of attacks it takes to finish off a gnome, does
	// work, so we'll binary search on that instead.
	outcomes := make(map[int]int) // outcomes indexed by number of attacks

	// Our search space is "backwards" since a lower number of attacks results
	// in a higher AP.  So search will return the number of attacks value that
	// results in our first loss.  We'll need the one before that.
	attacks := -1 + sort.Search(200, func(n int) bool {
		// Compute the AP that's needed to finish a 200HP gnome
		ap := int(math.Ceil(200 / float64(n)))

		var success bool
		outcomes[n], success = Simulate(ap)
		return !success
	})

	fmt.Println(outcomes[attacks])
}

func Simulate(ap int) (int, bool) {
	cavern, units := InputToCavern(), InputToUnits(ap)

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
			// Terminate with a failure if an elf dies.
			if unit.Kind == "E" && unit.HP <= 0 {
				return 0, false
			}

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
	return round * hps, true
}

func TurnOrder(units []Unit) func(int, int) bool {
	return func(i int, j int) bool {
		return units[i].Y < units[j].Y ||
			(units[i].Y == units[j].Y && units[i].X < units[j].X)
	}
}

func ReadingOrder(ps []lib.Point2D) func(int, int) bool {
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

func TakeTurn(unit *Unit, cavern lib.Grid2D[bool], units []Unit) {
	if unit.HP <= 0 {
		return
	}

	var occupied lib.Set[lib.Point2D]
	for _, u := range units {
		if u.HP > 0 && u.Point2D != unit.Point2D {
			occupied.Add(u.Point2D)
		}
	}

	enemies := make(map[lib.Point2D]int)
	for index, u := range units {
		if u.HP > 0 && u.Kind != unit.Kind {
			enemies[u.Point2D] = index
		}
	}

	// Attempt to move.  Start by computing all possible targets for this unit.
	// A target is an open cell adjacent to an enemy.
	var candidates lib.Set[lib.Point2D]
	for target := range enemies {
		for _, p := range target.OrthogonalNeighbors() {
			if cavern.GetPoint(p) && !occupied.Contains(p) {
				candidates.Add(p)
			}
		}
	}

	// If we're already at one of the candidate positions then no move is
	// necessary.  We'll use an empty targets slice in this situation.
	var targets []lib.Point2D
	if !candidates.Contains(unit.Point2D) {
		targets = candidates.Entries()
		sort.Slice(targets, ReadingOrder(targets))
	}

	// This unit can move to one of its neighboring cells.  Choose the neighboring
	// cell that's closest to a target cell.
	var choice lib.Point2D
	best := math.MaxInt
	for _, end := range targets {
		for _, start := range []lib.Point2D{unit.Up(), unit.Left(), unit.Right(), unit.Down()} {
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
	for _, p := range []lib.Point2D{unit.Up(), unit.Left(), unit.Right(), unit.Down()} {
		if index, found := enemies[p]; found && (attack == -1 || units[index].HP < units[attack].HP) {
			attack = index
		}
	}
	if attack != -1 {
		units[attack].HP -= unit.AP
	}
}

func Distance(start, end lib.Point2D, cavern lib.Grid2D[bool], occupied lib.Set[lib.Point2D]) int {
	children := func(p lib.Point2D) []lib.Point2D {
		var children []lib.Point2D
		for _, neighbor := range []lib.Point2D{p.Up(), p.Left(), p.Right(), p.Down()} {
			if cavern.InBoundsPoint(neighbor) && cavern.GetPoint(neighbor) && !occupied.Contains(neighbor) {
				children = append(children, neighbor)
			}
		}
		return children
	}

	isGoal := func(p lib.Point2D) bool {
		return p == end
	}

	path, found := lib.BreadthFirstSearch(start, children, isGoal)
	if !found {
		return math.MaxInt
	}
	return len(path)
}

type Unit struct {
	lib.Point2D
	Kind string
	HP   int
	AP   int
}

func InputToCavern() lib.Grid2D[bool] {
	return lib.InputToGrid2D(func(x int, y int, s string) bool {
		return s != "#"
	})
}

func InputToUnits(elfAP int) []Unit {
	lines := lib.InputToLines()

	var units []Unit
	for y := 0; y < len(lines); y++ {
		for x, c := range lines[y] {
			var ap = 3
			if c == 'E' {
				ap = elfAP
			}
			if c == 'G' || c == 'E' {
				units = append(units, Unit{
					Kind:    string(c),
					Point2D: lib.Point2D{X: x, Y: y},
					HP:      200,
					AP:      ap,
				})
			}
		}
	}

	return units
}
