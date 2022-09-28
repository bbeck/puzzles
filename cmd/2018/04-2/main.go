package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"sort"
	"strings"
)

func main() {
	var maxID, maxMinute, maxSleep int
	for _, guard := range InputToGuards() {
		for tm := 0; tm < 60; tm++ {
			if guard.Sleep[tm] > maxSleep {
				maxID = guard.ID
				maxMinute = tm
				maxSleep = guard.Sleep[tm]
			}
		}
	}

	fmt.Println(maxID * maxMinute)
}

type Guard struct {
	ID    int
	Sleep [60]int
}

func InputToGuards() []Guard {
	lines := aoc.InputToLines(2018, 4)
	sort.Strings(lines) // Sort into time order

	sleep := make(map[int][60]int) // number of times each guard was asleep at the given minute

	var current int // the current guard
	var asleep int  // the minute the current guard fell asleep

	for _, line := range lines {
		line = strings.ReplaceAll(line, "[", "")
		line = strings.ReplaceAll(line, "]", "")
		line = strings.ReplaceAll(line, "#", "")
		line = strings.ReplaceAll(line, ":", " ")
		line = strings.ReplaceAll(line, "-", " ")
		fields := strings.Fields(line)

		minute := aoc.ParseInt(fields[4])

		if strings.Contains(line, "begins shift") {
			current = aoc.ParseInt(fields[6])
		}

		if strings.Contains(line, "falls asleep") {
			asleep = minute
		}

		if strings.Contains(line, "wakes up") {
			for tm := asleep; tm < minute; tm++ {
				schedule := sleep[current]
				schedule[tm]++
				sleep[current] = schedule
			}
		}
	}

	var guards []Guard
	for id, schedule := range sleep {
		guards = append(guards, Guard{ID: id, Sleep: schedule})
	}

	return guards
}
