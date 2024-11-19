package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	state := State{
		Turn: "player",
		Player: Player{
			HitPoints: 50,
			Mana:      500,
		},
		Boss: InputToBoss(),
	}

	isGoal := func(s State) bool {
		return s.Boss.HitPoints <= 0
	}

	cost := func(from, to State) int {
		return to.Spent - from.Spent
	}

	heuristic := func(s State) int {
		// At most, we can do 7 hit points of damage per turn (magic missile + active
		// poison).  We pay the cost of the magic missile every turn, but poison only
		// every 6th turn.
		turns := s.Boss.HitPoints / 7
		return 53*turns + 173*turns/6
	}

	_, mana, found := lib.AStarSearch(state, children, isGoal, cost, heuristic)
	if found {
		fmt.Println(mana)
	}
}

func children(parent State) []State {
	// At the start of each turn the player loses 1 hit point.
	parent.Player.HitPoints--

	// The match is over if either character drops to 0 hit points or the player
	// doesn't have enough mana to cast a spell (on their turn).
	if parent.Player.HitPoints <= 0 || parent.Boss.HitPoints <= 0 {
		return nil
	}
	if parent.Turn == "player" && parent.Player.Mana < 53 {
		return nil
	}

	// Apply effects
	if parent.Player.DrainTurns > 0 {
		parent.Player.DrainTurns--
		parent.Player.HitPoints += 2
		parent.Boss.HitPoints -= 2
	}
	if parent.Player.ShieldTurns > 0 {
		parent.Player.ShieldTurns--
		parent.Player.Armor = 7
	} else {
		parent.Player.Armor = 0
	}
	if parent.Player.PoisonTurns > 0 {
		parent.Player.PoisonTurns--
		parent.Boss.HitPoints -= 3
	}
	if parent.Player.RechargeTurns > 0 {
		parent.Player.RechargeTurns--
		parent.Player.Mana += 101
	}

	if parent.Turn == "boss" {
		return []State{
			parent.BossAttacks(),
		}
	}

	// It's the player's turn
	var children []State

	if parent.Player.Mana >= 53 {
		children = append(children, parent.PlayerCastsMagicMissile())
	}
	if parent.Player.Mana >= 73 {
		children = append(children, parent.PlayerCastsDrain())
	}
	if parent.Player.Mana >= 113 && parent.Player.ShieldTurns == 0 {
		children = append(children, parent.PlayerCastsShield())
	}
	if parent.Player.Mana >= 173 && parent.Player.PoisonTurns == 0 {
		children = append(children, parent.PlayerCastsPoison())
	}
	if parent.Player.Mana >= 229 && parent.Player.RechargeTurns == 0 {
		children = append(children, parent.PlayerCastsRecharge())
	}

	return children
}

type Player struct {
	HitPoints int
	Mana      int
	Armor     int

	// Since this struct is going to be used in maps and sets it can't contain any
	// fields that make it unsuitable as a map key.
	DrainTurns    int
	ShieldTurns   int
	PoisonTurns   int
	RechargeTurns int
}

func (p Player) Copy() Player {
	return p
}

type Boss struct {
	HitPoints int
	Damage    int
}

type State struct {
	Player Player
	Boss   Boss

	// Whose turn it is, "boss" or "player"
	Turn string

	// How much mana has been spent
	Spent int
}

func (s State) PlayerCastsMagicMissile() State {
	s.Player.Mana -= 53
	s.Boss.HitPoints -= 4
	s.Turn = "boss"
	s.Spent += 53
	return s
}

func (s State) PlayerCastsDrain() State {
	s.Player.Mana -= 73
	s.Player.HitPoints += 2
	s.Boss.HitPoints -= 2
	s.Turn = "boss"
	s.Spent += 73
	return s
}

func (s State) PlayerCastsShield() State {
	s.Player.Mana -= 113
	s.Player.ShieldTurns = 6
	s.Turn = "boss"
	s.Spent += 113
	return s
}

func (s State) PlayerCastsPoison() State {
	s.Player.Mana -= 173
	s.Player.PoisonTurns = 6
	s.Turn = "boss"
	s.Spent += 173
	return s
}

func (s State) PlayerCastsRecharge() State {
	s.Player.Mana -= 229
	s.Player.RechargeTurns = 5
	s.Turn = "boss"
	s.Spent += 229
	return s
}

func (s State) BossAttacks() State {
	s.Player.HitPoints -= lib.Max(1, s.Boss.Damage-s.Player.Armor)
	s.Turn = "player"
	return s
}

func InputToBoss() Boss {
	var boss Boss
	for _, line := range lib.InputToLines() {
		fmt.Sscanf(line, "Hit Points: %d", &boss.HitPoints)
		fmt.Sscanf(line, "Damage: %d", &boss.Damage)
	}

	return boss
}
