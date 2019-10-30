package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

type Sue struct {
	id     int
	fields map[string]int
}

func main() {
	target := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	sues := InputToSues(2015, 16)
	for _, sue := range sues {
		if Matches(sue, target) {
			fmt.Println(sue.id)
		}
	}
}

func Matches(sue Sue, target map[string]int) bool {
	for field, expected := range target {
		if actual, ok := sue.fields[field]; ok {
			if expected != actual {
				return false
			}
		}
	}

	return true
}

func InputToSues(year, day int) []Sue {
	var sues []Sue
	for _, line := range aoc.InputToLines(2015, 16) {
		parts := strings.SplitN(line, ": ", 2)

		id, err := strconv.Atoi(strings.Replace(parts[0], "Sue ", "", 1))
		if err != nil {
			log.Fatalf("unable to parse id: %s", parts[0])
		}

		fields := make(map[string]int)
		for _, field := range strings.Split(parts[1], ", ") {
			tokens := strings.SplitN(field, ": ", 2)

			name := tokens[0]
			count, err := strconv.Atoi(tokens[1])
			if err != nil {
				log.Fatalf("unable to parse field (id=%s): %s: %+v", id, field, err)
			}

			fields[name] = count
		}

		sues = append(sues, Sue{id, fields})
	}

	return sues
}
