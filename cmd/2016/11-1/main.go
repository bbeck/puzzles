package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	// promethium = P
	// cobalt = C
	// curium = U
	// ruthenium = R
	// plutonium = L
	state := &State{
		elevator: 0,
		floors: [][]string{
			{"PG", "PM"},
			{"CG", "UG", "RG", "LG"},
			{"CM", "UM", "RM", "LM"},
			{},
		},
	}

	goal := Goal(state)

	fmt.Printf("state id: %s\n", state.ID())
	fmt.Printf("goal id: %s\n", goal.ID())

	var length int
	aoc.BreadthFirstSearch(state, func(current aoc.Node) bool {
		if current.ID() == goal.ID() {
			length = current.(*State).depth
			return true
		}

		return false
	})

	fmt.Printf("length: %d\n", length)
}

type State struct {
	elevator int
	floors   [][]string
	id       string
	depth    int
}

func (s *State) IsValid() bool {
	// if a chip is ever left in the same area as another RTG, and it's not
	// connected to its own RTG, the chip will be fried
	valid := func(items []string) bool {
		// Index all of the items we have
		generators := make(map[byte]bool)
		chips := make(map[byte]bool)
		for _, item := range items {
			if item[1] == 'G' {
				generators[item[0]] = true
			} else {
				chips[item[0]] = true
			}
		}

		if len(generators) == 0 {
			return true
		}

		for chip := range chips {
			if !generators[chip] {
				return false
			}
		}

		return true
	}

	for _, items := range s.floors {
		if !valid(items) {
			return false
		}
	}

	return true
}

func (s *State) ID() string {
	if s.id != "" {
		return s.id
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("E:%d ", s.elevator+1))

	for floor, items := range s.floors {
		ids := append([]string(nil), items...)
		sort.Strings(ids)

		builder.WriteString(fmt.Sprintf("F%d:%s ", floor+1, strings.Join(ids, ",")))
	}

	s.id = builder.String()
	return s.id
}

func (s *State) Children() []aoc.Node {
	var neighbors []aoc.Node
	N := len(s.floors)
	items := s.floors[s.elevator]

	var neighbor *State
	if s.elevator > 0 {
		// The elevator can move down
		for i := 0; i < len(items); i++ {
			// The elevator can take a single item
			floors := Copy(s.floors)
			floors[s.elevator] = Without(floors[s.elevator], items[i], "")
			floors[s.elevator-1] = append(floors[s.elevator-1], items[i])
			neighbor = &State{
				elevator: s.elevator - 1,
				floors:   floors,
				depth:    s.depth + 1,
			}
			if neighbor.IsValid() {
				neighbors = append(neighbors, neighbor)
			}

			// The elevator can take a second item
			for j := i + 1; j < len(items); j++ {
				floors := Copy(s.floors)
				floors[s.elevator] = Without(floors[s.elevator], items[i], items[j])
				floors[s.elevator-1] = append(append(floors[s.elevator-1], items[i]), items[j])

				neighbor = &State{
					elevator: s.elevator - 1,
					floors:   floors,
					depth:    s.depth + 1,
				}
				if neighbor.IsValid() {
					neighbors = append(neighbors, neighbor)
				}
			}
		}
	}

	if s.elevator < N-1 {
		// The elevator can move up
		for i := 0; i < len(items); i++ {
			// The elevator can take a single item
			floors := Copy(s.floors)
			floors[s.elevator] = Without(floors[s.elevator], items[i], "")
			floors[s.elevator+1] = append(floors[s.elevator+1], items[i])
			neighbor = &State{
				elevator: s.elevator + 1,
				floors:   floors,
				depth:    s.depth + 1,
			}
			if neighbor.IsValid() {
				neighbors = append(neighbors, neighbor)
			}

			// The elevator can take a second item
			for j := i + 1; j < len(items); j++ {
				floors := Copy(s.floors)
				floors[s.elevator] = Without(floors[s.elevator], items[i], items[j])
				floors[s.elevator+1] = append(append(floors[s.elevator+1], items[i]), items[j])

				neighbor = &State{
					elevator: s.elevator + 1,
					floors:   floors,
					depth:    s.depth + 1,
				}
				if neighbor.IsValid() {
					neighbors = append(neighbors, neighbor)
				}
			}
		}
	}

	return neighbors
}

func (s *State) Print() {
	for floor := len(s.floors) - 1; floor >= 0; floor-- {
		fmt.Printf("F%d ", floor+1)

		if s.elevator == floor {
			fmt.Print("E ")
		} else {
			fmt.Print(". ")
		}

		for _, item := range s.floors[floor] {
			fmt.Printf("%s ", item)
		}

		fmt.Println()
	}
}

func Copy(floors [][]string) [][]string {
	next := make([][]string, len(floors))
	for i := 0; i < len(floors); i++ {
		next[i] = append([]string(nil), floors[i]...)
	}

	return next
}

func Without(floor []string, item1, item2 string) []string {
	next := make([]string, 0, len(floor)-1)
	for _, item := range floor {
		if item == item1 || item == item2 {
			continue
		}
		next = append(next, item)
	}
	return next
}

func Goal(s *State) *State {
	var top []string
	for _, items := range s.floors {
		top = append(top, items...)
	}

	floors := make([][]string, len(s.floors))
	for i := 0; i < len(s.floors)-1; i++ {
		floors[i] = []string{}
	}
	floors[len(s.floors)-1] = top

	return &State{
		elevator: len(s.floors) - 1,
		floors:   floors,
	}
}
