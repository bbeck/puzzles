package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	boss := InputToBoss()

	var best int
	EnumeratePlayers(func(player Character, cost int) {
		if !PlayerWins(player, boss) {
			best = puz.Max(best, cost)
		}
	})

	fmt.Println(best)
}

var Weapons = []struct {
	Name   string
	Cost   int
	Damage int
}{
	{"Dagger", 8, 4},
	{"Shortsword", 10, 5},
	{"Warhammer", 25, 6},
	{"Longsword", 40, 7},
	{"Greataxe", 74, 8},
}

var Armors = []struct {
	Name  string
	Cost  int
	Armor int
}{
	{"None", 0, 0},
	{"Leather", 13, 1},
	{"Chainmail", 31, 2},
	{"Splintmail", 53, 3},
	{"Bandedmail", 75, 4},
	{"Platemail", 102, 5},
}

var Rings = []struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}{
	{"None", 0, 0, 0},
	{"Damage +1", 25, 1, 0},
	{"Damage +2", 50, 2, 0},
	{"Damage +3", 100, 3, 0},
	{"Defense +1", 20, 0, 1},
	{"Defense +2", 40, 0, 2},
	{"Defense +3", 80, 0, 3},
}

func EnumeratePlayers(fn func(Character, int)) {
	for _, weapon := range Weapons {
		for _, armor := range Armors {
			for _, ring1 := range Rings {
				for _, ring2 := range Rings {
					if ring1 == ring2 && ring1.Name != "None" {
						// Not allowed to buy two of the same ring
						continue
					}

					player := Character{
						HitPoints: 100,
						Damage:    weapon.Damage + ring1.Damage + ring2.Damage,
						Armor:     armor.Armor + ring1.Armor + ring2.Armor,
					}
					cost := weapon.Cost + armor.Cost + ring1.Cost + ring2.Cost
					fn(player, cost)
				}
			}
		}
	}
}

func PlayerWins(player Character, boss Character) bool {
	for {
		// player turn first
		boss.HitPoints -= puz.Max(player.Damage-boss.Armor, 1)
		if boss.HitPoints <= 0 {
			return true
		}

		player.HitPoints -= puz.Max(boss.Damage-player.Armor, 1)
		if player.HitPoints <= 0 {
			return false
		}
	}
}

type Character struct {
	HitPoints int
	Damage    int
	Armor     int
}

func InputToBoss() Character {
	var boss Character
	for _, line := range puz.InputToLines() {
		fmt.Sscanf(line, "Hit Points: %d", &boss.HitPoints)
		fmt.Sscanf(line, "Damage: %d", &boss.Damage)
		fmt.Sscanf(line, "Armor: %d", &boss.Armor)
	}

	return boss
}
