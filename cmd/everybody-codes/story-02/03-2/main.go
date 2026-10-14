package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	dice, track := InputToDiceAndTrack()
	var D, T = len(dice), len(track)

	var turns = make([]int, D)
	for i := range dice {
		var pos int
		for pos < T {
			turns[i]++
			if dice[i].Roll() == track[pos] {
				pos++
			}
		}
	}

	var indices = make([]int, D)
	for i := range indices {
		indices[i] = i + 1
	}
	slices.SortFunc(indices, func(i, j int) int {
		return cmp.Compare(turns[i-1], turns[j-1])
	})

	var order = make([]string, D)
	for i := range indices {
		order[i] = strconv.Itoa(indices[i])
	}

	fmt.Println(strings.Join(order, ","))
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

func InputToDiceAndTrack() ([]Die, []int) {
	chunk := in.ChunkS[Die]()
	dice := chunk.LinesToS(func(in in.Scanner[Die]) Die {
		splits := in.SplitS(" ")
		faces := splits[1].Ints()
		seed := splits[2].Int()
		return Die{Faces: faces, Seed: seed, Pulse: seed, RollNumber: 1}
	})

	var track []int
	for in.HasNext() {
		track = append(track, int(in.Byte()-'0'))
	}

	return dice, track
}
