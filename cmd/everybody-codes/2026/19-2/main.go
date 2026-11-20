package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid, nums := InputToGrids()

	var best int
	EnumeratePermutations(len(colors), func(permutation []int) bool {
		var fc FrequencyCounter[string]
		grid.ForEach(func(x int, y int, c1 string) {
			fc.Add(mixtures[c1+colors[permutation[nums.Get(x, y)-1]]])
		})

		var counts = fc.Entries()
		best = Max(best, counts[0].Count*counts[1].Count*counts[2].Count)
		return false
	})
	fmt.Println(best)
}

var colors = []string{"C", "M", "Y", "K"}

var mixtures = map[string]string{
	"CC": "C", "CM": "B", "CY": "G", "CK": "K",
	"MC": "B", "MM": "M", "MY": "R", "MK": "K",
	"YC": "G", "YM": "R", "YY": "Y", "YK": "K",
	"KC": "K", "KM": "K", "KY": "K", "KK": "K",
}

func InputToGrids() (Grid2D[string], Grid2D[int]) {
	chunk := in.ChunkS[string]()
	return chunk.StringGrid2D(), in.ToIntGrid2D()
}
