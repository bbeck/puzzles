package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	replacer := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")

	var ids []int
	for _, line := range aoc.InputToLines(2020, 05) {
		ids = append(ids, ParseBinary(replacer.Replace(line)))
	}

	sort.Ints(ids)

	for i := 0; i < len(ids)-1; i++ {
		if ids[i]+1 != ids[i+1] {
			fmt.Println(ids[i] + 1)
		}
	}
}

func ParseBinary(s string) int {
	n, err := strconv.ParseInt(s, 2, 16)
	if err != nil {
		log.Fatalf("unable to parse binary number: %s", s)
	}

	return int(n)
}
