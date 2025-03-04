package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
	"math"
)

func main() {
	rules := make(map[uint64]Grid)
	for _, rule := range InputToRules() {
		for _, from := range rule.from.Transformations() {
			rules[from.ID()] = rule.to
		}
	}

	g := Grid{Grid2D: Grid2D[bool]{
		Cells:  []bool{false, true, false, false, false, true, true, true, true},
		Width:  3,
		Height: 3,
	}}
	for n := 0; n < 5; n++ {
		g = Enhance(g, rules)
	}

	var count int
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.Get(x, y) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func Enhance(g Grid, rules map[uint64]Grid) Grid {
	N := 2
	if g.Width%2 != 0 {
		N = 3
	}

	next := NewGrid2D[bool](g.Width/N*(N+1), g.Height/N*(N+1))
	for cy := 0; cy < g.NumChunks(); cy++ {
		for cx := 0; cx < g.NumChunks(); cx++ {
			// Compute the top left coordinate of this chunk in the next grid
			x0 := cx * (N + 1)
			y0 := cy * (N + 1)

			// Determine the next chunk and fill it into the next grid
			c := rules[g.GetChunk(cx, cy).ID()]
			for dy := 0; dy < c.Height; dy++ {
				for dx := 0; dx < c.Width; dx++ {
					next.Set(x0+dx, y0+dy, c.Get(dx, dy))
				}
			}
		}
	}

	return Grid{next}
}

type Grid struct {
	Grid2D[bool]
}

func (g Grid) ID() uint64 {
	id := uint64(g.Width)
	g.ForEach(func(x, y int, value bool) {
		id <<= 1
		if value {
			id |= 1
		}
	})
	return id
}

func (g Grid) NumChunks() int {
	N := 2
	if g.Width%2 != 0 {
		N = 3
	}

	return g.Width / N
}

func (g Grid) GetChunk(cx, cy int) Grid {
	N := 2
	if g.Width%2 != 0 {
		N = 3
	}

	chunk := NewGrid2D[bool](N, N)
	for dy := 0; dy < N; dy++ {
		for dx := 0; dx < N; dx++ {
			chunk.Set(dx, dy, g.Get(cx*N+dx, cy*N+dy))
		}
	}

	return Grid{Grid2D: chunk}
}

func (g Grid) Transformations() []Grid {
	rotate := func(g Grid) Grid {
		return Grid{g.RotateRight()}
	}
	flipH := func(g Grid) Grid {
		r := NewGrid2D[bool](g.Width, g.Height)
		for y := 0; y < g.Height; y++ {
			for x := 0; x < g.Width; x++ {
				r.Set(x, y, g.Get(g.Width-x-1, y))
			}
		}
		return Grid{r}
	}
	flipV := func(g Grid) Grid {
		r := NewGrid2D[bool](g.Width, g.Height)
		for y := 0; y < g.Height; y++ {
			for x := 0; x < g.Width; x++ {
				r.Set(x, y, g.Get(x, g.Height-y-1))
			}
		}
		return Grid{r}
	}

	return []Grid{
		g,
		flipH(g),
		flipV(g),
		flipH(flipV(g)),
		rotate(g),
		flipH(rotate(g)),
		flipV(rotate(g)),
		flipV(flipH(rotate(g))),
	}
}

type Rule struct {
	from, to Grid
}

func InputToRules() []Rule {
	in.Remove("/")

	grid := func(s string) Grid {
		n := int(math.Sqrt(float64(len(s))))

		var cells []bool
		for _, ch := range s {
			cells = append(cells, ch == '#')
		}

		return Grid{Grid2D: Grid2D[bool]{Cells: cells, Height: n, Width: n}}
	}

	return in.LinesToS(func(in in.Scanner[Rule]) Rule {
		var lhs, rhs = in.Cut(" => ")
		return Rule{from: grid(lhs), to: grid(rhs)}
	})
}
