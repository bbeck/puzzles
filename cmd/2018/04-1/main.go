package main

import (
	"fmt"
	"math/bits"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	activities := InputToActivities(2018, 4)

	// Find the guard that was asleep the most overall number of minutes.
	counts := make(map[int]int)
	for _, activity := range activities {
		counts[activity.guard] += bits.OnesCount64(activity.schedule)
	}

	var guard, count int
	for id, c := range counts {
		if c > count {
			guard = id
			count = c
		}
	}
	fmt.Printf("guard %d was alseep for %d minutes\n", guard, count)

	// Now that we've selected a guard, determine which minute he was most asleep
	// for.
	counts = make(map[int]int)
	for _, activity := range activities {
		if activity.guard != guard {
			continue
		}

		for m := 0; m <= 60; m++ {
			if activity.schedule&(1<<m) > 0 {
				counts[m]++
			}
		}
	}

	var minute, max int
	for m, c := range counts {
		if c > max {
			minute = m
			max = c
		}
	}
	fmt.Printf("guard was asleep %d times during minute %d\n", max, minute)
	fmt.Printf("id * minute: %d\n", guard*minute)
}

type Activity struct {
	date     string
	guard    int
	schedule uint64
}

func InputToActivities(year, day int) []Activity {
	lines := aoc.InputToLines(year, day)

	// The lines are not sorted in time order, so we need to sort them.
	// Fortunately they use a variant of ISO8601 so a simple textual sort should
	// work.
	sort.Strings(lines)

	var activities []Activity

	var activity *Activity
	var asleep int
	for _, line := range lines {
		date := strings.ReplaceAll(strings.Split(line, " ")[0], "[", "")
		minute := aoc.ParseInt(strings.Split(strings.Split(line, ":")[1], "]")[0])

		if strings.Contains(line, "begins shift") {
			if activity != nil {
				activities = append(activities, *activity)
			}

			id := aoc.ParseInt(strings.Split(strings.Split(line, "#")[1], " ")[0])
			activity = &Activity{date: date, guard: id}
			continue
		}

		if strings.Contains(line, "falls asleep") {
			activity.date = date
			asleep = minute
			continue
		}

		if strings.Contains(line, "wakes up") {
			for m := asleep; m < minute; m++ {
				activity.schedule |= 1 << m
			}
			continue
		}
	}

	activities = append(activities, *activity)

	return activities
}
