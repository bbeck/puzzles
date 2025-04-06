package main

import (
	"fmt"
	"sort"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	groups := InputToGroups()

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
		var hasImmune, hasInfection bool
		for _, group := range groups {
			if group.Kind == "immune" && group.Count > 0 {
				hasImmune = true
			}
			if group.Kind == "infection" && group.Count > 0 {
				hasInfection = true
			}
		}

		if !hasImmune || !hasInfection {
			break
		}
	}

	var sum int
	for _, group := range groups {
		sum += group.Count
	}
	fmt.Println(sum)
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
		if dd < dc || (dd == dc && pd < pc) || (dd == dc && pd == pc && id < ic) {
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
