package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	start, buses := InputToBuses(2020, 13)
	for tm := start; ; tm++ {
		if bus := FindDepartingBus(tm, buses); bus != -1 {
			fmt.Println(bus * (tm - start))
			break
		}
	}
}

func FindDepartingBus(tm int, buses []int) int {
	for _, bus := range buses {
		if tm%bus == 0 {
			return bus
		}
	}

	return -1
}

func InputToBuses(year, day int) (int, []int) {
	lines := aoc.InputToLines(year, day)
	tm := aoc.ParseInt(lines[0])

	var ids []int
	for _, id := range strings.Split(lines[1], ",") {
		if id != "x" {
			ids = append(ids, aoc.ParseInt(id))
		}
	}

	return tm, ids
}
