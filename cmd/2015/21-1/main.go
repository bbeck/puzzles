package main

import (
	"fmt"
	"math"

	"github.com/bbeck/advent-of-code/aoc"
)

type Character struct {
	hp     int
	damage int
	armor  int
}

type Item struct {
	name   string
	cost   int
	damage int
	armor  int
}

var weapons = []Item{
	{"Dagger", 8, 4, 0},
	{"Shortsword", 10, 5, 0},
	{"Warhammer", 25, 6, 0},
	{"Longsword", 40, 7, 0},
	{"Greataxe", 74, 8, 0},
}

var armors = []Item{
	{"None", 0, 0, 0},
	{"Leather", 13, 0, 1},
	{"Chainmail", 31, 0, 2},
	{"Splintmail", 53, 0, 3},
	{"Bandedmail", 75, 0, 4},
	{"Platemail", 102, 0, 5},
}

var rings = []Item{
	{"None", 0, 0, 0},
	{"ComputeDamage +1", 25, 1, 0},
	{"ComputeDamage +2", 50, 2, 0},
	{"ComputeDamage +3", 100, 3, 0},
	{"Defense +1", 20, 0, 1},
	{"Defense +2", 40, 0, 2},
	{"Defense +3", 80, 0, 3},
}

func main() {
	boss := InputToCharacter(2015, 21)
	fmt.Printf("boss: %+v\n", boss)

	best := math.MaxInt64
	EnumeratePlayers(func(player Character, cost int) {
		if cost < best && PlayerWins(player, boss) {
			best = cost
		}
	})

	fmt.Printf("best: %d\n", best)
}

func EnumeratePlayers(fn func(player Character, cost int)) {
	for _, weapon := range weapons {
		for _, armor := range armors {
			for _, ring1 := range rings {
				for _, ring2 := range rings {
					if ring1.name != "None" && ring1 == ring2 {
						continue
					}

					fn(Character{
						hp:     100,
						damage: weapon.damage + armor.damage + ring1.damage + ring2.damage,
						armor:  weapon.armor + armor.armor + ring1.armor + ring2.armor,
					}, weapon.cost+armor.cost+ring1.cost+ring2.cost)
				}
			}
		}
	}
}

func PlayerWins(player Character, boss Character) bool {
	for {
		// player turn first
		boss.hp -= player.damage - boss.armor
		if boss.hp <= 0 {
			return true
		}

		player.hp -= boss.damage - player.armor
		if player.hp <= 0 {
			return false
		}
	}
}

func InputToCharacter(year, day int) Character {
	var hp, damage, armor int
	for _, line := range aoc.InputToLines(year, day) {
		if _, err := fmt.Sscanf(line, "Hit Points: %d", &hp); err == nil {
			continue
		}

		if _, err := fmt.Sscanf(line, "Damage: %d", &damage); err == nil {
			continue
		}

		if _, err := fmt.Sscanf(line, "Armor: %d", &armor); err == nil {
			continue
		}
	}

	return Character{hp, damage, armor}
}
