package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	seats := InputToSeats()
	for {
		next := Next(seats)
		if Equals(seats, next) {
			break
		}

		seats = next
	}

	var count int
	seats.ForEach(func(_ aoc.Point2D, value uint8) {
		if value == Occupied {
			count++
		}
	})
	fmt.Println(count)
}

const (
	Empty    = 'L'
	Occupied = '#'
)

func Next(seats aoc.Grid2D[uint8]) aoc.Grid2D[uint8] {
	next := aoc.NewGrid2D[uint8](seats.Width, seats.Height)
	seats.ForEach(func(p aoc.Point2D, value uint8) {
		var count int
		seats.ForEachNeighbor(p, func(n aoc.Point2D, v uint8) {
			if v == Occupied {
				count++
			}
		})

		if value == Empty && count == 0 {
			next.Add(p, Occupied)
		} else if value == Occupied && count >= 4 {
			next.Add(p, Empty)
		} else {
			next.Add(p, value)
		}
	})
	return next
}

func Equals(a, b aoc.Grid2D[uint8]) bool {
	for y := 0; y < a.Height; y++ {
		for x := 0; x < a.Width; x++ {
			if a.GetXY(x, y) != b.GetXY(x, y) {
				return false
			}
		}
	}
	return true
}

func InputToSeats() aoc.Grid2D[uint8] {
	return aoc.InputToGrid2D(2020, 11, func(x int, y int, s string) uint8 {
		return s[0]
	})
}
