package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	activities := InputToActivities(2018, 4)

	// Of all of the guards, which guard is most frequently asleep on the same
	// minute?  We'll build a map from guard to []int indicating the count for
	// each minute.
	counts := make(map[int][]int)
	for _, activity := range activities {
		for m := 0; m < 60; m++ {
			if activity.schedule&(1<<m) > 0 {
				if counts[activity.guard] == nil {
					counts[activity.guard] = make([]int, 60)
				}
				counts[activity.guard][m]++
			}
		}
	}

	// Now find the largest count in the matrix.
	var guard, minute, max int
	for id, cs := range counts {
		for m, count := range cs {
			if count > max {
				guard = id
				minute = m
				max = count
			}
		}
	}

	fmt.Printf("guard %d was asleep %d times for minute %d\n", guard, max, minute)
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
