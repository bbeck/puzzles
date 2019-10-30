package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	groups := InputToGroups(2018, 24)

	for round := 1; !Done(groups); round++ {
		fmt.Printf("=== Round %d ===\n", round)

		//
		// Status phase
		//
		sort.Slice(groups, func(i, j int) bool {
			return groups[i].id < groups[j].id
		})

		var immunes []*Group
		var infections []*Group
		for _, group := range groups {
			if group.kind == "immune" {
				immunes = append(immunes, group)
			} else {
				infections = append(infections, group)
			}
		}

		fmt.Println("  == Immune System ==")
		for _, group := range immunes {
			fmt.Printf("    Group %d contains %d units\n", group.id, group.count)
		}
		fmt.Println()

		fmt.Println("  == Infection ==")
		for _, group := range infections {
			fmt.Printf("    Group %d contains %d units\n", group.id, group.count)
		}
		fmt.Println()

		//
		// Target selection phase
		//
		fmt.Println("  == Target Selection ==")

		// sort by effective power, largest to smallest
		sort.Slice(groups, func(i, j int) bool {
			ip := groups[i].count * groups[i].ap
			jp := groups[j].count * groups[j].ap

			return ip > jp || (ip == jp && groups[i].initiative > groups[j].initiative)
		})

		chosen := make(map[*Group]bool)
		assignments := make(map[*Group]*Group)
		for _, group := range groups {
			var target *Group
			var targetDamage int

			for _, potentialTarget := range groups {
				if potentialTarget.kind == group.kind {
					continue
				}

				if chosen[potentialTarget] {
					continue
				}

				if potentialTarget.count == 0 {
					continue
				}

				// In the case of a tie, the defending group with the largest effective
				// power wins, if there's still a tie then it goes to the group with the
				// highest initiative
				damage := group.ComputeDamage(potentialTarget)
				fmt.Printf("    %s group %d would deal defending group %d %d damage\n",
					group.kind, group.id, potentialTarget.id, damage)
				if damage == 0 {
					continue
				}

				if target == nil {
					target = potentialTarget
					targetDamage = damage
					continue
				}

				if damage > targetDamage {
					target = potentialTarget
					targetDamage = damage
					continue
				}

				tEffectivePower := target.count * target.ap
				ptEffectivePower := potentialTarget.count * potentialTarget.ap
				tInitiative := target.initiative
				ptInitiative := potentialTarget.initiative
				if tEffectivePower < ptEffectivePower || (tEffectivePower == ptEffectivePower && tInitiative < ptInitiative) {
					target = potentialTarget
					targetDamage = damage
					continue
				}
			}

			if target != nil {
				// fmt.Printf("  %s group %d has chosen defending group %d\n",
				// 	group.kind, group.id, target.id)
				chosen[target] = true
				assignments[group] = target
			}
		}

		fmt.Println()

		// Attack phase
		sort.Slice(groups, func(i, j int) bool {
			return groups[i].initiative > groups[j].initiative
		})

		fmt.Println("  == Attack Phase ==")
		for _, group := range groups {
			if group.count == 0 {
				continue
			}

			target := assignments[group]
			if target == nil {
				continue
			}

			damage := group.ComputeDamage(target)
			if damage == 0 {
				continue
			}

			delta := damage / target.hp
			fmt.Printf("    %s group %d attacks defending group %d, killing %d units\n",
				group.kind, group.id, target.id, delta)

			target.count -= delta
			if target.count < 0 {
				target.count = 0
			}
		}

		fmt.Println()
	}

	var winner string
	var winnerCount int
	for _, group := range groups {
		if group.count > 0 {
			winner = group.kind
			winnerCount += group.count
		}
	}

	fmt.Printf("%s wins with %d units remaining\n", winner, winnerCount)
}

func Done(groups []*Group) bool {
	counts := make(map[string]int)
	for _, group := range groups {
		counts[group.kind] += group.count
	}

	for _, count := range counts {
		if count == 0 {
			return true
		}
	}

	return false
}

func Contains(needle string, haystack []string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}

	return false
}

type Group struct {
	kind       string
	id         int
	count      int
	hp         int
	immunities []string
	weaknesses []string
	ap         int
	at         string
	initiative int
}

func (g *Group) ComputeDamage(target *Group) int {
	if Contains(g.at, target.immunities) {
		return 0
	}

	if Contains(g.at, target.weaknesses) {
		return 2 * g.count * g.ap
	}

	return g.count * g.ap
}

func InputToGroups(year, day int) []*Group {
	var kind string
	var id int
	var groups []*Group
	for _, line := range aoc.InputToLines(year, day) {
		if len(line) == 0 {
			continue
		}

		if line == "Immune System:" {
			kind = "immune"
			id = 1
			continue
		}

		if line == "Infection:" {
			kind = "infection"
			id = 1
			continue
		}

		fields := MatchGroup(line)
		groups = append(groups, &Group{
			kind:       kind,
			id:         id,
			count:      aoc.ParseInt(fields["count"]),
			hp:         aoc.ParseInt(fields["hp"]),
			immunities: Immunities(fields["modifiers"]),
			weaknesses: Weaknesses(fields["modifiers"]),
			ap:         aoc.ParseInt(fields["ap"]),
			at:         fields["at"],
			initiative: aoc.ParseInt(fields["initiative"]),
		})
		id++
	}

	return groups
}

var groupRegex = regexp.MustCompile(`(?P<count>\d+) units each with (?P<hp>\d+) hit points(?: \((?P<modifiers>.*)\))? with an attack that does (?P<ap>\d+) (?P<at>.*) damage at initiative (?P<initiative>\d+)`)

func MatchGroup(s string) map[string]string {
	fields := make(map[string]string)

	names := groupRegex.SubexpNames()
	parts := groupRegex.FindStringSubmatch(s)
	for i := 1; i < len(names); i++ {
		fields[names[i]] = parts[i]
	}

	return fields
}

func Immunities(s string) []string {
	var immunities []string
	for _, part := range strings.Split(s, "; ") {
		if strings.HasPrefix(part, "immune to ") {
			for _, immunity := range strings.Split(part[10:], ", ") {
				immunities = append(immunities, immunity)
			}
		}
	}
	return immunities
}

func Weaknesses(s string) []string {
	var weaknesses []string
	for _, part := range strings.Split(s, "; ") {
		if strings.HasPrefix(part, "weak to ") {
			for _, weakness := range strings.Split(part[8:], ", ") {
				weaknesses = append(weaknesses, weakness)
			}
		}
	}
	return weaknesses
}
