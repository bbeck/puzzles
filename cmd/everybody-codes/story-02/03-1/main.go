package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	dice := InputToDice()

	var roll, total int
	for total < 10000 {
		roll++

		for i := range dice {
			total += dice[i].Roll()
		}
	}
	fmt.Println(roll)
}

type Die struct {
	Faces      []int
	Seed       int
	Pulse      int
	RollNumber int
	Last       int
}

func (d *Die) Roll() int {
	spin := d.RollNumber * d.Pulse
	d.Last = (d.Last + spin) % len(d.Faces)
	d.Pulse = (d.Pulse+spin)%d.Seed + 1 + d.RollNumber + d.Seed
	d.RollNumber = d.RollNumber + 1
	return d.Faces[d.Last]
}

func InputToDice() []Die {
	return in.LinesToS(func(in in.Scanner[Die]) Die {
		splits := in.SplitS(" ")
		faces := splits[1].Ints()
		seed := splits[2].Int()
		return Die{Faces: faces, Seed: seed, Pulse: seed, RollNumber: 1}
	})
}
