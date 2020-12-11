package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ferry := InputToFerry(2020, 11)

	last := ferry.String()
	for {
		ferry = Next(ferry)

		current := ferry.String()
		if current == last {
			break
		}
		last = current
	}

	var count int
	for _, status := range ferry.seats {
		if status == Occupied {
			count++
		}
	}
	fmt.Println(count)
}

func Next(f Ferry) Ferry {
	next := Ferry{
		seats:  make(map[aoc.Point2D]string),
		width:  f.width,
		height: f.height,
	}

	for seat, status := range f.seats {
		var occupied int
		for _, neighbor := range f.Neighbors(seat) {
			if f.seats[neighbor] == Occupied {
				occupied++
			}
		}

		if status != Occupied && occupied == 0 {
			next.seats[seat] = Occupied
		} else if status == Occupied && occupied >= 5 {
			next.seats[seat] = Empty
		} else {
			next.seats[seat] = status
		}
	}

	return next
}

var (
	Floor    = ""
	Empty    = "L"
	Occupied = "#"
)

type Ferry struct {
	seats         map[aoc.Point2D]string
	width, height int
}

func (f Ferry) Neighbors(seat aoc.Point2D) []aoc.Point2D {
	var neighbors []aoc.Point2D
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			// Continue going in the current direction until a non-floor seat is found
			// or we go out of bounds.
			p := seat
			for {
				p = aoc.Point2D{X: p.X + dx, Y: p.Y + dy}
				if p.X < 0 || p.X > f.width || p.Y < 0 || p.Y > f.height {
					break
				}

				if f.seats[p] != Floor {
					neighbors = append(neighbors, p)
					break
				}
			}
		}
	}

	return neighbors
}

func (f Ferry) String() string {
	var sb strings.Builder
	for y := 0; y <= f.height; y++ {
		for x := 0; x <= f.width; x++ {
			p := aoc.Point2D{X: x, Y: y}
			switch f.seats[p] {
			case Floor:
				sb.WriteString(".")
			case Empty:
				sb.WriteString("L")
			case Occupied:
				sb.WriteString("#")
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func InputToFerry(year, day int) Ferry {
	seats := make(map[aoc.Point2D]string)
	positions := make([]aoc.Point2D, 0)
	for y, line := range aoc.InputToLines(year, day) {
		for x, c := range line {
			p := aoc.Point2D{X: x, Y: y}
			if c == 'L' {
				seats[p] = Empty
			}
			positions = append(positions, p)
		}
	}

	_, _, width, height := aoc.GetBounds(positions)
	return Ferry{
		seats:  seats,
		width:  width,
		height: height,
	}
}
