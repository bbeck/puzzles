package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

const (
	W = 13
	H = 7
)

func main() {
	start := InputToInitialState()

	heuristic := func(s State) int {
		var cost int
		for _, p := range Hallway {
			if kind := s.Board[p.X][p.Y]; kind != '.' {
				cost += lib.Abs(Rooms[kind][0].X-p.X) * Energy[kind]
			}
		}
		for _, rs := range Rooms {
			for _, p := range rs {
				if kind := s.Board[p.X][p.Y]; kind != '.' {
					cost += (lib.Abs(Rooms[kind][0].X-p.X) + 2) * Energy[kind]
				}
			}
		}

		return cost
	}

	cost := func(from, to State) int {
		return to.Energy - from.Energy
	}

	id := func(s State) [W][H]byte {
		return s.Board
	}

	_, energy, _ := lib.AStarSearchWithIdentity(start, Children, IsGoal, cost, heuristic, id)
	fmt.Println(energy)
}

var Hallway = []lib.Point2D{
	{X: 1, Y: 1},
	{X: 2, Y: 1},
	{X: 4, Y: 1},
	{X: 6, Y: 1},
	{X: 8, Y: 1},
	{X: 10, Y: 1},
	{X: 11, Y: 1},
}

var Rooms = map[byte][]lib.Point2D{
	'A': {{X: 3, Y: 2}, {X: 3, Y: 3}, {X: 3, Y: 4}, {X: 3, Y: 5}},
	'B': {{X: 5, Y: 2}, {X: 5, Y: 3}, {X: 5, Y: 4}, {X: 5, Y: 5}},
	'C': {{X: 7, Y: 2}, {X: 7, Y: 3}, {X: 7, Y: 4}, {X: 7, Y: 5}},
	'D': {{X: 9, Y: 2}, {X: 9, Y: 3}, {X: 9, Y: 4}, {X: 9, Y: 5}},
}

var Energy = map[byte]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}

func IsGoal(s State) bool {
	for kind, ps := range Rooms {
		for _, p := range ps {
			if s.Board[p.X][p.Y] != kind {
				return false
			}
		}
	}
	return true
}

func Children(s State) []State {
	var children []State

	// Consider moving from the hallway into a room
	for _, p := range Hallway {
		kind := s.Board[p.X][p.Y]
		if kind == '.' {
			continue
		}

		var occupied bool
		var target lib.Point2D
		for _, r := range Rooms[kind] {
			k := s.Board[r.X][r.Y]
			if k == '.' {
				target = r
			} else if k != kind {
				occupied = true
			}
		}

		if ok, length := HasPath(s, p, target); !occupied && ok {
			children = append(children, State{
				Board:  Move(s.Board, p, target),
				Energy: s.Energy + length*Energy[kind],
			})
		}
	}

	// Consider moving the topmost entry in a room into the hallway, but only
	// if it's the wrong kind or is blocking a wrong kind.
	for kind, rs := range Rooms {
		var top *lib.Point2D
		var blocking bool

		for _, r := range rs {
			if s.Board[r.X][r.Y] != '.' && top == nil {
				cpy := r
				top = &cpy
				continue
			}

			if top != nil && s.Board[r.X][r.Y] != kind {
				blocking = true
				break
			}
		}

		if top == nil {
			continue
		}

		if blocking || s.Board[top.X][top.Y] != kind {
			for _, target := range Hallway {
				if ok, length := HasPath(s, *top, target); ok && s.Board[target.X][target.Y] == '.' {
					children = append(children, State{
						Board:  Move(s.Board, *top, target),
						Energy: s.Energy + length*Energy[s.Board[top.X][top.Y]],
					})
				}
			}
		}
	}

	return children
}

func HasPath(s State, a, b lib.Point2D) (bool, int) {
	var length int
	for a.Y != 1 {
		a = a.Up()
		length++

		if s.Board[a.X][a.Y] != '.' {
			return false, 0
		}
	}

	for a.X != b.X {
		if a.X < b.X {
			a = a.Right()
		} else {
			a = a.Left()
		}
		length++

		if s.Board[a.X][a.Y] != '.' {
			return false, 0
		}
	}

	for a.Y != b.Y {
		a = a.Down()
		length++

		if s.Board[a.X][a.Y] != '.' {
			return false, 0
		}
	}

	return true, length
}

func Move(board [W][H]byte, a, b lib.Point2D) [W][H]byte {
	board[b.X][b.Y] = board[a.X][a.Y]
	board[a.X][a.Y] = '.'
	return board
}

type State struct {
	Board  [W][H]byte
	Energy int
}

func InputToInitialState() State {
	folded := lib.InputToLines()
	lines := append([]string{}, folded[:3]...)
	lines = append(lines, "  #D#C#B#A#  ", "  #D#B#A#C#  ")
	lines = append(lines, folded[3:]...)

	var board [W][H]byte
	for y, line := range lines {
		for x, c := range line {
			board[x][y] = byte(c)
		}
	}

	return State{Board: board}
}
