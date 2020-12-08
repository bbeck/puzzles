package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	replacer := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")

	var max int
	for _, line := range aoc.InputToLines(2020, 5) {
		id := ParseBinary(replacer.Replace(line))
		if id > max {
			max = id
		}
	}

	fmt.Println(max)
}

func ParseBinary(s string) int {
	n, err := strconv.ParseInt(s, 2, 16)
	if err != nil {
		log.Fatalf("unable to parse binary number: %s", s)
	}

	return int(n)
}
