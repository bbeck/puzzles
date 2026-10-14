package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	dice, grid := InputToDiceAndGrid()

	var collected Set[Point2D]
	for i := range dice {
		var current Set[Point2D]
		grid.ForEachPoint(func(p Point2D, _ int) { current.Add(p) })

		for len(current) > 0 {
			value := dice[i].Roll()

			var next Set[Point2D]
			for p := range current {
				if !grid.InBoundsPoint(p) || grid.GetPoint(p) != value {
					continue
				}

				collected.Add(p)
				next.Add(p, p.Up(), p.Right(), p.Down(), p.Left())
			}
			current = next
		}
	}
	fmt.Println(len(collected))
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

func InputToDiceAndGrid() ([]Die, Grid2D[int]) {
	chunk := in.ChunkS[Die]()
	dice := chunk.LinesToS(func(in in.Scanner[Die]) Die {
		fields := in.FieldsS()
		faces := fields[1].Ints()
		seed := fields[2].Int()
		return Die{Faces: faces, Seed: seed, Pulse: seed, RollNumber: 1}
	})

	grid := in.ToIntGrid2D()
	return dice, grid
}
