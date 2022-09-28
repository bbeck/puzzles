package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var best int
	sort.Search(1<<10, func(boost int) bool {
		if result := Run(boost); result > 0 {
			best = result
			return true
		}
		return false
	})

	fmt.Println(best)
}

func Run(boost int) int {
	groups := InputToGroups()

	// Apply the boost
	for _, group := range groups {
		if group.Kind == "immune" {
			group.AttackPower += boost
		}
	}

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

		var used aoc.Set[*Group]
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
				defender.Count -= aoc.Min(defender.Count, Damage(attacker, defender)/defender.HitPoints)
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

func ChooseTarget(attacker *Group, groups []*Group, used aoc.Set[*Group]) *Group {
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
	ID          int
	Kind        string
	Count       int
	HitPoints   int
	AttackPower int
	AttackType  string
	Initiative  int
	Immunities  aoc.Set[string]
	Weaknesses  aoc.Set[string]
}

var regex = regexp.MustCompile(strings.TrimSpace(strings.Join([]string{
	`(?P<count>\d+) units`,
	`each with (?P<hp>\d+) hit points`,
	`(?:\((?P<modifiers>.*)\))?`,
	`with an attack that does (?P<ap>\d+) (?P<at>\S+) damage`,
	`at initiative (?P<initiative>\d+)`,
}, "\\s?")))

func InputToGroups() []*Group {
	var groups []*Group

	var kind string
	for _, line := range aoc.InputToLines(2018, 24) {
		if len(line) == 0 {
			continue
		}

		if line == "Immune System:" {
			kind = "immune"
			continue
		}

		if line == "Infection:" {
			kind = "infection"
			continue
		}

		fields := MatchFields(line, regex)
		groups = append(groups, &Group{
			Kind:        kind,
			Count:       aoc.ParseInt(fields["count"]),
			HitPoints:   aoc.ParseInt(fields["hp"]),
			AttackPower: aoc.ParseInt(fields["ap"]),
			AttackType:  fields["at"],
			Initiative:  aoc.ParseInt(fields["initiative"]),
			Immunities:  ParseModifiers(fields["modifiers"], "immune"),
			Weaknesses:  ParseModifiers(fields["modifiers"], "weak"),
		})
	}

	return groups
}

func MatchFields(s string, regex *regexp.Regexp) map[string]string {
	fields := make(map[string]string)

	names := regex.SubexpNames()
	matches := regex.FindStringSubmatch(s)
	for i := 1; i < len(names); i++ {
		fields[names[i]] = matches[i]
	}
	return fields
}

func ParseModifiers(s string, kind string) aoc.Set[string] {
	s = strings.ReplaceAll(s, " to ", " ")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ";", "")

	var modifiers aoc.Set[string]
	var save bool
	for _, field := range strings.Fields(s) {
		if field == kind {
			save = true
			continue
		} else if field == "immune" || field == "weak" {
			save = false
			continue
		} else if save {
			modifiers.Add(field)
		}
	}

	return modifiers
}
