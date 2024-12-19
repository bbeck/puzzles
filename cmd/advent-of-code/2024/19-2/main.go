package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	towels, designs := InputToTowelsAndDesigns()

	var count int
	for _, d := range designs {
		count += Count(d, towels, map[string]int{"": 1})
	}
	fmt.Println(count)
}

func Count(design string, towels []string, memo map[string]int) int {
	if v, ok := memo[design]; ok {
		return v
	}

	var count int
	for _, towel := range towels {
		if !strings.HasPrefix(design, towel) {
			continue
		}
		count += Count(design[len(towel):], towels, memo)
	}

	memo[design] = count
	return count
}

func InputToTowelsAndDesigns() ([]string, []string) {
	lines := InputToLines()

	towels := strings.Fields(strings.ReplaceAll(lines[0], ",", ""))
	return towels, lines[2:]
}
