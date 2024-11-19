package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	scanners := InputToScanners()

	transforms := lib.Make2D[func(lib.Point3D) lib.Point3D](len(scanners), len(scanners))
	for i := 0; i < len(scanners); i++ {
		for j := i + 1; j < len(scanners); j++ {
			if Overlaps(scanners[i], scanners[j]) {
				transforms[i][j] = Transform(scanners[i], scanners[j])
				transforms[j][i] = Transform(scanners[j], scanners[i])
			}
		}
	}

	// Use the Floyd-Warshall algorithm to determine the transform between all
	// pairs of scanners.
	D := lib.Make2D[int](len(scanners), len(scanners))
	N := lib.Make2D[int](len(scanners), len(scanners))

	for i := 0; i < len(scanners); i++ {
		for j := 0; j < len(scanners); j++ {
			if transforms[i][j] != nil {
				D[i][j] = 1
				N[i][j] = j
			} else {
				D[i][j] = 1e6
				N[i][j] = -1
			}
		}
		D[i][i] = 0
		N[i][i] = i
	}

	for k := 0; k < len(scanners); k++ {
		for i := 0; i < len(scanners); i++ {
			for j := 0; j < len(scanners); j++ {
				if D[i][j] > D[i][k]+D[k][j] {
					D[i][j] = D[i][k] + D[k][j]
					N[i][j] = N[i][k]
				}
			}
		}
	}

	// Transform the beacons for each scanner into the same coordinate system.
	var beacons lib.Set[lib.Point3D]
	for i := 0; i < len(scanners); i++ {
		bs := scanners[i].Beacons

		prev := i
		for prev != 0 {
			next := N[prev][0]
			bs = Map(bs, transforms[prev][next])
			prev = next
		}
		beacons.Add(bs...)
	}
	fmt.Println(len(beacons))
}

func Overlaps(a, b Scanner) bool {
	return len(a.AllDistances.Intersect(b.AllDistances)) >= 66
}

// Transform builds a transformation function to convert the beacons from a's
// coordinate system to b's coordinate system.
func Transform(a, b Scanner) func(lib.Point3D) lib.Point3D {
	// First, determine any two points that are the same in both scanners.  We'll
	// use this later to determine what translation is necessary.
	var pa, pb lib.Point3D
outer:
	for i := 0; i < len(a.Beacons); i++ {
		for j := 0; j < len(b.Beacons); j++ {
			// Only 11 beacons need to overlap since we're forcing the ith and jth
			// beacons to be the same.
			if len(a.Distances[i].Intersect(b.Distances[j])) >= 11 {
				pa, pb = a.Beacons[i], b.Beacons[j]
				break outer
			}
		}
	}

	// Next, try each possible rotation.  We'll rotate each of the points in a by
	// the rotation, then use our two points that are known to be the same to
	// determine the needed translation.
	var transform func(lib.Point3D) lib.Point3D

	bBeacons := lib.SetFrom(b.Beacons...)
	for _, rotate := range Rotations {
		paRotated := rotate(pa)

		// Use our same points to determine translation
		dx, dy, dz := pb.X-paRotated.X, pb.Y-paRotated.Y, pb.Z-paRotated.Z
		translate := func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: p.X + dx, Y: p.Y + dy, Z: p.Z + dz} }

		aBeacons := lib.SetFrom(Map(Map(a.Beacons, rotate), translate)...)
		if len(aBeacons.Intersect(bBeacons)) >= 12 {
			transform = func(p lib.Point3D) lib.Point3D { return translate(rotate(p)) }
			break
		}
	}

	return transform
}

var Rotations = []func(lib.Point3D) lib.Point3D{
	// http://www.euclideanspace.com/maths/algebra/matrix/transforms/examples/index.htm
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: +1 * p.X, Y: +1 * p.Y, Z: +1 * p.Z} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: +1 * p.X, Y: +1 * p.Z, Z: -1 * p.Y} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: +1 * p.X, Y: -1 * p.Y, Z: -1 * p.Z} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: +1 * p.X, Y: -1 * p.Z, Z: +1 * p.Y} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: +1 * p.Y, Y: -1 * p.X, Z: +1 * p.Z} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: +1 * p.Y, Y: +1 * p.Z, Z: +1 * p.X} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: +1 * p.Y, Y: +1 * p.X, Z: -1 * p.Z} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: +1 * p.Y, Y: -1 * p.Z, Z: -1 * p.X} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: -1 * p.X, Y: -1 * p.Y, Z: +1 * p.Z} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: -1 * p.X, Y: -1 * p.Z, Z: -1 * p.Y} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: -1 * p.X, Y: +1 * p.Y, Z: -1 * p.Z} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: -1 * p.X, Y: +1 * p.Z, Z: +1 * p.Y} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: -1 * p.Y, Y: +1 * p.X, Z: +1 * p.Z} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: -1 * p.Y, Y: -1 * p.Z, Z: +1 * p.X} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: -1 * p.Y, Y: -1 * p.X, Z: -1 * p.Z} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: -1 * p.Y, Y: +1 * p.Z, Z: -1 * p.X} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: +1 * p.Z, Y: +1 * p.Y, Z: -1 * p.X} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: +1 * p.Z, Y: +1 * p.X, Z: +1 * p.Y} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: +1 * p.Z, Y: -1 * p.Y, Z: +1 * p.X} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: +1 * p.Z, Y: -1 * p.X, Z: -1 * p.Y} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: -1 * p.Z, Y: -1 * p.Y, Z: -1 * p.X} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: -1 * p.Z, Y: -1 * p.X, Z: +1 * p.Y} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: -1 * p.Z, Y: +1 * p.Y, Z: +1 * p.X} },
	func(p lib.Point3D) lib.Point3D { return lib.Point3D{X: -1 * p.Z, Y: +1 * p.X, Z: -1 * p.Y} },
}

func Map[T, V any](inputs []T, fn func(T) V) []V {
	outputs := make([]V, len(inputs))
	for i, in := range inputs {
		outputs[i] = fn(in)
	}
	return outputs
}

type Scanner struct {
	ID           string
	Beacons      []lib.Point3D
	AllDistances lib.Set[int]
	Distances    []lib.Set[int]
}

func (s *Scanner) AddBeacon(p lib.Point3D) {
	var distances lib.Set[int]
	for i, b := range s.Beacons {
		d := p.ManhattanDistance(b)
		s.AllDistances.Add(d)
		s.Distances[i].Add(d)
		distances.Add(d)
	}

	s.Beacons = append(s.Beacons, p)
	s.Distances = append(s.Distances, distances)
}

func InputToScanners() []Scanner {
	var scanners []Scanner

	var current Scanner
	for _, line := range lib.InputToLines() {
		if line == "" {
			if current.ID != "" {
				scanners = append(scanners, current)
			}

			current = Scanner{}
			continue
		}

		if strings.HasPrefix(line, "--") {
			current.ID = line
			continue
		}

		var beacon lib.Point3D
		fmt.Sscanf(line, "%d,%d,%d", &beacon.X, &beacon.Y, &beacon.Z)
		current.AddBeacon(beacon)
	}
	scanners = append(scanners, current)

	return scanners
}
