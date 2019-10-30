package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := Grid([][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	})
	rules := InputToRules(2017, 21)

	for iter := 0; iter < 5; iter++ {
		grid = grid.Enhance(rules)
	}

	var count int
	for _, row := range grid {
		for _, c := range row {
			if c {
				count++
			}
		}
	}

	fmt.Printf("count: %d\n", count)
}

type Grid [][]bool

func (g Grid) Print() {
	chars := map[bool]string{
		false: ".",
		true:  "#",
	}

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			fmt.Print(chars[g[y][x]])
		}
		fmt.Println()
	}
}

func (g Grid) Enhance(rules []Rule) Grid {
	// We have to subdivide the grid into either 2x2 or 3x3 chunks which
	// themselves are just small grids.
	N := len(g)

	var stride = 3
	if N%2 == 0 {
		stride = 2
	}

	var get func(x, y int) Grid
	if stride == 2 {
		get = func(x, y int) Grid {
			return [][]bool{
				{g[y+0][x+0], g[y+0][x+1]},
				{g[y+1][x+0], g[y+1][x+1]},
			}
		}
	} else {
		get = func(x, y int) Grid {
			return [][]bool{
				{g[y+0][x+0], g[y+0][x+1], g[y+0][x+2]},
				{g[y+1][x+0], g[y+1][x+1], g[y+1][x+2]},
				{g[y+2][x+0], g[y+2][x+1], g[y+2][x+2]},
			}
		}
	}

	// Navigate through the grid in chunks, building up a new grid by enhancing
	// one chunk at a time.  The new grid will be square with a side length of
	// (N / stride) * (stride + 1).
	var next Grid
	for y := 0; y < N; y += stride {
		for i := 0; i < stride+1; i++ {
			next = append(next, make([]bool, 0))
		}

		for x := 0; x < N; x += stride {
			// This is now either a 2x2 or 3x3 chunk of the grid that we need to
			// enhance.  Find the rule that matches it, and use it to build a new
			// chunk.
			chunk := get(x, y)

			var found bool
			for _, rule := range rules {
				if rule.Matches(chunk) {
					for dy := 0; dy < len(rule.to); dy++ {
						next[(y/stride)*(stride+1)+dy] = append(next[(y/stride)*(stride+1)+dy], rule.to[dy]...)
					}

					found = true
					break
				}
			}

			if !found {
				fmt.Println("couldn't find match for chunk:")
				chunk.Print()
				log.Fatal("failed")
			}
		}
	}

	return next
}

type Rule struct {
	from, to Grid
}

func (r Rule) Matches(g Grid) bool {
	// Make sure the chunk has the same dimensions as the rule.
	if len(r.from) != len(g) {
		return false
	}

	// We have to consider all possible combinations of rotations and flips of the
	// grid for a match.
	if len(r.from) == 2 {
		//        Flips
		//    AB  BA  DC  CD
		//    CD  DC  BA  AB
		// R
		// o  CA  DB  BD  AC
		// t  DB  CA  AC  BD
		// a
		// t  DC  CD  AB  BA
		// e  BA  AB  CD  DC
		// s
		//    BD  AC  CA  DB
		//    AC  BD  DB  CA
		a, b, c, d := g[0][0], g[0][1], g[1][0], g[1][1]
		A, B, C, D := r.from[0][0], r.from[0][1], r.from[1][0], r.from[1][1]
		return (A == a && B == b && C == c && D == d) ||
			(B == a && A == b && D == c && C == d) ||
			(D == a && C == b && B == c && A == d) ||
			(C == a && D == b && A == c && B == d) ||
			(C == a && A == b && D == c && B == d) ||
			(D == a && B == b && C == c && A == d) ||
			(B == a && D == b && A == c && C == d) ||
			(A == a && C == b && B == c && D == d) ||
			(D == a && C == b && B == c && A == d) ||
			(C == a && D == b && A == c && B == d) ||
			(A == a && B == b && C == c && D == d) ||
			(B == a && A == b && D == c && C == d) ||
			(B == a && D == b && A == c && C == d) ||
			(A == a && C == b && B == c && D == d) ||
			(C == a && A == b && D == c && B == d) ||
			(D == a && B == b && C == c && A == d)
	} else {
		//          Flips
		//    ABC  CBA  IHG  GHI
		//    DEF  FED  FED  DEF
		//    GHI  IHG  CBA  ABC
		//
		// R  GDA  IFC  CFI  ADG
		// o  HEB  HEB  BEH  BEH
		// t  IFC  GDA  ADG  CFI
		// a
		// t  IHG  GHI  ABC  CBA
		// e  FED  DEF  DEF  FED
		// s  CBA  ABC  GHI  IHG
		//
		//    CFI  ADG  GDA  IFC
		//    BEH  BEH  HEB  HEB
		//    ADG  CFI  IFC  GDA
		a, b, c := g[0][0], g[0][1], g[0][2]
		d, e, f := g[1][0], g[1][1], g[1][2]
		g, h, i := g[2][0], g[2][1], g[2][2]
		A, B, C := r.from[0][0], r.from[0][1], r.from[0][2]
		D, E, F := r.from[1][0], r.from[1][1], r.from[1][2]
		G, H, I := r.from[2][0], r.from[2][1], r.from[2][2]
		return (A == a && B == b && C == c && D == d && E == e && F == f && G == g && H == h && I == i) ||
			(C == a && B == b && A == c && F == d && E == e && D == f && I == g && H == h && G == i) ||
			(I == a && H == b && G == c && F == d && E == e && D == f && C == g && B == h && A == i) ||
			(G == a && H == b && I == c && D == d && E == e && F == f && A == g && B == h && C == i) ||
			(G == a && D == b && A == c && H == d && E == e && B == f && I == g && F == h && C == i) ||
			(I == a && F == b && C == c && H == d && E == e && B == f && G == g && D == h && A == i) ||
			(C == a && F == b && I == c && B == d && E == e && H == f && A == g && D == h && G == i) ||
			(A == a && D == b && G == c && B == d && E == e && H == f && C == g && F == h && I == i) ||
			(I == a && H == b && G == c && F == d && E == e && D == f && C == g && B == h && A == i) ||
			(G == a && H == b && I == c && D == d && E == e && F == f && A == g && B == h && C == i) ||
			(A == a && B == b && C == c && D == d && E == e && F == f && G == g && H == h && I == i) ||
			(C == a && B == b && A == c && F == d && E == e && D == f && I == g && H == h && G == i) ||
			(C == a && F == b && I == c && B == d && E == e && H == f && A == g && D == h && G == i) ||
			(A == a && D == b && G == c && B == d && E == e && H == f && C == g && F == h && I == i) ||
			(G == a && D == b && A == c && H == d && E == e && B == f && I == g && F == h && C == i) ||
			(I == a && F == b && C == c && H == d && E == e && B == f && G == g && D == h && A == i)
	}
}

func InputToRules(year, day int) []Rule {
	parse := func(s string) Grid {
		var grid [][]bool
		for y, row := range strings.Split(s, "/") {
			grid = append(grid, make([]bool, len(row)))

			for x, c := range row {
				if c == '#' {
					grid[y][x] = true
				}
			}
		}

		return grid
	}

	var rules []Rule
	for _, line := range aoc.InputToLines(year, day) {
		parts := strings.Split(line, " => ")
		lhs, rhs := parts[0], parts[1]

		rules = append(rules, Rule{from: parse(lhs), to: parse(rhs)})
	}

	return rules
}
