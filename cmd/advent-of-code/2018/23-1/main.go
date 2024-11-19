package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"sort"
)

func main() {
	bots := InputToNanobots()
	sort.Slice(bots, func(i, j int) bool {
		return bots[i].R > bots[j].R
	})

	var strongest = bots[0]
	var count int
	for _, b := range bots {
		if strongest.ManhattanDistance(b.Point3D) <= strongest.R {
			count++
		}
	}
	fmt.Println(count)
}

type Nanobot struct {
	lib.Point3D
	R int
}

func InputToNanobots() []Nanobot {
	return lib.InputLinesTo(func(line string) Nanobot {
		var bot Nanobot
		fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &bot.X, &bot.Y, &bot.Z, &bot.R)
		return bot
	})
}
