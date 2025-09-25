package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	words, lines := InputToMessage()

	// We need to search in both directions so include the reverse words so we
	// only need to search from left to right.
	for _, word := range words.Entries() {
		reversed := string(Reversed([]byte(word)))
		words.Add(reversed)
	}

	var seen Set[Point2D] // {X: index within line, Y: line number}
	for _, word := range words.Entries() {
		for y, line := range lines {
			var start int
			for {
				index := strings.Index(line[start:], word)
				if index == -1 {
					break
				}

				for dx := 0; dx < len(word); dx++ {
					seen.Add(Point2D{X: start + index + dx, Y: y})
				}
				start += index + 1
			}
		}
	}
	fmt.Println(len(seen))
}

func InputToMessage() (Set[string], []string) {
	_, words := in.Cut(":")
	in.Line() // blank line
	return SetFrom(strings.Split(words, ",")...), in.Lines()
}
