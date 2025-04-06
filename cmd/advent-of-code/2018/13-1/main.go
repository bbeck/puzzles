package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
	"sort"
)

func main() {
	carts, track := InputToCartsAndTrack()

	var collision *Point2D
	for tm := 0; collision == nil; tm++ {
		// Rearrange the carts to be in their move order.
		sort.Slice(carts, func(i, j int) bool {
			return carts[i].Location.Y < carts[j].Location.Y ||
				(carts[i].Location.Y == carts[j].Location.Y && carts[i].Location.X < carts[j].Location.X)
		})

		// Keep track of the location of each cart to easily find collisions.
		var locations Set[Point2D]
		for _, cart := range carts {
			locations.Add(cart.Location)
		}

		// Move each cart, checking for a collision with another cart afterwards.
		for i := 0; i < len(carts); i++ {
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

			locations.Remove(carts[i].Location)
			carts[i].Forward(1)

			if !locations.Add(carts[i].Location) {
				collision = &carts[i].Location
				break
			}
		}
	}

	fmt.Printf("%d,%d\n", collision.X, collision.Y)
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
