package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`mul[(]([0-9]+),([0-9]+)[)]|do[(][)]|don't[(][)]`)

func main() {
	enabled := true

	var sum int
	for _, m := range regex.FindAllString(InputToString(), -1) {
		m = strings.ReplaceAll(m, "(", " ")
		m = strings.ReplaceAll(m, ")", " ")
		m = strings.ReplaceAll(m, ",", " ")

		switch fields := strings.Fields(m); fields[0] {
		case "do":
			enabled = true
		case "don't":
			enabled = false
		case "mul":
			if enabled {
				sum += ParseInt(fields[1]) * ParseInt(fields[2])
			}
		}
	}
	fmt.Println(sum)
}
