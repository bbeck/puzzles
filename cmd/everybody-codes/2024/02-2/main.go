package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"strings"
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
	lines := InputToLines()
	_, words, _ := strings.Cut(lines[0], ":")
	return SetFrom(strings.Split(words, ",")...), lines[2:]
}
