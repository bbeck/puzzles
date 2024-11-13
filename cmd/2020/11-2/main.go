package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	seats := InputToSeats()

	neighbors := make(map[puz.Point2D][]puz.Point2D)
	seats.ForEachPoint(func(p puz.Point2D, value uint8) {
		neighbors[p] = GetVisibleNeighbors(p, seats)
	})

	for {
		next := Next(seats, neighbors)
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
	NotASeat = '.'
	Empty    = 'L'
	Occupied = '#'
)

func Next(seats puz.Grid2D[uint8], neighbors map[puz.Point2D][]puz.Point2D) puz.Grid2D[uint8] {
	next := puz.NewGrid2D[uint8](seats.Width, seats.Height)
	seats.ForEachPoint(func(p puz.Point2D, value uint8) {
		var count int
		for _, neighbor := range neighbors[p] {
			if seats.GetPoint(neighbor) == Occupied {
				count++
			}
		}

		if value == Empty && count == 0 {
			next.SetPoint(p, Occupied)
		} else if value == Occupied && count >= 5 {
			next.SetPoint(p, Empty)
		} else {
			next.SetPoint(p, value)
		}
	})
	return next
}

func GetVisibleNeighbors(p puz.Point2D, seats puz.Grid2D[uint8]) []puz.Point2D {
	var neighbors []puz.Point2D
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			q := puz.Point2D{X: p.X + dx, Y: p.Y + dy}
			for seats.InBoundsPoint(q) && seats.GetPoint(q) == NotASeat {
				q.X += dx
				q.Y += dy
			}

			if seats.InBoundsPoint(q) {
				neighbors = append(neighbors, q)
			}
		}
	}

	return neighbors
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
