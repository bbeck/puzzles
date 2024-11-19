package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	seats := InputToSeats()

	neighbors := make(map[lib.Point2D][]lib.Point2D)
	seats.ForEachPoint(func(p lib.Point2D, value uint8) {
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

func Next(seats lib.Grid2D[uint8], neighbors map[lib.Point2D][]lib.Point2D) lib.Grid2D[uint8] {
	next := lib.NewGrid2D[uint8](seats.Width, seats.Height)
	seats.ForEachPoint(func(p lib.Point2D, value uint8) {
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

func GetVisibleNeighbors(p lib.Point2D, seats lib.Grid2D[uint8]) []lib.Point2D {
	var neighbors []lib.Point2D
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			q := lib.Point2D{X: p.X + dx, Y: p.Y + dy}
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

func Equals(a, b lib.Grid2D[uint8]) bool {
	for y := 0; y < a.Height; y++ {
		for x := 0; x < a.Width; x++ {
			if a.Get(x, y) != b.Get(x, y) {
				return false
			}
		}
	}
	return true
}

func InputToSeats() lib.Grid2D[uint8] {
	return lib.InputToGrid2D(func(x int, y int, s string) uint8 {
		return s[0]
	})
}
