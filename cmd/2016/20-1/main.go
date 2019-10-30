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
	ranges := InputToRanges(2016, 20)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	ip := 0
	for _, r := range ranges {
		if r.start > ip {
			fmt.Printf("ip: %d is not filtered\n", ip)
			break
		}

		ip = r.end + 1
	}
}

type Range struct {
	start, end int
}

func InputToRanges(year, day int) []Range {
	var ranges []Range
	for _, line := range aoc.InputToLines(year, day) {
		fields := strings.Split(line, "-")
		if len(fields) != 2 {
			log.Fatalf("unable to parse range: %s", line)
		}

		start, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalf("unable to parse range: %s", line)
		}

		end, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("unable to parse range: %s", line)
		}

		ranges = append(ranges, Range{start, end})
	}

	return ranges
}
