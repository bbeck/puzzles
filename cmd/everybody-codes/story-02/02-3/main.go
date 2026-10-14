package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

const N int = 100000

var Arrows = []string{"R", "G", "B"}

func main() {
	var balloons = strings.Repeat(in.String(), N)
	var half = len(balloons) / 2

	var left, right Deque[string]
	for i := range half {
		left.PushBack(balloons[i : i+1])
		right.PushBack(balloons[half+i : half+i+1])
	}

	var count int
	for !left.Empty() {
		if Arrows[count%3] != left.PeekFront() || left.Len() != right.Len() {
			// We're going to pop 1 balloon
			left.PopFront()

			// The deque lengths may be out of sync so attempt to rebalance.
			for left.Len() < right.Len() {
				left.PushBack(right.PopFront())
			}
		} else {
			// We're going to pop 2 balloons
			left.PopFront()
			right.PopFront()
		}

		count++
	}
	fmt.Println(count)
}
