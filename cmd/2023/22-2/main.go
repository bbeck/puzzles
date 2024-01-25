package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"math"
	"sort"
)

func main() {
	bricks := InputToBricks()

	// Order the bricks so the ones closest to the ground are first.
	sort.Slice(bricks, func(i, j int) bool {
		iz := aoc.Min(bricks[i].MinZ, bricks[i].MaxZ)
		jz := aoc.Min(bricks[j].MinZ, bricks[j].MaxZ)
		return iz < jz
	})

	// Drop the bricks, building the volume containing all bricks.
	highest := aoc.Make2D[int](10, 10)
	volume := aoc.Make3D[int](10, 10, 400)
	for i := range bricks {
		// Determine how much this brick can drop.
		dz := math.MaxInt
		for x := bricks[i].MinX; x <= bricks[i].MaxX; x++ {
			for y := bricks[i].MinY; y <= bricks[i].MaxY; y++ {
				dz = aoc.Max(0, aoc.Min(dz, bricks[i].MinZ-highest[x][y]-1))
			}
		}

		bricks[i].MinZ -= dz
		bricks[i].MaxZ -= dz

		for x := bricks[i].MinX; x <= bricks[i].MaxX; x++ {
			for y := bricks[i].MinY; y <= bricks[i].MaxY; y++ {
				highest[x][y] = bricks[i].MaxZ

				for z := bricks[i].MinZ; z <= bricks[i].MaxZ; z++ {
					volume[x][y][z] = bricks[i].ID
				}
			}
		}
	}

	// Determine which blocks are supporting each other.
	supporting := make(map[int]aoc.Set[int])
	supportedBy := make(map[int]aoc.Set[int])
	for _, b := range bricks {
		for x := b.MinX; x <= b.MaxX; x++ {
			for y := b.MinY; y <= b.MaxY; y++ {
				if oid := volume[x][y][b.MaxZ+1]; oid != 0 {
					supporting[b.ID] = supporting[b.ID].UnionElems(oid)
					supportedBy[oid] = supportedBy[oid].UnionElems(b.ID)
				}
			}
		}
	}

	var sum int
	for _, b := range bricks {
		sum += Count(b.ID, supporting, supportedBy)
	}
	fmt.Println(sum)
}

func Count(n int, supporting, supportedBy map[int]aoc.Set[int]) int {
	var removed aoc.Set[int]

	toRemove := aoc.DequeFrom(n)
	for !toRemove.Empty() {
		id := toRemove.PopFront()
		removed.Add(id)

		for oid := range supporting[id].Difference(removed) {
			s := supportedBy[oid].Difference(removed)
			if len(s) == 0 {
				toRemove.PushBack(oid)
			}
		}
	}
	return len(removed) - 1
}

type Brick struct {
	aoc.Cube
	ID int
}

func InputToBricks() []Brick {
	var id int
	return aoc.InputLinesTo(2023, 22, func(line string) Brick {
		id++

		var c aoc.Cube
		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &c.MinX, &c.MinY, &c.MinZ, &c.MaxX, &c.MaxY, &c.MaxZ)
		return Brick{c, id}
	})
}
