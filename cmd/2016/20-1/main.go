package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ranges := InputToRanges()
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	var ip int
	for _, r := range ranges {
		if r.Start > ip {
			break
		}

		// It's possible that the ending ip of this range is earlier than our
		// current ip.  When this happens don't update where we currently are
		// or else we'll double-check ips.
		if r.End+1 > ip {
			ip = r.End + 1
		}
	}

	fmt.Println(ip)
}

type Range struct {
	Start, End int
}

func InputToRanges() []Range {
	return aoc.InputLinesTo(2016, 20, func(line string) (Range, error) {
		start, end, _ := strings.Cut(line, "-")
		return Range{
			Start: aoc.ParseInt(start),
			End:   aoc.ParseInt(end),
		}, nil
	})
}
