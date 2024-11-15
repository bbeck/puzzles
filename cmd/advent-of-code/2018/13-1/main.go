package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"sort"
)

func main() {
	carts, track := InputToCarts(), InputToTrack()

	var collision *puz.Point2D
	for tm := 0; collision == nil; tm++ {
		// Rearrange the carts to be in their move order.
		sort.Slice(carts, func(i, j int) bool {
			return carts[i].Location.Y < carts[j].Location.Y ||
				(carts[i].Location.Y == carts[j].Location.Y && carts[i].Location.X < carts[j].Location.X)
		})

		// Keep track of the location of each cart to easily find collisions.
		var locations puz.Set[puz.Point2D]
		for _, cart := range carts {
			locations.Add(cart.Location)
		}

		// Move each cart, checking for a collision with another cart afterwards.
		for i := 0; i < len(carts); i++ {
			switch track[carts[i].Location] {
			case "/":
				if carts[i].Heading == puz.Up || carts[i].Heading == puz.Down {
					carts[i].TurnRight()
				} else {
					carts[i].TurnLeft()
				}

			case "\\":
				if carts[i].Heading == puz.Up || carts[i].Heading == puz.Down {
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
	puz.Turtle
	Turns int
}

func InputToTrack() map[puz.Point2D]string {
	track := make(map[puz.Point2D]string)
	for y, line := range puz.InputToLines() {
		for x, c := range line {
			p := puz.Point2D{X: x, Y: y}
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
	for y, line := range puz.InputToLines() {
		for x, c := range line {
			cart := Cart{
				Turtle: puz.Turtle{Location: puz.Point2D{X: x, Y: y}},
			}

			switch c {
			case '^':
				cart.Heading = puz.Up
			case '>':
				cart.Heading = puz.Right
			case 'v':
				cart.Heading = puz.Down
			case '<':
				cart.Heading = puz.Left
			default:
				continue
			}

			carts = append(carts, cart)
		}
	}

	return carts
}
