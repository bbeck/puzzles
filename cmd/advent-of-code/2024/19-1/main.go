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
		if IsPossible(d, towels) {
			count++
		}
	}
	fmt.Println(count)
}

func IsPossible(design string, towels []string) bool {
	if design == "" {
		return true
	}

	for _, towel := range towels {
		if !strings.HasPrefix(design, towel) {
			continue
		}

		if IsPossible(design[len(towel):], towels) {
			return true
		}
	}

	return false
}

func InputToTowelsAndDesigns() ([]string, []string) {
	lines := InputToLines()

	towels := strings.Fields(strings.ReplaceAll(lines[0], ",", ""))
	return towels, lines[2:]
}
