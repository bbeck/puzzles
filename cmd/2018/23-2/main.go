package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	bots := InputToNanobots(2018, 23)

	// We're asked to find the point that's contained with the bounding box of the
	// most bots.  For this we'll use a simulated annealing which is
	// probabilistic so it may take a lot of runs to get the optimal answer.  To
	// mitigate this we'll tune it to run slowly in hopes that it'll converge on
	// the global minimum.

	// A random neighbor of a given point.
	neighbor := func(p Point3D) Point3D {
		switch rand.Intn(6) {
		case 0:
			return Point3D{p.X - rand.Intn(10000), p.Y, p.Z}
		case 1:
			return Point3D{p.X + rand.Intn(10000), p.Y, p.Z}
		case 2:
			return Point3D{p.X, p.Y - rand.Intn(10000), p.Z}
		case 3:
			return Point3D{p.X, p.Y + rand.Intn(10000), p.Z}
		case 4:
			return Point3D{p.X, p.Y, p.Z - rand.Intn(10000)}
		case 5:
			return Point3D{p.X, p.Y, p.Z + rand.Intn(10000)}
		default:
			log.Fatal("unhandled case")
			return p
		}
	}

	// Evaluate the point that we're at to determine how good it is.  This is the
	// function that we're going to attempt to minimize -- so the better the
	// solution is, the lower of a value this should return.  For our purposes
	// returning the number of bots that this point isn't within range of should
	// do well.
	cost := func(p Point3D) int {
		var count int
		for _, bot := range bots {
			if !bot.InRange(p) {
				count++
			}
		}

		return count
	}

	// The acceptance function that determines whether or not we accept the
	// solution with the old cost or the solution with the new cost.  This is the
	// probabilistic part of simulated annealing.  We model acceptance by
	// comparing e^((old - new)/T) because when the new solution is better
	// (lower) than the old solution (old - new) will be positive and cause the
	// exponential to evaluate to a larger number.  When the old solution is
	// better then (old - new) will be negative and cause the exponential to be
	// close to zero.  Thus we'll tend to accept better answers over worse ones
	// depending on the random number generated.
	accept := func(old, new int, t float64) bool {
		return math.Exp(float64(old-new)/t) > rand.Float64()
	}

	// We need to start with a solution, for this we'll pick one the location of
	// one of the bots at random.
	bestLocation := bots[rand.Intn(len(bots))].location
	bestCost := cost(bestLocation)

	// We'll start at a hot temperature and after each iteration scale it by
	// alpha until we reach our minimum temperature.  At that point we're done.
	Tmax := 100000.
	Tmin := 0.00001
	alpha := 0.9

	for T := Tmax; T > Tmin; T = T * alpha {
		// Try a lot of neighbors in the area of our best solution in the hopes of
		// finding the best one.
		for i := 0; i < 10000; i++ {
			nextLocation := neighbor(bestLocation)
			nextCost := cost(nextLocation)

			if accept(bestCost, nextCost, T) {
				bestLocation = nextLocation
				bestCost = nextCost
			}
		}
	}

	distance := bestLocation.X + bestLocation.Y + bestLocation.Z
	fmt.Printf("solution: %s, cost: %d\n", bestLocation, bestCost)
	fmt.Printf("manhattan distance from origin: %d\n", distance)
}

type Nanobot struct {
	location Point3D
	radius   int
}

// Determine if a point is in range of this nanobot.
func (n Nanobot) InRange(p Point3D) bool {
	dx := n.location.X - p.X
	if dx < 0 {
		dx = -dx
	}
	dy := n.location.Y - p.Y
	if dy < 0 {
		dy = -dy
	}
	dz := n.location.Z - p.Z
	if dz < 0 {
		dz = -dz
	}

	return dx+dy+dz <= n.radius
}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.X, p.Y, p.Z)
}

func InputToNanobots(year, day int) []Nanobot {
	var nanobots []Nanobot
	for _, line := range aoc.InputToLines(year, day) {
		var x, y, z, r int
		if _, err := fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r); err != nil {
			log.Fatalf("unable to parse line: %s", line)
		}

		nanobots = append(nanobots, Nanobot{Point3D{x, y, z}, r})
	}

	return nanobots
}
