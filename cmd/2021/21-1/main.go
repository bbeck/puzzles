package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	positions := InputToStartingPositions()
	scores := []int{0, 0}

	var die = Die{value: 1}
	var rolls int
	for player := 0; ; player = (player + 1) % 2 {
		position := positions[player]
		distance := die.Roll() + die.Roll() + die.Roll()
		rolls += 3

		position += distance
		for position > 10 {
			position -= 10
		}
		scores[player] += position

		positions[player] = position
		if scores[player] >= 1000 {
			break
		}
	}
	fmt.Println(aoc.MinInt(scores[0], scores[1]) * rolls)
}

type Die struct {
	value int
}

func (d *Die) Roll() int {
	value := d.value
	d.value++
	if d.value > 100 {
		d.value = 1
	}
	return value
}

func InputToStartingPositions() []int {
	lines := aoc.InputToLines(2021, 21)

	return []int{
		aoc.ParseInt(strings.Split(lines[0], ": ")[1]),
		aoc.ParseInt(strings.Split(lines[1], ": ")[1]),
	}
}
