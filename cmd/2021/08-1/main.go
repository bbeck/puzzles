package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"log"
)

func main() {
	entries := InputToEntries()

	var count int
	for _, entry := range entries {
		for _, output := range entry.Outputs {
			if len(output) == 2 || len(output) == 4 || len(output) == 3 || len(output) == 7 {
				count++
			}
		}
	}
	fmt.Println(count)
}

type Entry struct {
	Signals []string
	Outputs []string
}

func InputToEntries() []Entry {
	var entries []Entry
	for _, line := range aoc.InputToLines(2021, 8) {
		var s0, s1, s2, s3, s4, s5, s6, s7, s8, s9 string
		var o0, o1, o2, o3 string
		_, err := fmt.Sscanf(line, "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&s0, &s1, &s2, &s3, &s4, &s5, &s6, &s7, &s8, &s9, &o0, &o1, &o2, &o3)
		if err != nil {
			log.Fatal(err)
		}

		entries = append(entries, Entry{
			Signals: []string{s0, s1, s2, s3, s4, s5, s6, s7, s8, s9},
			Outputs: []string{o0, o1, o2, o3},
		})
	}
	return entries
}

// var fs []int
// for _, s := range strings.Split(strings.TrimSpace(line), ",") {
// fs = append(fs, aoc.ParseInt(s))
// }
// return fs
