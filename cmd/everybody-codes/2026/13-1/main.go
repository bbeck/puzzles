package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	hands := InputToHands()

	var count int
	count += CountOverlaps(hands[0], hands[1])
	count += CountOverlaps(hands[0], hands[2])
	count += CountOverlaps(hands[1], hands[2])
	fmt.Println(count)
}

func CountOverlaps(a, b Hand) int {
	N := a.Stops
	aRemaining, bRemaining := a.Remaining, b.Remaining
	aDistance, bDistance := a.Direction*a.Distance, b.Direction*b.Distance

	var count int

	var gap = (a.Current - b.Current + N) % N
	for aRemaining > 0 || bRemaining > 0 {
		var delta int
		if aRemaining > 0 {
			delta += aDistance
			aRemaining--
		}
		if bRemaining > 0 {
			delta -= bDistance
			bRemaining--
		}

		// If the gap is increasing (`delta` > 0) then we want to find a, b
		// such that:
		//   (a−1)*N <= gap < a*N  AND  b*N <= gap+delta < (b+1)*N
		//
		// Thus a*N is the multiple of N immediately after `before` and b*N is the
		// multiple of N immediately before or equal to `after`.
		//
		// Solving for a, b we have:
		//   a = floor(gap/N)+1
		//   b = floor((gap+delta)/N)
		//
		// This results in there being b-a+1 multiples of N between the current and
		// next positions.  If the gap is decreasing the math is the same except we
		// need to negate both inequalities.
		//
		// Lastly, we have to focus on the floor computation.  This needs to round
		// down towards the nearest multiple of N, not towards zero like integer
		// division does (we add N to each numerator to make everything positive so
		// normal integer division works).
		if delta > 0 {
			count += (gap+delta+N)/N - (gap+N)/N
		} else {
			count += (-gap-delta+N)/N - (-gap+N)/N
		}

		gap = (gap + delta + N) % N
	}
	return count
}

type Hand struct {
	Stops, Direction, Current, Remaining, Distance int
}

func InputToHands() [3]Hand {
	var dirs = map[string]int{"CW": 1, "CCW": -1}

	var hands [3]Hand
	for in.HasNext() {
		switch in.OneOf("hands", "direction", "current", "move") {
		case "hands":
			hands[0].Stops, hands[1].Stops, hands[2].Stops = in.Int(), in.Int(), in.Int()
		case "direction":
			in.Byte() // space after "direction"
			ds := in.Split(":")
			hands[0].Direction, hands[1].Direction, hands[2].Direction = dirs[ds[0]], dirs[ds[1]], dirs[ds[2]]
		case "current":
			hands[0].Current, hands[1].Current, hands[2].Current = in.Int(), in.Int(), in.Int()
		case "move":
			hands[0].Remaining, hands[1].Remaining, hands[2].Remaining = in.Int(), in.Int(), in.Int()
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
	return hands
}
