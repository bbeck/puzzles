package main

import (
	"fmt"
	"sort"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	groups := InputToGroups()

	var best int
	sort.Search(1<<10, func(boost int) bool {
		// Apply the boost
		var updated []*Group
		for _, group := range groups {
			g := *group
			if g.Kind == "immune" {
				g.AttackPower += boost
			}
			updated = append(updated, &g)
		}

		if result := Run(updated); result > 0 {
			best = result
			return true
		}
		return false
	})

	fmt.Println(best)
}

func Run(groups []*Group) int {
	ByEffectivePower := func(i, j int) bool {
		gi, gj := groups[i], groups[j]
		if pi, pj := EffectivePower(gi), EffectivePower(gj); pi != pj {
			return pi > pj
		}
		return gi.Initiative > gj.Initiative
	}

	ByInitiative := func(i, j int) bool {
		return groups[i].Initiative > groups[j].Initiative
	}

	for {
		//
		// Target selection phase
		//
		sort.Slice(groups, ByEffectivePower)

		var used Set[*Group]
		targets := make(map[*Group]*Group)
		for _, attacker := range groups {
			defender := ChooseTarget(attacker, groups, used)
			if defender != nil {
				targets[attacker] = defender
				used.Add(defender)
			}
		}

		// If there are no attacks to be made then we've deadlocked.
		if len(targets) == 0 {
			break
		}

		//
		// Attack phase
		//
		sort.Slice(groups, ByInitiative)

		for _, attacker := range groups {
			defender := targets[attacker]
			if defender != nil {
				defender.Count -= Min(defender.Count, Damage(attacker, defender)/defender.HitPoints)
			}
		}

		//
		// Determination phase
		//
		counts := Count(groups)
		if counts["immune"] == 0 || counts["infection"] == 0 {
			break
		}
	}

	counts := Count(groups)
	if counts["infection"] > 0 {
		return 0
	}
	return counts["immune"]
}

func Count(groups []*Group) map[string]int {
	counts := make(map[string]int)
	for _, group := range groups {
		counts[group.Kind] += group.Count
	}
	return counts
}

func EffectivePower(g *Group) int {
	return g.Count * g.AttackPower
}

func Damage(attacker, defender *Group) int {
	multiplier := 1
	if defender.Immunities.Contains(attacker.AttackType) {
		multiplier = 0
	} else if defender.Weaknesses.Contains(attacker.AttackType) {
		multiplier = 2
	}

	return multiplier * attacker.Count * attacker.AttackPower
}

func ChooseTarget(attacker *Group, groups []*Group, used Set[*Group]) *Group {
	var defender *Group
	for _, candidate := range groups {
		if attacker.Kind == candidate.Kind || candidate.Count <= 0 || used.Contains(candidate) {
			continue
		}

		dc := Damage(attacker, candidate)
		if dc <= 0 {
			continue
		}

		if defender == nil {
			defender = candidate
			continue
		}

		dd := Damage(attacker, defender)
		pd, pc := EffectivePower(defender), EffectivePower(candidate)
		id, ic := defender.Initiative, candidate.Initiative
		if defender == nil || dd < dc || (dd == dc && pd < pc) || (dd == dc && pd == pc && id < ic) {
			defender = candidate
		}
	}

	return defender
}

type Group struct {
	Kind        string
	Count       int
	HitPoints   int
	AttackPower int
	AttackType  string
	Initiative  int
	Immunities  Set[string]
	Weaknesses  Set[string]
}

func InputToGroups() []*Group {
	var groups []*Group
	var kind string

	for in.HasNext() {
		switch {
		case in.HasPrefix("Immune System:"):
			kind = "immune"
			in.Line()

		case in.HasPrefix("Infection:"):
			kind = "infection"
			in.Line()

		case in.HasPrefix("\n"):
			in.Line()

		default:
			var modifiers in.Scanner[any]

			var group = Group{Kind: kind}
			in.Scanf(
				"%d units each with %d hit points"+
					"%s"+
					"with an attack that does %d %s damage "+
					"at initiative %d",
				&group.Count,
				&group.HitPoints,
				&modifiers,
				&group.AttackPower,
				&group.AttackType,
				&group.Initiative,
			)

			group.Immunities, group.Weaknesses = ParseModifiers(modifiers)
			groups = append(groups, &group)
		}
	}

	return groups
}

func ParseModifiers(in in.Scanner[any]) (Set[string], Set[string]) {
	in.Remove("(", ")", ";", ",")

	var immunities, weaknesses Set[string]
	var current *Set[string]
	for _, field := range in.Fields() {
		switch field {
		case "immune":
			current = &immunities
		case "weak":
			current = &weaknesses
		default:
			if current != nil {
				current.Add(field)
			}
		}
	}

	return immunities, weaknesses
}
