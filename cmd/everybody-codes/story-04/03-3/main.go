package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	width, height, horizontal, vertical := InputToPattern()
	tW, tH := 2*len(vertical), 2*len(horizontal)

	hasHorizontal := func(x, y int) bool {
		ch := horizontal[y%len(horizontal)]
		return (ch == '0' && x%2 == 0) || (ch == '1' && x%2 == 1)
	}
	hasVertical := func(x, y int) bool {
		ch := vertical[x%len(vertical)]
		return (ch == '0' && y%2 == 0) || (ch == '1' && y%2 == 1)
	}
	isIsolated := func(x, y int) bool {
		return hasHorizontal(x, y) && hasHorizontal(x, y+1) && hasVertical(x, y) && hasVertical(x+1, y)
	}

	// Compute the colors for just the base tile.
	var colors = make(map[Point2D]bool)
	for y := range tH {
		for x := range tW {
			p := Point2D{X: x, Y: y}

			if x > 0 && !hasVertical(x, y) {
				colors[p] = colors[p.Left()]
			} else if x > 0 {
				colors[p] = !colors[p.Left()]
			} else if y > 0 && !hasHorizontal(x, y) {
				colors[p] = colors[p.Up()]
			} else if y > 0 {
				colors[p] = !colors[p.Up()]
			}
		}
	}

	// Now we can determine our counts.  To count properly we have to consider
	// four different areas: full tiles, vertical strip, horizontal strip and
	// the corner.
	//
	// ┌───────┬───────┬─────┬───┐
	// │       │       │     │   │
	// │   T   │   T   │ ... │ V │
	// │       │       │     │   │
	// ├───────┼───────┼─────┼───┤
	// │  ...  │  ...  │ ... │...│
	// ├───────┼───────┼─────┼───┤
	// │   H   │   H   │ ... │ C │
	// └───────┴───────┴─────┴───┘
	count := func(w, h int, mult int, fc *FrequencyCounter[bool]) {
		for y := range h {
			for x := range w {
				if isIsolated(x, y) {
					fc.AddWithCount(colors[Point2D{X: x, Y: y}], mult)
				}
			}
		}
	}

	var counts FrequencyCounter[bool]
	count(tW, tH, (width/tW)*(height/tH), &counts) // T
	count(width%tW, tH, height/tH, &counts)        // V
	count(tW, height%tH, width/tW, &counts)        // H
	count(width%tW, height%tH, 1, &counts)         // C

	fmt.Println(counts.Entries()[0].Count)
}

func InputToPattern() (int, int, string, string) {
	var kv = make(map[string]string)
	in.LinesToS(func(in in.Scanner[any]) any {
		key, value := in.Cut("=")
		kv[key] = value
		return nil
	})
	return ParseInt(kv["width"]), ParseInt(kv["height"]), kv["horizontal-offsets"], kv["vertical-offsets"]
}
