package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib/in"
	"sort"
)

func main() {
	guards := InputToGuards()

	var guard, minute, most int
	for id := range guards {
		for tm := 0; tm < 60; tm++ {
			if guards[id][tm] > most {
				guard = id
				minute = tm
				most = guards[id][tm]
			}
		}
	}

	fmt.Println(guard * minute)
}

func InputToGuards() map[int][60]int {
	type DatedMessage struct {
		in.Scanner[DatedMessage]

		Date   string
		Minute int
	}

	var messages = in.LinesToS(func(s in.Scanner[DatedMessage]) DatedMessage {
		var year, month, day, hour, minute int
		var message in.Scanner[DatedMessage]
		s.Scanf("[%d-%d-%d %d:%d] %s", &year, &month, &day, &hour, &minute, &message)

		return DatedMessage{
			Scanner: message,
			Date:    fmt.Sprintf("%4d-%02d-%02d %02d:%02d", year, month, day, hour, minute),
			Minute:  minute,
		}
	})

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].Date < messages[j].Date
	})

	var guards = make(map[int][60]int)

	var id, asleep int
	for _, message := range messages {
		switch {
		case message.HasPrefix("Guard"):
			message.Scanf("Guard #%d begins shift", &id)

		case message.HasPrefix("falls asleep"):
			asleep = message.Minute

		case message.HasPrefix("wakes up"):
			for tm := asleep; tm < message.Minute; tm++ {
				guard := guards[id]
				guard[tm]++
				guards[id] = guard
			}
		}
	}

	return guards
}
