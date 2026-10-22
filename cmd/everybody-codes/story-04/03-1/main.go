package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	width, height, horizontal, vertical := InputToPattern()

	hasHorizontal := func(x, y int) bool {
		ch := horizontal[y%len(horizontal)]
		return (ch == '0' && x%2 == 0) || (ch == '1' && x%2 == 1)
	}
	hasVertical := func(x, y int) bool {
		ch := vertical[x%len(vertical)]
		return (ch == '0' && y%2 == 0) || (ch == '1' && y%2 == 1)
	}

	var count int
	for y := range height {
		for x := range width {
			if hasHorizontal(x, y) && hasHorizontal(x, y+1) && hasVertical(x, y) && hasVertical(x+1, y) {
				count++
			}
		}
	}
	fmt.Println(count)
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
