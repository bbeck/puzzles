package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

type State struct {
	Name      string
	Remaining int
	Traveled  int
	Points    int
}

func main() {
	reindeer := InputToReindeer()

	states := make(map[string]*State)
	for _, reindeer := range reindeer {
		states[reindeer.Name] = &State{Name: "fly", Remaining: reindeer.Fly}
	}

	for tm := 0; tm < 2503; tm++ {
		var distances []int
		for _, reindeer := range reindeer {
			state := states[reindeer.Name]

			if state.Name == "fly" {
				state.Traveled += reindeer.Speed
			}
			distances = append(distances, state.Traveled)

			// Check for a state transition
			state.Remaining--
			if state.Remaining == 0 && state.Name == "fly" {
				state.Name = "rest"
				state.Remaining = reindeer.Rest
			} else if state.Remaining == 0 && state.Name == "rest" {
				state.Name = "fly"
				state.Remaining = reindeer.Fly
			}
		}

		// Award points to the reindeer in the lead
		furthest := lib.Max(distances...)
		for _, state := range states {
			if state.Traveled == furthest {
				state.Points++
			}
		}
	}

	var best int
	for _, state := range states {
		best = lib.Max(best, state.Points)
	}
	fmt.Println(best)
}

type ReindeerOld struct {
	name                  string
	speed, duration, rest int

	traveled           int
	resting            bool
	tmToNextTransition int
}

type Reindeer struct {
	Name  string
	Speed int
	Fly   int
	Rest  int
}

func InputToReindeer() []Reindeer {
	return lib.InputLinesTo(func(line string) Reindeer {
		line = strings.ReplaceAll(line, " can fly ", " ")
		line = strings.ReplaceAll(line, " km/s for ", " ")
		line = strings.ReplaceAll(line, " seconds, but then must rest for ", " ")
		line = strings.ReplaceAll(line, " seconds.", "")

		var reindeer Reindeer
		if _, err := fmt.Sscanf(line, "%s %d %d %d", &reindeer.Name, &reindeer.Speed, &reindeer.Fly, &reindeer.Rest); err != nil {
			log.Fatalf("unable to parse line: %v", err)
		}
		return reindeer
	})
}
