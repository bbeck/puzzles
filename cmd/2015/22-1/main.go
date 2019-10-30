package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

type State struct {
	parent   *State
	bossTurn bool

	player struct {
		hp     int
		armor  int
		mana   int
		spells map[string]int
		spent  int
	}

	boss struct {
		hp     int
		damage int
	}
}

func (s *State) Copy() *State {
	state := &State{
		parent:   s,
		bossTurn: s.bossTurn,

		player: struct {
			hp     int
			armor  int
			mana   int
			spells map[string]int
			spent  int
		}{
			hp:     s.player.hp,
			armor:  s.player.armor,
			mana:   s.player.mana,
			spells: make(map[string]int),
			spent:  s.player.spent,
		},

		boss: struct {
			hp     int
			damage int
		}{
			hp:     s.boss.hp,
			damage: s.boss.damage,
		},
	}

	for spell, duration := range s.player.spells {
		state.player.spells[spell] = duration
	}

	return state
}

func (s *State) ID() string {
	parent := s.parent
	s.parent = nil
	id := fmt.Sprintf("%+v", *s)
	s.parent = parent
	return id
}

func (s *State) Children() []aoc.Node {
	// If either of us have died then we're done and there are no children.
	if s.player.hp <= 0 || s.boss.hp <= 0 {
		return nil
	}

	// Apply any active effects
	if s.player.spells["Shield"] > 0 {
		s.player.spells["Shield"]--
		if s.player.spells["Shield"] == 0 {
			s.player.armor -= 7
		}
	}

	if s.player.spells["Poison"] > 0 {
		s.player.spells["Poison"]--
		s.boss.hp -= 3
		if s.boss.hp <= 0 {
			return []aoc.Node{s}
		}
	}

	if s.player.spells["Recharge"] > 0 {
		s.player.spells["Recharge"]--
		s.player.mana += 101
	}

	if s.bossTurn {
		damage := s.boss.damage - s.player.armor
		if damage <= 0 {
			damage = 1
		}

		child := s.Copy()
		child.bossTurn = false
		child.player.hp -= damage
		return []aoc.Node{child}
	}

	// It's the player's turn.  Evaluate each option.
	var children []aoc.Node

	if s.player.mana >= 53 {
		child := s.Copy()
		child.bossTurn = true
		child.player.mana -= 53
		child.player.spent += 53
		child.boss.hp -= 4
		children = append(children, child)
	}

	if s.player.mana >= 73 {
		child := s.Copy()
		child.bossTurn = true
		child.player.mana -= 73
		child.player.spent += 73
		child.player.hp += 2
		child.boss.hp -= 2
		children = append(children, child)
	}

	if s.player.mana >= 113 && s.player.spells["Shield"] == 0 {
		child := s.Copy()
		child.bossTurn = true
		child.player.mana -= 113
		child.player.spent += 113
		child.player.spells["Shield"] = 6
		child.player.armor += 7
		children = append(children, child)
	}

	if s.player.mana >= 173 && s.player.spells["Poison"] == 0 {
		child := s.Copy()
		child.bossTurn = true
		child.player.mana -= 173
		child.player.spent += 173
		child.player.spells["Poison"] = 6
		children = append(children, child)
	}

	if s.player.mana >= 229 && s.player.spells["Recharge"] == 0 {
		child := s.Copy()
		child.bossTurn = true
		child.player.mana -= 229
		child.player.spent += 229
		child.player.spells["Recharge"] = 5
		children = append(children, child)
	}

	return children
}

func main() {
	state := InputToState(2015, 22)

	var best *State
	aoc.BreadthFirstSearch(state, func(node aoc.Node) bool {
		state := node.(*State)

		// We died, this isn't a goal state.
		if state.player.hp <= 0 {
			return false
		}

		if state.boss.hp <= 0 && (best == nil || state.player.spent < best.player.spent) {
			best = state
		}

		// Keep searching since we might find another goal that uses less mana.
		return false
	})

	DisplayState(best)

	fmt.Printf("spent: %d\n", best.player.spent)
}

func DisplayState(state *State) {
	if state == nil {
		return
	}

	DisplayState(state.parent)
	fmt.Println()
	fmt.Println("===================================")
	fmt.Println()

	if !state.bossTurn {
		fmt.Println("-- Player turn --")
	} else {
		fmt.Println("-- Boss turn --")
	}

	fmt.Printf("- Player has %d hit points, %d armor, %d mana\n", state.player.hp, state.player.armor, state.player.mana)
	fmt.Printf("- Boss has %d hit points\n", state.boss.hp)
	fmt.Printf("- Player spells: %+v\n", state.player.spells)

	if state.bossTurn && state.boss.hp > 0 {
		damage := state.boss.damage - state.player.armor
		if damage <= 0 {
			damage = 1
		}
		fmt.Printf("Boss attacks for %d damage.\n", damage)
	}
}

func InputToState(year, day int) *State {
	var hp, damage int
	for _, line := range aoc.InputToLines(year, day) {
		if _, err := fmt.Sscanf(line, "Hit Points: %d", &hp); err == nil {
			continue
		}

		if _, err := fmt.Sscanf(line, "ComputeDamage: %d", &damage); err == nil {
			continue
		}
	}

	return &State{
		player: struct {
			hp     int
			armor  int
			mana   int
			spells map[string]int
			spent  int
		}{
			hp:     50,
			mana:   500,
			spells: make(map[string]int),
		},

		boss: struct {
			hp     int
			damage int
		}{
			hp:     hp,
			damage: damage,
		},
	}
}
