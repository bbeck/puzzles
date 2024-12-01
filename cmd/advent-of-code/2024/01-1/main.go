package main

import (
	"fmt"
	"sort"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	as, bs := InputToNums()
	sort.Ints(as)
	sort.Ints(bs)

	var sum int
	for i := range len(as) {
		sum += Abs(as[i] - bs[i])
	}
	fmt.Println(sum)
}

func InputToNums() ([]int, []int) {
	var as, bs []int
	for _, line := range InputToLines() {
		fields := strings.Fields(line)
		as = append(as, ParseInt(fields[0]))
		bs = append(bs, ParseInt(fields[1]))
	}
	return as, bs
}
