package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
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
	seats.ForEach(func(_, _ int, value uint8) {
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

func Next(seats puz.Grid2D[uint8]) puz.Grid2D[uint8] {
	next := puz.NewGrid2D[uint8](seats.Width, seats.Height)
	seats.ForEach(func(x, y int, value uint8) {
		var count int
		seats.ForEachNeighbor(x, y, func(_, _ int, v uint8) {
			if v == Occupied {
				count++
			}
		})

		if value == Empty && count == 0 {
			next.Set(x, y, Occupied)
		} else if value == Occupied && count >= 4 {
			next.Set(x, y, Empty)
		} else {
			next.Set(x, y, value)
		}
	})
	return next
}

func Equals(a, b puz.Grid2D[uint8]) bool {
	for y := 0; y < a.Height; y++ {
		for x := 0; x < a.Width; x++ {
			if a.Get(x, y) != b.Get(x, y) {
				return false
			}
		}
	}
	return true
}

func InputToSeats() puz.Grid2D[uint8] {
	return puz.InputToGrid2D(2020, 11, func(x int, y int, s string) uint8 {
		return s[0]
	})
}
