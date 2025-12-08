package main

import (
	"fmt"
	"maps"
	"slices"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var ps []Point3D
	for in.HasNext() {
		ps = append(ps, Point3D{X: in.Int(), Y: in.Int(), Z: in.Int()})
	}

	var distances PriorityQueue[[2]Point3D]
	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			distances.Push([2]Point3D{ps[i], ps[j]}, Distance2(ps[i], ps[j]))
		}
	}

	var ds DisjointSet[Point3D]
	for range 1000 {
		pair := distances.Pop()
		ds.UnionWithAdd(pair[0], pair[1])
	}

	var sizes = make(map[Point3D]int)
	for _, p := range ps {
		id, _ := ds.Find(p)
		sizes[id] = ds.Size(id)
	}

	var largest = slices.Sorted(maps.Values(sizes))
	fmt.Println(Product(largest[len(largest)-3:]...))
}

func Distance2(p, q Point3D) int {
	dx, dy, dz := p.X-q.X, p.Y-q.Y, p.Z-q.Z
	return dx*dx + dy*dy + dz*dz
}
