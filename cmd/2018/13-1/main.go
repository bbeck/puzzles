package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"sort"
)

func main() {
	carts, track := InputToCarts(), InputToTrack()

	var collision *aoc.Point2D
	for tm := 0; collision == nil; tm++ {
		// Rearrange the carts to be in their move order.
		sort.Slice(carts, func(i, j int) bool {
			return carts[i].Location.Y < carts[j].Location.Y ||
				(carts[i].Location.Y == carts[j].Location.Y && carts[i].Location.X < carts[j].Location.X)
		})

		// Keep track of the location of each cart to easily find collisions.
		var locations aoc.Set[aoc.Point2D]
		for _, cart := range carts {
			locations.Add(cart.Location)
		}

		// Move each cart, checking for a collision with another cart afterwards.
		for i := 0; i < len(carts); i++ {
			switch track[carts[i].Location] {
			case "/":
				if carts[i].Heading == aoc.Up || carts[i].Heading == aoc.Down {
					carts[i].TurnRight()
				} else {
					carts[i].TurnLeft()
				}

			case "\\":
				if carts[i].Heading == aoc.Up || carts[i].Heading == aoc.Down {
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
	aoc.Turtle
	Turns int
}

func InputToTrack() map[aoc.Point2D]string {
	track := make(map[aoc.Point2D]string)
	for y, line := range aoc.InputToLines(2018, 13) {
		for x, c := range line {
			p := aoc.Point2D{X: x, Y: y}
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
	for y, line := range aoc.InputToLines(2018, 13) {
		for x, c := range line {
			cart := Cart{
				Turtle: aoc.Turtle{Location: aoc.Point2D{X: x, Y: y}},
			}

			switch c {
			case '^':
				cart.Heading = aoc.Up
			case '>':
				cart.Heading = aoc.Right
			case 'v':
				cart.Heading = aoc.Down
			case '<':
				cart.Heading = aoc.Left
			default:
				continue
			}

			carts = append(carts, cart)
		}
	}

	return carts
}
