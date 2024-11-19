package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"sort"
)

func main() {
	carts, track := InputToCarts(), InputToTrack()

	for tm := 0; len(carts) > 1; tm++ {
		// Rearrange the carts to be in their move order.
		sort.Slice(carts, func(i, j int) bool {
			return carts[i].Location.Y < carts[j].Location.Y ||
				(carts[i].Location.Y == carts[j].Location.Y && carts[i].Location.X < carts[j].Location.X)
		})

		// Keep track of the indices of which carts are removed (so we don't modify
		// the slice while iterating over it).
		var removed lib.Set[int]

		// Move each cart, checking for a collision with another cart afterwards.
		for i := 0; i < len(carts); i++ {
			if removed.Contains(i) {
				continue
			}

			switch track[carts[i].Location] {
			case "/":
				if carts[i].Heading == lib.Up || carts[i].Heading == lib.Down {
					carts[i].TurnRight()
				} else {
					carts[i].TurnLeft()
				}

			case "\\":
				if carts[i].Heading == lib.Up || carts[i].Heading == lib.Down {
					carts[i].TurnLeft()
				} else {
					carts[i].TurnRight()
				}

			case "+":
				switch carts[i].Turns % 3 {
				case 0:
					carts[i].TurnLeft()
				case 2:
					carts[i].TurnRight()
				}
				carts[i].Turns++
			}

			carts[i].Forward(1)

			// Check for a collision with another cart
			for j := 0; j < len(carts); j++ {
				if i == j || removed.Contains(j) {
					continue
				}

				if carts[i].Location == carts[j].Location {
					removed.Add(i, j)
				}
			}
		}

		// Update the carts slice to not include any removed carts.
		var next []Cart
		for i := 0; i < len(carts); i++ {
			if !removed.Contains(i) {
				next = append(next, carts[i])
			}
		}
		carts = next
	}

	fmt.Printf("%d,%d\n", carts[0].Location.X, carts[0].Location.Y)
}

type Cart struct {
	lib.Turtle
	Turns int
}

func InputToTrack() map[lib.Point2D]string {
	track := make(map[lib.Point2D]string)
	for y, line := range lib.InputToLines() {
		for x, c := range line {
			p := lib.Point2D{X: x, Y: y}
			switch c {
			case ' ':
				continue
			case '^':
				track[p] = "|"
			case '>':
				track[p] = "-"
			case 'v':
				track[p] = "|"
			case '<':
				track[p] = "-"
			default:
				track[p] = string(c)
			}
		}
	}
	return track
}

func InputToCarts() []Cart {
	var carts []Cart
	for y, line := range lib.InputToLines() {
		for x, c := range line {
			cart := Cart{
				Turtle: lib.Turtle{Location: lib.Point2D{X: x, Y: y}},
			}

			switch c {
			case '^':
				cart.Heading = lib.Up
			case '>':
				cart.Heading = lib.Right
			case 'v':
				cart.Heading = lib.Down
			case '<':
				cart.Heading = lib.Left
			default:
				continue
			}

			carts = append(carts, cart)
		}
	}

	return carts
}
