package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	world := InputToWorld(2018, 17)
	Flow(world, aoc.Point2D{X: 500, Y: 1})

	fmt.Printf("complete\n")
	fmt.Println(world)

	var count int
	for p, cell := range world.grid {
		if p.Y < world.minY || p.Y > world.maxY {
			continue
		}

		if cell == WATER || cell == FLOW {
			count++
		}
	}
	fmt.Printf("tiles that water can reach: %d\n", count)
}

func Flow(world *World, p aoc.Point2D) {
	if p.Y > world.maxY {
		return
	}

	current := world.grid[p]
	if current != EMPTY {
		return
	}

	world.grid[p] = FLOW

	below := world.grid[p.Down()]
	if below == EMPTY {
		Flow(world, p.Down())
	}

	// At this point everything below us has been flooded properly.  If we're
	// currently on top of more flow then we're done we can't move to the sides.
	// Otherwise we're on top of standing water or a wall and need to spread water
	// to the sides.
	below = world.grid[p.Down()]
	if below == FLOW || p.Down().Y > world.maxY {
		return
	}

	Spread(world, p)
}

func Spread(world *World, p aoc.Point2D) bool {
	if world.grid[p] != FLOW {
		log.Fatalf("attempting to spread on a non-flow cell, p: %s", p)
		return false
	}

	// We can't short circuit here, we must always try to spread in both
	// directions.
	containedLeft := SpreadLeft(world, p.Left())
	containedRight := SpreadRight(world, p.Right())

	// If the left and right are contained, we can convert all of our flow into
	// standing water.
	if containedLeft && containedRight {
		for x := p; world.grid[x] != WALL; x = x.Left() {
			world.grid[x] = WATER
		}
		for x := p; world.grid[x] != WALL; x = x.Right() {
			world.grid[x] = WATER
		}

		return true
	}

	return false
}

// SpreadLeft will keep spreading flow to the left until it either falls off of
// a cliff or runs into a way.  The returned boolean indicates if a wall was hit
// which indicates that from this side it's okay to convert the flow into
// standing water since it's contained.
func SpreadLeft(world *World, p aoc.Point2D) bool {
	if world.grid[p] == WALL || world.grid[p] == WATER {
		return true
	}

	below := world.grid[p.Down()]
	if below == EMPTY {
		Flow(world, p)
		return false
	}

	if below == FLOW {
		return false
	}

	world.grid[p] = FLOW
	return SpreadLeft(world, p.Left())
}

func SpreadRight(world *World, p aoc.Point2D) bool {
	if world.grid[p] == WALL || world.grid[p] == WATER {
		return true
	}

	below := world.grid[p.Down()]
	if below == EMPTY {
		Flow(world, p)
		return false
	}

	if below == FLOW {
		return false
	}

	world.grid[p] = FLOW
	return SpreadRight(world, p.Right())
}

const (
	EMPTY  int = 0
	SPRING int = 1
	WALL   int = 3
	FLOW   int = 4
	WATER  int = 5
)

type World struct {
	minX, maxX int
	minY, maxY int
	grid       map[aoc.Point2D]int
}

func (w *World) String() string {
	var builder strings.Builder

	// header
	builder.WriteString("     ")
	for x := w.minX - 1; x <= w.maxX+1; x++ {
		builder.WriteString(fmt.Sprintf("%d", x/100%10))
	}
	builder.WriteString("\n")
	builder.WriteString("     ")
	for x := w.minX - 1; x <= w.maxX+1; x++ {
		builder.WriteString(fmt.Sprintf("%d", x/10%10))
	}
	builder.WriteString("\n")
	builder.WriteString("     ")
	for x := w.minX - 1; x <= w.maxX+1; x++ {
		builder.WriteString(fmt.Sprintf("%d", x%10))
	}
	builder.WriteString("\n")

	for y := w.minY; y <= w.maxY; y++ {
		builder.WriteString(fmt.Sprintf("%4d ", y))
		for x := w.minX - 1; x <= w.maxX+1; x++ {
			p := aoc.Point2D{X: x, Y: y}
			switch w.grid[p] {
			case SPRING:
				builder.WriteString("+")
			case WALL:
				builder.WriteString("█")
			case FLOW:
				builder.WriteString("|")
			case WATER:
				builder.WriteString("~")
			default:
				builder.WriteString(" ")
			}
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func InputToWorld(year, day int) *World {
	minX, minY := math.MaxInt64, math.MaxInt64
	maxX, maxY := math.MinInt64, math.MinInt64
	grid := map[aoc.Point2D]int{}

	for _, line := range aoc.InputToLines(year, day) {
		var xmin, xmax, ymin, ymax int
		if _, err := fmt.Sscanf(line, "x=%d, y=%d..%d", &xmin, &ymin, &ymax); err == nil {
			xmax = xmin
		} else if _, err := fmt.Sscanf(line, "y=%d, x=%d..%d", &ymin, &xmin, &xmax); err == nil {
			ymax = ymin
		} else {
			log.Fatalf("unable to parse input: %s", line)
		}

		if xmin < minX {
			minX = xmin
		}
		if xmax > maxX {
			maxX = xmax
		}
		if ymin < minY {
			minY = ymin
		}
		if ymax > maxY {
			maxY = ymax
		}

		for y := ymin; y <= ymax; y++ {
			for x := xmin; x <= xmax; x++ {
				p := aoc.Point2D{X: x, Y: y}
				grid[p] = WALL
			}
		}
	}

	return &World{
		minX: minX,
		maxX: maxX,
		minY: minY,
		maxY: maxY,
		grid: grid,
	}
}
