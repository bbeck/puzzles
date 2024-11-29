package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"slices"
)

func main() {
	grid := InputToGrid2D(func(_ int, _ int, s string) string {
		return s
	})

	for {
		if !Infer(grid) {
			break
		}
	}

	var sum int
	for y := 0; y < grid.Height-2; y += 6 {
		for x := 0; x < grid.Width-2; x += 6 {
			word := Word(grid, x, y)
			if !slices.Contains(word, ".") {
				sum += Power(word)
			}
		}
	}
	fmt.Println(sum)
}

func Infer(grid Grid2D[string]) bool {
	var changed bool
	grid.ForEach(func(x int, y int, s string) {
		if s == "." {
			bx, by := 6*(x/6), 6*(y/6)
			rs := Row(grid, bx, y)
			cs := Col(grid, x, by)
			rh := SetFrom(rs[0], rs[1], rs[6], rs[7])
			ch := SetFrom(cs[0], cs[1], cs[6], cs[7])
			rc := SetFrom(rs[2], rs[3], rs[4], rs[5])
			cc := SetFrom(cs[2], cs[3], cs[4], cs[5])

			// If the row and column headers share a character then this cell must
			// be that character since we can't repeat characters within a word.
			if i := rh.Intersect(ch); len(i) == 1 && !i.Contains("?") {
				grid.Set(x, y, i.Entries()[0])
				changed = true
			}

			// Each header character must appear in the row.
			if d := rh.Difference(rc); len(d) == 1 && !d.Contains("?") {
				grid.Set(x, y, d.Entries()[0])
				changed = true
			}

			// Each header character must appear in the column.
			if d := ch.Difference(cc); len(d) == 1 && !d.Contains("?") {
				grid.Set(x, y, d.Entries()[0])
				changed = true
			}
		}

		if s == "?" {
			isColHeader :=
				(grid.InBounds(x-1, y) && grid.Get(x-1, y) == "*") ||
					(grid.InBounds(x-2, y) && grid.Get(x-2, y) == "*") ||
					(grid.InBounds(x-3, y) && grid.Get(x-3, y) == "*") ||
					(grid.InBounds(x-4, y) && grid.Get(x-4, y) == "*")
			isRowHeader := !isColHeader

			// If this is in a row header, and the row is completely filled out then
			// we can determine the value of this cell.
			if isRowHeader {
				for dx := -2; dx <= 2; dx += 4 {
					if !grid.InBounds(x+dx, y) {
						continue
					}

					bx := 6 * ((x + dx) / 6)
					rs := Row(grid, bx, y)
					rh := SetFrom(rs[0], rs[1], rs[6], rs[7])
					rc := SetFrom(rs[2], rs[3], rs[4], rs[5])

					if d := rc.Difference(rh); !d.Contains(".") && len(d) == 1 {
						grid.Set(x, y, d.Entries()[0])
						changed = true
					}
				}
			}

			if isColHeader {
				for dy := -2; dy <= 2; dy += 4 {
					if !grid.InBounds(x, y+dy) {
						continue
					}

					by := 6 * ((y + dy) / 6)
					cs := Col(grid, x, by)
					ch := SetFrom(cs[0], cs[1], cs[6], cs[7])
					cc := SetFrom(cs[2], cs[3], cs[4], cs[5])

					if d := cc.Difference(ch); !d.Contains(".") && len(d) == 1 {
						grid.Set(x, y, d.Entries()[0])
						changed = true
					}
				}
			}
		}
	})
	return changed
}

func Row(grid Grid2D[string], x, y int) []string {
	var row []string
	for dx := 0; dx < 8; dx++ {
		row = append(row, grid.Get(x+dx, y))
	}
	return row
}

func Col(grid Grid2D[string], x, y int) []string {
	var col []string
	for dy := 0; dy < 8; dy++ {
		col = append(col, grid.Get(x, y+dy))
	}
	return col
}

func Word(grid Grid2D[string], x, y int) []string {
	var word []string
	for dy := 2; dy < 6; dy++ {
		for dx := 2; dx < 6; dx++ {
			word = append(word, grid.Get(x+dx, y+dy))
		}
	}
	return word
}

func Power(word []string) int {
	powers := map[string]int{
		"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9,
		"J": 10, "K": 11, "L": 12, "M": 13, "N": 14, "O": 15, "P": 16, "Q": 17,
		"R": 18, "S": 19, "T": 20, "U": 21, "V": 22, "W": 23, "X": 24, "Y": 25,
		"Z": 26,
	}

	var power int
	for i, s := range word {
		power += (i + 1) * powers[s]
	}
	return power
}
