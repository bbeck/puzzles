package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	seats := InputToSeats()

	neighbors := make(map[aoc.Point2D][]aoc.Point2D)
	for y := 0; y < seats.Height; y++ {
		for x := 0; x < seats.Width; x++ {
			p := aoc.Point2D{X: x, Y: y}
			neighbors[p] = GetVisibleNeighbors(p, seats)
		}
	}

	for {
		next := Next(seats, neighbors)
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
	NotASeat = '.'
	Empty    = 'L'
	Occupied = '#'
)

func Next(seats aoc.Grid2D[uint8], neighbors map[aoc.Point2D][]aoc.Point2D) aoc.Grid2D[uint8] {
	next := aoc.NewGrid2D[uint8](seats.Width, seats.Height)
	seats.ForEach(func(p aoc.Point2D, value uint8) {
		var count int
		for _, neighbor := range neighbors[p] {
			if seats.Get(neighbor) == Occupied {
				count++
			}
		}

		if value == Empty && count == 0 {
			next.Add(p, Occupied)
		} else if value == Occupied && count >= 5 {
			next.Add(p, Empty)
		} else {
			next.Add(p, value)
		}
	})
	return next
}

func GetVisibleNeighbors(p aoc.Point2D, seats aoc.Grid2D[uint8]) []aoc.Point2D {
	var neighbors []aoc.Point2D
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			q := aoc.Point2D{X: p.X + dx, Y: p.Y + dy}
			for seats.InBounds(q) && seats.Get(q) == NotASeat {
				q.X += dx
				q.Y += dy
			}

			if seats.InBounds(q) {
				neighbors = append(neighbors, q)
			}
		}
	}

	return neighbors
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
