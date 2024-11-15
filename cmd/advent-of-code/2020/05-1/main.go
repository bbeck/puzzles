package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	replacer := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")

	var max int
	for _, line := range puz.InputToLines() {
		max = puz.Max(max, puz.ParseIntWithBase(replacer.Replace(line), 2))
	}

	fmt.Println(max)
}
