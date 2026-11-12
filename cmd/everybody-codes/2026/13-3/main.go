package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	hands := InputToHands()

	var count int
	for _, hs := range hands {
		count += CountOverlaps(hs[0], hs[1])
		count += CountOverlaps(hs[0], hs[2])
		count += CountOverlaps(hs[1], hs[2])
	}
	fmt.Println(count)
}

func CountOverlaps(a, b Hand) int {
	// The approach here is going to be to consider two stretches of time.  The
	// stretch where both hands are moving and the stretch where only one hand is
	// moving.  For simplicity, we'll assume that hand a always takes fewer steps
	// than hand b.
	if b.Remaining < a.Remaining {
		a, b = b, a
	}

	// The length of each stretch.
	len1 := Min(a.Remaining, b.Remaining)
	len2 := Max(a.Remaining, b.Remaining) - len1

	// Compute the gap at the beginning and end of the first stretch.
	gap0 := a.Current - b.Current
	gap1 := gap0 + len1*(a.Direction*a.Distance-b.Direction*b.Distance)

	var count int
	if delta1 := a.Direction*a.Distance - b.Direction*b.Distance; delta1 > 0 {
		count += (Modulo(gap0, a.Stops) + len1*delta1) / a.Stops
	} else {
		count += (Modulo(-gap0, a.Stops) - len1*delta1) / a.Stops
	}
	if delta2 := -b.Direction * b.Distance; delta2 > 0 {
		count += (Modulo(gap1, a.Stops) + len2*delta2) / a.Stops
	} else {
		count += (Modulo(-gap1, a.Stops) - len2*delta2) / a.Stops
	}
	return count
}

type Hand struct {
	Stops, Direction, Current, Remaining, Distance int
}

func InputToHands() [][3]Hand {
	var dirs = map[string]int{"CW": 1, "CCW": -1}

	var allHands [][3]Hand
	for in.HasNext() {
		var hands [3]Hand

		chunk := in.ChunkS[string]()
		for chunk.HasNext() {
			switch chunk.OneOf("hands", "direction", "current", "move") {
			case "hands":
				hands[0].Stops, hands[1].Stops, hands[2].Stops = chunk.Int(), chunk.Int(), chunk.Int()
			case "direction":
				chunk.Byte() // space after "direction"
				ds := chunk.Split(":")
				hands[0].Direction, hands[1].Direction, hands[2].Direction = dirs[ds[0]], dirs[ds[1]], dirs[ds[2]]
			case "current":
				hands[0].Current, hands[1].Current, hands[2].Current = chunk.Int(), chunk.Int(), chunk.Int()
			case "move":
				hands[0].Remaining, hands[1].Remaining, hands[2].Remaining = chunk.Int(), chunk.Int(), chunk.Int()
			}
		}

		// Normalize the hands to have the same number of stops.  We do this by
		// scaling things up to the LCM of the number of stops.
		N := LCM(hands[0].Stops, hands[1].Stops, hands[2].Stops)
		for i := range hands {
			hands[i].Distance = N / hands[i].Stops
			hands[i].Current *= N / hands[i].Stops
			hands[i].Stops = N
		}

		allHands = append(allHands, hands)
	}

	return allHands
}
