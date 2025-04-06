package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
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
	Point3D
	R int
}

func InputToNanobots() []Nanobot {
	return in.LinesToS[Nanobot](func(s in.Scanner[Nanobot]) Nanobot {
		return Nanobot{
			Point3D: Point3D{X: in.Int(), Y: in.Int(), Z: in.Int()},
			R:       in.Int(),
		}
	})
}
