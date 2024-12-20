package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	ranges := InputToRanges()
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	var ip, count int
	for _, r := range ranges {
		if r.Start > ip {
			count += r.Start - ip
		}

		// It's possible that the ending ip of this range is earlier than our
		// current ip.  When this happens don't update where we currently are
		// or else we'll double-check ips.
		if r.End+1 > ip {
			ip = r.End + 1
		}
	}

	fmt.Println(count)
}

type Range struct {
	Start, End int
}

func InputToRanges() []Range {
	return lib.InputLinesTo(func(line string) Range {
		start, end, _ := strings.Cut(line, "-")
		return Range{
			Start: lib.ParseInt(start),
			End:   lib.ParseInt(end),
		}
	})
}
