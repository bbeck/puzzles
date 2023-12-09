package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	var sum int
	for _, game := range InputToGames() {
		if game.IsOk() {
			sum += game.ID
		}
	}

	fmt.Println(sum)
}

type Game struct {
	ID      int
	Subsets []map[string]int
}

func (g *Game) IsOk() bool {
	for _, s := range g.Subsets {
		for color, num := range s {
			switch {
			case color == "red" && num > 12:
				return false
			case color == "green" && num > 13:
				return false
			case color == "blue" && num > 14:
				return false
			}
		}
	}
	return true
}

func InputToGames() []Game {
	return aoc.InputLinesTo(2023, 2, func(line string) Game {
		line = strings.ReplaceAll(line, "Game ", "")
		line = strings.ReplaceAll(line, ":", "")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "; ", " ; ")
		fields := strings.Fields(line)

		subsets := []map[string]int{{}}
		for i := 1; i < len(fields); i++ {
			if fields[i] == ";" {
				subsets = append(subsets, make(map[string]int))
				continue
			}

			subsets[len(subsets)-1][fields[i+1]] = aoc.ParseInt(fields[i])
			i++
		}

		return Game{
			ID:      aoc.ParseInt(fields[0]),
			Subsets: subsets,
		}
	})
}
