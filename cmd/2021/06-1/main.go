package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	counts := make(map[int]int)
	for _, fish := range InputToFish() {
		counts[fish]++
	}

	for day := 1; day <= 80; day++ {
		next := make(map[int]int)
		for tm, count := range counts {
			if tm == 0 {
				next[6] += count
				next[8] += count
				continue
			}
			next[tm-1] += count
		}
		counts = next
	}

	fmt.Println(puz.Sum(puz.GetMapValues(counts)...))
}

func InputToFish() []int {
	line := puz.InputToString(2021, 6)

	var fs []int
	for _, s := range strings.Split(strings.TrimSpace(line), ",") {
		fs = append(fs, puz.ParseInt(s))
	}
	return fs
}
