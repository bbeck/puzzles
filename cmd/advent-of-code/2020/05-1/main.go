package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	replacer := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")

	var max int
	for _, line := range lib.InputToLines() {
		max = lib.Max(max, lib.ParseIntWithBase(replacer.Replace(line), 2))
	}

	fmt.Println(max)
}
