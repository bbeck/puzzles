package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	var sum int
	for _, game := range InputToGames() {
		sum += game.Power()
	}

	fmt.Println(sum)
}

type Game struct {
	ID      int
	Subsets []map[string]int
}

func (g *Game) Power() int {
	var maxRed, maxBlue, maxGreen int
	for _, s := range g.Subsets {
		maxRed = puz.Max(maxRed, s["red"])
		maxGreen = puz.Max(maxGreen, s["green"])
		maxBlue = puz.Max(maxBlue, s["blue"])
	}

	return maxRed * maxBlue * maxGreen
}

func InputToGames() []Game {
	return puz.InputLinesTo(2023, 2, func(line string) Game {
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

			subsets[len(subsets)-1][fields[i+1]] = puz.ParseInt(fields[i])
			i++
		}

		return Game{
			ID:      puz.ParseInt(fields[0]),
			Subsets: subsets,
		}
	})
}
