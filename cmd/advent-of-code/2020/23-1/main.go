package main

import (
	"container/ring"
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	circle, index := InputToCircle()

	for i := 1; i <= 100; i++ {
		removed := circle.Unlink(3)
		destination := Destination(circle, removed)

		next := circle.Next()

		circle = index[destination]
		circle.Link(removed)
		circle = next
	}

	circle = index[1].Next()
	circle.Do(func(value any) {
		if value.(int) != 1 {
			fmt.Print(value)
		}
	})
	fmt.Println()
}

func Destination(circle, removed *ring.Ring) int {
	var used lib.Set[int]
	removed.Do(func(value any) {
		used.Add(value.(int))
	})

	destination := circle.Value.(int) - 1
	for {
		if destination == 0 {
			destination = 9
		}

		if !used.Contains(destination) {
			return destination
		}

		destination -= 1
	}
}

func InputToCircle() (*ring.Ring, map[int]*ring.Ring) {
	digits := lib.Digits(lib.InputToInt())

	circle := ring.New(len(digits))
	index := make(map[int]*ring.Ring)
	for _, d := range digits {
		circle.Value = d
		index[d] = circle
		circle = circle.Next()
	}

	return circle, index
}
