package main

import (
	"fmt"
	"sort"

	"github.com/bbeck/puzzles/lib/in"
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
	return in.LinesToS(func(in in.Scanner[Range]) Range {
		var start, end int
		in.Scanf("%d-%d", &start, &end)
		return Range{Start: start, End: end}
	})
}
