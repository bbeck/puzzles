package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
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
		furthest := Max(distances...)
		for _, state := range states {
			if state.Traveled == furthest {
				state.Points++
			}
		}
	}

	var best int
	for _, state := range states {
		best = Max(best, state.Points)
	}
	fmt.Println(best)
}

type Reindeer struct {
	Name  string
	Speed int
	Fly   int
	Rest  int
}

func InputToReindeer() []Reindeer {
	var reindeer []Reindeer
	for in.HasNext() {
		var r Reindeer
		in.Scanf("%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &r.Name, &r.Speed, &r.Fly, &r.Rest)
		reindeer = append(reindeer, r)
	}
	return reindeer
}
