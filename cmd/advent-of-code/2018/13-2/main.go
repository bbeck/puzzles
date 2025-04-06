package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
	"sort"
)

func main() {
	carts, track := InputToCartsAndTrack()

	for tm := 0; len(carts) > 1; tm++ {
		// Rearrange the carts to be in their move order.
		sort.Slice(carts, func(i, j int) bool {
			return carts[i].Location.Y < carts[j].Location.Y ||
				(carts[i].Location.Y == carts[j].Location.Y && carts[i].Location.X < carts[j].Location.X)
		})

		// Keep track of the indices of which carts are removed (so we don't modify
		// the slice while iterating over it).
		var removed Set[int]

		// Move each cart, checking for a collision with another cart afterwards.
		for i := 0; i < len(carts); i++ {
			if removed.Contains(i) {
				continue
			}

			switch track[carts[i].Location] {
			case "/":
				if carts[i].Heading == Up || carts[i].Heading == Down {
					carts[i].TurnRight()
				} else {
					carts[i].TurnLeft()
				}

			case "\\":
				if carts[i].Heading == Up || carts[i].Heading == Down {
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
	Turtle
	Turns int
}

func InputToCartsAndTrack() ([]Cart, map[Point2D]string) {
	var carts []Cart
	var track = make(map[Point2D]string)

	for y := 0; in.HasNext(); y++ {
		for x, ch := range in.Line() {
			p := Point2D{X: x, Y: y}

			switch ch {
			case '^':
				track[p] = "|"
				carts = append(carts, Cart{Turtle: Turtle{Location: p, Heading: Up}})

			case '>':
				track[p] = "-"
				carts = append(carts, Cart{Turtle: Turtle{Location: p, Heading: Right}})

			case 'v':
				track[p] = "|"
				carts = append(carts, Cart{Turtle: Turtle{Location: p, Heading: Down}})

			case '<':
				track[p] = "-"
				carts = append(carts, Cart{Turtle: Turtle{Location: p, Heading: Left}})

			default:
				track[p] = string(ch)
			}
		}
	}

	return carts, track
}
