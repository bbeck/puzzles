package main

import (
	"container/heap"
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	bots := InputToNanobots()

	// We're going to work on iteratively subdividing the space into smaller
	// cubes constantly focusing on the cube that intersects with the largest
	// number of bots.  When multiple cubes intersect the same number of bots
	// then the focus will be on the one closest to the origin.  Once we've
	// whittled the space down to a single point we know we're finished.
	//
	// We use a container/heap instead of an aoc.PriorityQueue[T] because we
	// need a compound key that doesn't fit easily into the singular priority
	// that aoc.PriorityQueue[T] uses.

	cubes := &Heap{
		Entry{
			Cube:  GetInitialCube(bots),
			Count: len(bots),
		},
	}
	heap.Init(cubes)

	var cube Cube
	for len(*cubes) > 0 {
		cube = heap.Pop(cubes).(Entry).Cube

		if cube.Volume() == 1 {
			break
		}

		for _, octant := range cube.Octants() {
			entry := Entry{
				Cube:  octant,
				Count: octant.GetNumIntersectingBots(bots),
			}
			heap.Push(cubes, entry)
		}
	}

	fmt.Println(aoc.Origin3D.ManhattanDistance(cube.Point3D))
}

func GetInitialCube(bots []Nanobot) Cube {
	// Ensure that the dimensions of the cube returned are a power of two so as
	// we subdivide we always end up with integer dimensions.
	max := 1
	for _, b := range bots {
		for max < aoc.Abs(b.X)+b.R || max < aoc.Abs(b.Y)+b.R || max < aoc.Abs(b.Z)+b.R {
			max *= 2
		}
	}

	return Cube{
		Point3D: aoc.Point3D{X: -max, Y: -max, Z: -max},
		W:       2 * max, H: 2 * max, D: 2 * max,
	}
}

type Cube struct {
	aoc.Point3D
	W, H, D int
}

func (c Cube) Volume() int {
	return c.W * c.H * c.D
}

func (c Cube) Octants() [8]Cube {
	x, y, z := c.X, c.Y, c.Z
	dx, dy, dz := c.W/2, c.H/2, c.D/2

	return [8]Cube{
		{Point3D: aoc.Point3D{X: x, Y: y, Z: z}, W: dx, H: dy, D: dz},
		{Point3D: aoc.Point3D{X: x + dx, Y: y, Z: z}, W: dx, H: dy, D: dz},
		{Point3D: aoc.Point3D{X: x, Y: y + dy, Z: z}, W: dx, H: dy, D: dz},
		{Point3D: aoc.Point3D{X: x, Y: y, Z: z + dz}, W: dx, H: dy, D: dz},
		{Point3D: aoc.Point3D{X: x + dx, Y: y + dy, Z: z}, W: dx, H: dy, D: dz},
		{Point3D: aoc.Point3D{X: x + dx, Y: y, Z: z + dz}, W: dx, H: dy, D: dz},
		{Point3D: aoc.Point3D{X: x, Y: y + dy, Z: z + dz}, W: dx, H: dy, D: dz},
		{Point3D: aoc.Point3D{X: x + dx, Y: y + dy, Z: z + dz}, W: dx, H: dy, D: dz},
	}
}

func (c Cube) GetNumIntersectingBots(bots []Nanobot) int {
	var count int
	for _, b := range bots {
		closest := aoc.Point3D{
			X: aoc.Min(aoc.Max(b.X, c.X), c.X+c.W-1),
			Y: aoc.Min(aoc.Max(b.Y, c.Y), c.Y+c.H-1),
			Z: aoc.Min(aoc.Max(b.Z, c.Z), c.Z+c.D-1),
		}

		if closest.ManhattanDistance(b.Point3D) <= b.R {
			count++
		}
	}

	return count
}

type Entry struct {
	Cube  Cube
	Count int
}

type Heap []Entry

func (h Heap) Len() int      { return len(h) }
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool {
	if h[i].Count != h[j].Count {
		return h[i].Count > h[j].Count
	}

	di := aoc.Origin3D.ManhattanDistance(h[i].Cube.Point3D)
	dj := aoc.Origin3D.ManhattanDistance(h[j].Cube.Point3D)
	return di < dj
}

func (h *Heap) Push(x any) { *h = append(*h, x.(Entry)) }
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Nanobot struct {
	aoc.Point3D
	R int
}

func InputToNanobots() []Nanobot {
	return aoc.InputLinesTo(2018, 23, func(line string) (Nanobot, error) {
		var p aoc.Point3D
		var r int
		_, err := fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &p.X, &p.Y, &p.Z, &r)
		return Nanobot{Point3D: p, R: r}, err
	})
}
