package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
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

	fmt.Println(lib.Sum(lib.GetMapValues(counts)...))
}

func InputToFish() []int {
	line := lib.InputToString()

	var fs []int
	for _, s := range strings.Split(strings.TrimSpace(line), ",") {
		fs = append(fs, lib.ParseInt(s))
	}
	return fs
}
