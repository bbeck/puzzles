package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	ids := InputToLines()

	for i := 0; i < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			if dist, common := Distance(ids[i], ids[j]); dist == 1 {
				fmt.Println(common)
			}
		}
	}
}

func Distance(a, b string) (int, string) {
	var dist int
	var sb strings.Builder
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
			continue
		}

		sb.WriteByte(a[i])
	}

	return dist, sb.String()
}
