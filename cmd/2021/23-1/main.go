package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"log"
	"sort"
	"strings"
)

var board = InputToBoard()

func main() {
	players := make(map[aoc.Point2D]Player)
	for _, p := range InputToPlayers() {
		players[p.Point2D] = p
	}
	start := State{
		Players: players,
	}

	isGoal := func(n aoc.Node) bool {
		state := n.(State)
		for _, player := range state.Players {
			if !state.IsDone(player) {
				return false
			}
		}
		return true
	}

	cost := func(from, to aoc.Node) int {
		return to.(State).Energy - from.(State).Energy
	}

	heuristic := func(n aoc.Node) int {
		return 1
	}

	_, energy, found := aoc.AStarSearch(start, isGoal, cost, heuristic)
	if !found {
		log.Fatal("no solution found")
	}
	fmt.Println(energy)
}

type State struct {
	Players map[aoc.Point2D]Player
	Energy  int
}

func (s State) ID() string {
	var players []Player
	for _, p := range s.Players {
		players = append(players, p)
	}

	sort.Slice(players, func(i, j int) bool {
		if players[i].Kind != players[j].Kind {
			return players[i].Kind < players[j].Kind
		}
		if players[i].Y != players[j].Y {
			return players[i].Y < players[j].Y
		}
		return players[i].X < players[j].X
	})

	var sb strings.Builder
	for i, p := range players {
		sb.WriteString(p.Kind)
		sb.WriteRune('@')
		sb.WriteString(p.String())
		if i < len(players)-1 {
			sb.WriteRune('|')
		}
	}
	return sb.String()
}

var RoomX = map[string]int{
	"A": 3,
	"B": 5,
	"C": 7,
	"D": 9,
}

var RoomY = []int{
	2,
	3,
}

// HallwayX contains the x-coordinates in the hallway that a player can stop at.
var HallwayX = []int{1, 2, 4, 6, 8, 10, 11}

func (s State) Children() []aoc.Node {
	var children []aoc.Node
	for _, p := range s.Players {
		// Consider moving the player if:
		//   - they are in the hallway
		//   - they are in a different player's room
		//   - they are in their room, but they are blocking another player
		isInHallway := p.Y == 1
		isInWrongRoom := p.X != RoomX[p.Kind]
		isBlocking := s.Players[p.Down()].Kind != "" && s.Players[p.Down()].Kind != p.Kind
		if !isInHallway && !isInWrongRoom && !isBlocking {
			continue
		}

		if !isInHallway {
			// This player is in a room, consider moving them into any hallway space that
			// isn't occupied and is reachable.
			for _, x := range HallwayX {
				target := aoc.Point2D{X: x, Y: 1}
				if _, occupied := s.Players[target]; occupied {
					continue
				}

				path := Path(p.Point2D, target)
				if !s.IsTraversable(path) {
					continue
				}

				children = append(children, s.Move(p, path[len(path)-1], len(path)))
			}
		}

		if isInHallway {
			// This player is in the hallway, their next move needs to be into their room.
			for _, y := range RoomY {
				target := aoc.Point2D{X: RoomX[p.Kind], Y: y}
				if _, occupied := s.Players[target]; occupied {
					continue
				}

				path := Path(p.Point2D, target)
				if !s.IsTraversable(path) {
					continue
				}

				pNext := Player{Point2D: target, Kind: p.Kind}
				if !s.IsDone(pNext) {
					continue
				}

				children = append(children, s.Move(p, path[len(path)-1], len(path)))
			}
		}

	}
	return children
}

var StepEnergy = map[string]int{
	"A": 1,
	"B": 10,
	"C": 100,
	"D": 1000,
}

func (s State) Move(player Player, point aoc.Point2D, steps int) State {
	players := map[aoc.Point2D]Player{
		point: {Point2D: point, Kind: player.Kind},
	}
	for opoint, oplayer := range s.Players {
		if oplayer == player {
			continue
		}

		players[opoint] = oplayer
	}

	return State{
		Players: players,
		Energy:  s.Energy + steps*StepEnergy[player.Kind],
	}
}

func (s State) String() string {
	var points []aoc.Point2D
	for p := range board {
		points = append(points, p)
	}
	minX, minY, maxX, maxY := aoc.GetBounds(points)

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("energy: %d\n", s.Energy))
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := aoc.Point2D{X: x, Y: y}

			if player, found := s.Players[p]; found {
				sb.WriteString(player.Kind)
			} else if board[p] {
				sb.WriteString(" ")
			} else {
				sb.WriteString("░")
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (s State) IsTraversable(path []aoc.Point2D) bool {
	// A path is traversable if there are no other players on it.
	for _, p := range path {
		if _, found := s.Players[p]; found {
			return false
		}
	}
	return true
}

// IsDone determines if the specified player is in a goal state for them.
func (s State) IsDone(p Player) bool {
	// They can't be in a hallway
	if p.Y == 1 {
		return false
	}

	// They have to be in their assigned room
	if p.X != RoomX[p.Kind] {
		return false
	}

	// There shouldn't be any other player kinds in their room with them
	for _, y := range RoomY {
		if other, found := s.Players[aoc.Point2D{X: p.X, Y: y}]; found {
			if other.Kind != p.Kind {
				return false
			}
		}
	}

	// There shouldn't be any blank spots below them
	for _, y := range RoomY {
		if y <= p.Y {
			continue
		}

		if _, found := s.Players[aoc.Point2D{X: p.X, Y: y}]; !found {
			return false
		}
	}

	return true
}

type Player struct {
	aoc.Point2D
	Kind string
}

func InputToPlayers() []Player {
	var players []Player
	for y, line := range aoc.InputToLines(2021, 23) {
		for x, c := range line {
			if c == 'A' || c == 'B' || c == 'C' || c == 'D' {
				players = append(players, Player{aoc.Point2D{X: x, Y: y}, string(c)})
			}
		}
	}

	return players
}

func Path(p, end aoc.Point2D) []aoc.Point2D {
	dx := 1
	if p.X > end.X {
		dx = -1
	}

	dy := 1
	if p.Y > end.Y {
		dy = -1
	}

	var path []aoc.Point2D
	moveLR := func() {
		for p.X != end.X {
			p = aoc.Point2D{X: p.X + dx, Y: p.Y}
			path = append(path, p)
		}
	}
	moveUD := func() {
		for p.Y != end.Y {
			p = aoc.Point2D{X: p.X, Y: p.Y + dy}
			path = append(path, p)
		}
	}

	if p.Y == 1 {
		// The starting point is in the hallway, first move left/right
		moveLR()
		moveUD()
	} else {
		// The starting point is in a room, first move up/down
		moveUD()
		moveLR()
	}

	return path
}

type Board map[aoc.Point2D]bool

func InputToBoard() Board {
	board := make(Board)
	for y, line := range aoc.InputToLines(2021, 23) {
		for x, c := range line {
			if c == '.' || c == 'A' || c == 'B' || c == 'C' || c == 'D' {
				board[aoc.Point2D{X: x, Y: y}] = true
			} else {
				board[aoc.Point2D{X: x, Y: y}] = false
			}
		}
	}

	return board
}
