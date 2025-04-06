package main

import (
	"container/heap"
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
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

	var cube Cube3D
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

	fmt.Println(Origin3D.ManhattanDistance(cube.Point3D))
}

func GetInitialCube(bots []Nanobot) Cube3D {
	// Ensure that the dimensions of the cube returned are a power of two so as
	// we subdivide we always end up with integer dimensions.
	max := 1
	for _, b := range bots {
		for max < Abs(b.X)+b.R || max < Abs(b.Y)+b.R || max < Abs(b.Z)+b.R {
			max *= 2
		}
	}

	return Cube3D{
		Point3D: Point3D{X: -max, Y: -max, Z: -max},
		W:       2 * max, H: 2 * max, D: 2 * max,
	}
}

type Cube3D struct {
	Point3D
	W, H, D int
}

func (c Cube3D) Volume() int {
	return c.W * c.H * c.D
}

func (c Cube3D) Octants() [8]Cube3D {
	x, y, z := c.X, c.Y, c.Z
	dx, dy, dz := c.W/2, c.H/2, c.D/2

	return [8]Cube3D{
		{Point3D: Point3D{X: x, Y: y, Z: z}, W: dx, H: dy, D: dz},
		{Point3D: Point3D{X: x + dx, Y: y, Z: z}, W: dx, H: dy, D: dz},
		{Point3D: Point3D{X: x, Y: y + dy, Z: z}, W: dx, H: dy, D: dz},
		{Point3D: Point3D{X: x, Y: y, Z: z + dz}, W: dx, H: dy, D: dz},
		{Point3D: Point3D{X: x + dx, Y: y + dy, Z: z}, W: dx, H: dy, D: dz},
		{Point3D: Point3D{X: x + dx, Y: y, Z: z + dz}, W: dx, H: dy, D: dz},
		{Point3D: Point3D{X: x, Y: y + dy, Z: z + dz}, W: dx, H: dy, D: dz},
		{Point3D: Point3D{X: x + dx, Y: y + dy, Z: z + dz}, W: dx, H: dy, D: dz},
	}
}

func (c Cube3D) GetNumIntersectingBots(bots []Nanobot) int {
	var count int
	for _, b := range bots {
		closest := Point3D{
			X: Min(Max(b.X, c.X), c.X+c.W-1),
			Y: Min(Max(b.Y, c.Y), c.Y+c.H-1),
			Z: Min(Max(b.Z, c.Z), c.Z+c.D-1),
		}

		if closest.ManhattanDistance(b.Point3D) <= b.R {
			count++
		}
	}

	return count
}

type Entry struct {
	Cube  Cube3D
	Count int
}

type Heap []Entry

func (h Heap) Len() int      { return len(h) }
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool {
	if h[i].Count != h[j].Count {
		return h[i].Count > h[j].Count
	}

	di := Origin3D.ManhattanDistance(h[i].Cube.Point3D)
	dj := Origin3D.ManhattanDistance(h[j].Cube.Point3D)
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
	Point3D
	R int
}

func InputToNanobots() []Nanobot {
	return in.LinesToS[Nanobot](func(s in.Scanner[Nanobot]) Nanobot {
		return Nanobot{
			Point3D: Point3D{X: in.Int(), Y: in.Int(), Z: in.Int()},
			R:       in.Int(),
		}
	})
}
