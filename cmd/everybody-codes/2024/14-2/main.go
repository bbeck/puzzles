package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var seen Set[Point3D]
	for in.HasNext() {
		var p Point3D
		for _, instruction := range strings.Split(in.Line(), ",") {
			var dx, dy, dz int
			switch instruction[0] {
			case 'U':
				dy = 1
			case 'D':
				dy = -1
			case 'L':
				dx = 1
			case 'R':
				dx = -1
			case 'F':
				dz = 1
			case 'B':
				dz = -1
			}

			N := ParseInt(instruction[1:])
			for n := 0; n < N; n++ {
				p = Point3D{X: p.X + dx, Y: p.Y + dy, Z: p.Z + dz}
				seen.Add(p)
			}
		}
	}
	
	fmt.Println(len(seen))
}
