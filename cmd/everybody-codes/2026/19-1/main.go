package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid1, grid2 := InputToGrids()

	var fc FrequencyCounter[string]
	grid1.ForEach(func(x int, y int, c1 string) {
		fc.Add(mixtures[c1+grid2.Get(x, y)])
	})

	var counts = fc.Entries()
	fmt.Println(counts[0].Count * counts[1].Count * counts[2].Count)
}

var mixtures = map[string]string{
	"CC": "C", "CM": "B", "CY": "G", "CK": "K",
	"MC": "B", "MM": "M", "MY": "R", "MK": "K",
	"YC": "G", "YM": "R", "YY": "Y", "YK": "K",
	"KC": "K", "KM": "K", "KY": "K", "KK": "K",
}

func InputToGrids() (Grid2D[string], Grid2D[string]) {
	chunk := in.ChunkS[string]()
	return chunk.StringGrid2D(), in.ToStringGrid2D()
}
