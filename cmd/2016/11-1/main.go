package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	state := InputToState()

	goal := func(s State) bool {
		if s.Elevator != FloorCount-1 {
			return false
		}

		for e := 0; e < ElementCount; e++ {
			for c := 0; c < ComponentCount; c++ {
				if !s.Floors[FloorCount-1][e][c] {
					return false
				}
			}
		}

		return true
	}

	cost := func(from, to State) int {
		return 1
	}

	heuristic := func(s State) int {
		// Pretend that we have an elevator that can magically transport 2 items at a time from
		// a floor to the top floor, moving one floor at a time and ignoring any other constraints.
		var steps int
		for floor := 0; floor < FloorCount-1; floor++ {
			var count int
			for e := 0; e < ElementCount; e++ {
				for c := 0; c < ComponentCount; c++ {
					if s.Floors[floor][e][c] {
						count++
					}
				}
			}

			steps += (FloorCount-floor-1)*count/2 + (FloorCount-floor-1)*(count%2)
		}

		return steps
	}

	_, length, found := aoc.AStarSearch(state, Children, goal, cost, heuristic)
	if !found {
		fmt.Println("no path found")
	}

	fmt.Println(length)
}

func Children(s State) []State {
	valid := func(floor [ElementCount][ComponentCount]bool) bool {
		// if a chip is ever left in the same area as another RTG, and it's not
		// connected to its own RTG, the chip will be fried.
		var numGenerators, numUnshieldedChips int
		for i := 0; i < ElementCount; i++ {
			if floor[i][GeneratorIndex] {
				numGenerators++
				continue
			}

			if floor[i][MicrochipIndex] {
				numUnshieldedChips++
			}
		}

		return numUnshieldedChips == 0 || numGenerators == 0
	}

	var children []State
	for dy := -1; dy <= 1; dy += 2 {
		if s.Elevator+dy < 0 || s.Elevator+dy >= FloorCount {
			continue
		}

		for i := 0; i < ElementCount*ComponentCount; i++ {
			e1 := i / ComponentCount
			c1 := i % ComponentCount
			if !s.Floors[s.Elevator][e1][c1] {
				continue
			}

			// Consider moving only one component
			child := s // Copy
			child.Elevator = s.Elevator + dy
			child.Floors[s.Elevator][e1][c1] = false
			child.Floors[s.Elevator+dy][e1][c1] = true
			if valid(child.Floors[s.Elevator]) && valid(child.Floors[s.Elevator+dy]) {
				children = append(children, child)
			}

			// Consider moving a second component
			for j := i + 1; j < ElementCount*ComponentCount; j++ {
				e2 := j / ComponentCount
				c2 := j % ComponentCount
				if !s.Floors[s.Elevator][e2][c2] {
					continue
				}

				child := s // Copy
				child.Elevator = s.Elevator + dy
				child.Floors[s.Elevator][e1][c1] = false
				child.Floors[s.Elevator][e2][c2] = false
				child.Floors[s.Elevator+dy][e1][c1] = true
				child.Floors[s.Elevator+dy][e2][c2] = true
				if valid(child.Floors[s.Elevator]) && valid(child.Floors[s.Elevator+dy]) {
					children = append(children, child)
				}
			}
		}
	}

	return children
}

const (
	FloorCount     = 4
	ElementCount   = 5
	ComponentCount = 2 // Generator or Microchip

	GeneratorIndex = 0
	MicrochipIndex = 1
)

var Elements = map[string]int{
	"promethium": 0,
	"cobalt":     1,
	"curium":     2,
	"ruthenium":  3,
	"plutonium":  4,
}

var Components = map[string]int{
	"generator": 0,
	"microchip": 1,
}

type State struct {
	Elevator int
	Floors   [FloorCount][ElementCount][ComponentCount]bool
}

func (s State) Dump() string {
	var Abbreviations = [ElementCount][ComponentCount]string{
		{"PG", "PM"},
		{"CG", "CM"},
		{"UG", "UM"},
		{"RG", "RM"},
		{"LG", "LM"},
	}

	var sb strings.Builder
	for floor := FloorCount - 1; floor >= 0; floor-- {
		sb.WriteString(fmt.Sprintf("F%d ", floor))
		if s.Elevator == floor {
			sb.WriteString("E ")
		} else {
			sb.WriteString(". ")
		}

		for elem := 0; elem < ElementCount; elem++ {
			for component := 0; component < ComponentCount; component++ {
				if s.Floors[floor][elem][component] {
					sb.WriteString(Abbreviations[elem][component])
					sb.WriteString(" ")
				} else {
					sb.WriteString(".  ")
				}
			}
		}

		if floor > 0 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func InputToState() State {
	var state State
	for floor, line := range aoc.InputToLines(2016, 11) {
		line = strings.ReplaceAll(line, "The ", "")
		line = strings.ReplaceAll(line, " a ", " ")
		line = strings.ReplaceAll(line, " and ", " ")
		line = strings.ReplaceAll(line, "floor contains ", "")
		line = strings.ReplaceAll(line, "-compatible", "")
		line = strings.ReplaceAll(line, " nothing relevant", "")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ".", "")
		parts := strings.Split(line, " ")

		for i := 1; i < len(parts); i += 2 {
			state.Floors[floor][Elements[parts[i]]][Components[parts[i+1]]] = true
		}
	}

	return state
}
