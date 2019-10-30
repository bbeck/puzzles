package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

type Reindeer struct {
	name                  string
	speed, duration, rest int

	traveled           int
	resting            bool
	tmToNextTransition int
}

func main() {
	reindeer := InputToReindeer(2015, 14)
	scores := make([]int, len(reindeer))

	for tm := 0; tm < 2503; tm++ {
		for i := 0; i < len(reindeer); i++ {
			if !reindeer[i].resting {
				reindeer[i].traveled += reindeer[i].speed
			}

			reindeer[i].tmToNextTransition--
			if reindeer[i].tmToNextTransition == 0 {
				// Time to switch states
				if reindeer[i].resting {
					reindeer[i].resting = false
					reindeer[i].tmToNextTransition = reindeer[i].duration
				} else {
					reindeer[i].resting = true
					reindeer[i].tmToNextTransition = reindeer[i].rest
				}
			}
		}

		// Determine how far the reindeer in the lead have traveled
		var best int
		for i := 0; i < len(reindeer); i++ {
			if reindeer[i].traveled > best {
				best = reindeer[i].traveled
			}
		}

		// Award each reindeer who has traveled the furthest a point
		for i := 0; i < len(reindeer); i++ {
			if reindeer[i].traveled == best {
				scores[i]++
			}
		}
	}

	var best int
	for _, score := range scores {
		if score > best {
			best = score
		}
	}

	fmt.Printf("best score: %d\n", best)
}

func InputToReindeer(year, day int) []Reindeer {
	var reindeer []Reindeer
	for _, line := range aoc.InputToLines(year, day) {
		var name string
		var speed, duration, rest int
		if _, err := fmt.Sscanf(line[:len(line)-1], "%s can fly %d km/s for %d seconds, but then must rest for %d seconds", &name, &speed, &duration, &rest); err != nil {
			log.Fatalf("unable to parse line: %s", line)
		}

		reindeer = append(reindeer, Reindeer{
			name:               name,
			speed:              speed,
			duration:           duration,
			rest:               rest,
			resting:            false,
			tmToNextTransition: duration,
		})
	}

	return reindeer
}
