package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	track, carts := InputToTrackAndCarts(2018, 13)

	for len(carts) > 1 {
		carts = Step(track, carts)
	}

	fmt.Printf("last cart location: %s\n", carts[0].location)
}

// Move all of the carts forward one step, stopping if a collision happens.
func Step(track Track, carts []*Cart) []*Cart {
	occupied := make(map[aoc.Point2D]bool)
	for _, c := range carts {
		occupied[c.location] = true
	}

	// order the carts by their move order, top left to bottom right
	sort.Slice(carts, func(i, j int) bool {
		location1, location2 := carts[i].location, carts[j].location
		if location1.Y < location2.Y {
			return true
		}

		if location1.Y == location2.Y && location1.X < location2.X {
			return true
		}

		return false
	})

	removed := make(map[*Cart]bool)
	for _, cart := range carts {
		if removed[cart] {
			continue
		}

		delete(occupied, cart.location)

		// Move the cart
		switch cart.direction {
		case NORTH:
			cart.location = cart.location.Up()
		case SOUTH:
			cart.location = cart.location.Down()
		case WEST:
			cart.location = cart.location.Left()
		case EAST:
			cart.location = cart.location.Right()
		}

		// Determine it's new direction
		switch track[cart.location] {
		case "/":
			switch cart.direction {
			case NORTH:
				cart.direction = EAST
			case SOUTH:
				cart.direction = WEST
			case WEST:
				cart.direction = SOUTH
			case EAST:
				cart.direction = NORTH
			}

		case "\\":
			switch cart.direction {
			case NORTH:
				cart.direction = WEST
			case SOUTH:
				cart.direction = EAST
			case WEST:
				cart.direction = NORTH
			case EAST:
				cart.direction = SOUTH
			}

		case "+":
			switch {
			case cart.direction == NORTH && cart.lastTurn == RIGHT:
				cart.lastTurn = LEFT
				cart.direction = WEST

			case cart.direction == NORTH && cart.lastTurn == STRAIGHT:
				cart.lastTurn = RIGHT
				cart.direction = EAST

			case cart.direction == NORTH && cart.lastTurn == LEFT:
				cart.lastTurn = STRAIGHT
				cart.direction = NORTH

			case cart.direction == SOUTH && cart.lastTurn == RIGHT:
				cart.lastTurn = LEFT
				cart.direction = EAST

			case cart.direction == SOUTH && cart.lastTurn == STRAIGHT:
				cart.lastTurn = RIGHT
				cart.direction = WEST

			case cart.direction == SOUTH && cart.lastTurn == LEFT:
				cart.lastTurn = STRAIGHT
				cart.direction = SOUTH

			case cart.direction == WEST && cart.lastTurn == RIGHT:
				cart.lastTurn = LEFT
				cart.direction = SOUTH

			case cart.direction == WEST && cart.lastTurn == STRAIGHT:
				cart.lastTurn = RIGHT
				cart.direction = NORTH

			case cart.direction == WEST && cart.lastTurn == LEFT:
				cart.lastTurn = STRAIGHT
				cart.direction = WEST

			case cart.direction == EAST && cart.lastTurn == RIGHT:
				cart.lastTurn = LEFT
				cart.direction = NORTH

			case cart.direction == EAST && cart.lastTurn == STRAIGHT:
				cart.lastTurn = RIGHT
				cart.direction = SOUTH

			case cart.direction == EAST && cart.lastTurn == LEFT:
				cart.lastTurn = STRAIGHT
				cart.direction = EAST
			}
		}

		// See if this new location collides with another cart.
		if occupied[cart.location] {
			for _, c := range carts {
				if c.location == cart.location {
					removed[c] = true
				}
			}
		}
		occupied[cart.location] = true
	}

	var next []*Cart
	for _, cart := range carts {
		if !removed[cart] {
			next = append(next, cart)
		}
	}

	return next
}

const (
	// The turns that can be made
	RIGHT int = iota
	LEFT
	STRAIGHT
)

const (
	// The directions that a cart can be traveling in
	NORTH int = iota
	SOUTH
	WEST
	EAST
)

type Track map[aoc.Point2D]string

func (t Track) String(carts []*Cart) string {
	cs := make(map[aoc.Point2D]*Cart)
	for _, cart := range carts {
		cs[cart.location] = cart
	}

	var builder strings.Builder
	for y := 0; ; y++ {
		if _, ok := t[aoc.Point2D{0, y}]; !ok {
			break
		}

		for x := 0; ; x++ {
			p := aoc.Point2D{X: x, Y: y}

			cart, ok := cs[p]
			if ok {
				if cart.direction == NORTH {
					builder.WriteString("^")
				} else if cart.direction == SOUTH {
					builder.WriteString("v")
				} else if cart.direction == WEST {
					builder.WriteString("<")
				} else if cart.direction == EAST {
					builder.WriteString(">")
				}
				continue
			}

			c, ok := t[p]
			if !ok {
				break
			}

			builder.WriteString(c)
		}

		builder.WriteString("\n")
	}

	return builder.String()
}

type Cart struct {
	location  aoc.Point2D
	direction int
	lastTurn  int // the direction we last turned
}

func (c *Cart) String() string {
	return c.location.String()
}

func InputToTrackAndCarts(year, day int) (Track, []*Cart) {
	track := make(Track)
	carts := make([]*Cart, 0)
	for y, line := range aoc.InputToLines(year, day) {
		for x, c := range line {
			p := aoc.Point2D{X: x, Y: y}

			switch c {
			case '^':
				carts = append(carts, &Cart{location: p, direction: NORTH})
				track[p] = "|"
			case 'v':
				carts = append(carts, &Cart{location: p, direction: SOUTH})
				track[p] = "|"
			case '<':
				carts = append(carts, &Cart{location: p, direction: WEST})
				track[p] = "-"
			case '>':
				carts = append(carts, &Cart{location: p, direction: EAST})
				track[p] = "-"
			default:
				track[p] = string(c)
			}
		}
	}

	return track, carts
}
