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
					supporting[b.ID] = supporting[b.ID].Union(aoc.SetFrom(oid))
					supportedBy[oid] = supportedBy[oid].Union(aoc.SetFrom(b.ID))
				}
			}
		}
	}

	// Determine how many bricks can be removed without causing other bricks to
	// fall down.  A brick can be removed if the bricks its supporting are
	// supported by another brick.
	var count int
outer:
	for _, b := range bricks {
		for oid := range supporting[b.ID] {
			if len(supportedBy[oid]) == 1 {
				continue outer
			}
		}
		count++
	}
	fmt.Println(count)
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
