package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"sort"
	"strings"
)

func main() {
	guards := InputToGuards()
	sort.Slice(guards, func(i, j int) bool {
		return lib.Sum(guards[i].Sleep[:]...) > lib.Sum(guards[j].Sleep[:]...)
	})

	max := lib.Max(guards[0].Sleep[:]...)

	var minute int
	for minute = 0; minute < 60; minute++ {
		if guards[0].Sleep[minute] == max {
			break
		}
	}

	fmt.Println(guards[0].ID * minute)
}

type Guard struct {
	ID    int
	Sleep [60]int
}

func InputToGuards() []Guard {
	lines := lib.InputToLines()
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

		minute := lib.ParseInt(fields[4])

		if strings.Contains(line, "begins shift") {
			current = lib.ParseInt(fields[6])
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
