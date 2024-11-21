package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	words, lines := InputToMessage()
	cols := Cols(lines)

	// We need to search in both directions so include the reverse words so we
	// only need to search from left to right.
	for _, word := range words.Entries() {
		reversed := string(Reversed([]byte(word)))
		words.Add(reversed)
	}

	var seen Set[Point2D]
	for _, word := range words.Entries() {
		for y, line := range Rows(word, lines) {
			for _, index := range Indices(word, line, len(lines[y])) {
				seen.Add(Point2D{X: index, Y: y})
			}
		}

		for x, line := range cols {
			for _, index := range Indices(word, line, len(line)) {
				seen.Add(Point2D{X: x, Y: index})
			}
		}
	}

	fmt.Println(len(seen))
}

func Rows(word string, lines []string) []string {
	var rows []string
	for _, row := range lines {
		rows = append(rows, row+row[:len(word)-1])
	}
	return rows
}

func Cols(lines []string) []string {
	var cols []string
	for x := 0; x < len(lines[0]); x++ {
		var sb strings.Builder
		for y := 0; y < len(lines); y++ {
			sb.WriteByte(lines[y][x])
		}
		cols = append(cols, sb.String())
	}
	return cols
}

func Indices(word, line string, L int) []int {
	var indices []int

	var start int
	for {
		index := strings.Index(line[start:], word)
		if index == -1 {
			break
		}

		for delta := 0; delta < len(word); delta++ {
			// The line sometimes has extra characters at the end because we support
			// wrapping rows around.  So instead of using the length of the line use
			// the length passed in.
			indices = append(indices, (start+index+delta)%L)
		}

		start += index + 1
	}

	return indices
}

func InputToMessage() (Set[string], []string) {
	lines := InputToLines()
	_, words, _ := strings.Cut(lines[0], ":")
	return SetFrom(strings.Split(words, ",")...), lines[2:]
}
