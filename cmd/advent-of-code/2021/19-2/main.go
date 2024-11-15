package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	scanners := InputToScanners()

	transforms := puz.Make2D[func(puz.Point3D) puz.Point3D](len(scanners), len(scanners))
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
	D := puz.Make2D[int](len(scanners), len(scanners))
	N := puz.Make2D[int](len(scanners), len(scanners))

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

	// Transform each scanner's location into a single coordinate system
	var ps []puz.Point3D
	for i := 0; i < len(scanners); i++ {
		p := puz.Origin3D

		prev := i
		for prev != 0 {
			next := N[prev][0]
			p = transforms[prev][next](p)
			prev = next
		}
		ps = append(ps, p)
	}

	var max int
	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			max = puz.Max(max, ps[i].ManhattanDistance(ps[j]))
		}
	}
	fmt.Println(max)
}

func Overlaps(a, b Scanner) bool {
	return len(a.AllDistances.Intersect(b.AllDistances)) >= 66
}

// Transform builds a transformation function to convert the beacons from a's
// coordinate system to b's coordinate system.
func Transform(a, b Scanner) func(puz.Point3D) puz.Point3D {
	// First, determine any two points that are the same in both scanners.  We'll
	// use this later to determine what translation is necessary.
	var pa, pb puz.Point3D
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
	var transform func(puz.Point3D) puz.Point3D

	bBeacons := puz.SetFrom(b.Beacons...)
	for _, rotate := range Rotations {
		paRotated := rotate(pa)

		// Use our same points to determine translation
		dx, dy, dz := pb.X-paRotated.X, pb.Y-paRotated.Y, pb.Z-paRotated.Z
		translate := func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: p.X + dx, Y: p.Y + dy, Z: p.Z + dz} }

		aBeacons := puz.SetFrom(Map(Map(a.Beacons, rotate), translate)...)
		if len(aBeacons.Intersect(bBeacons)) >= 12 {
			transform = func(p puz.Point3D) puz.Point3D { return translate(rotate(p)) }
			break
		}
	}

	return transform
}

var Rotations = []func(puz.Point3D) puz.Point3D{
	// http://www.euclideanspace.com/maths/algebra/matrix/transforms/examples/index.htm
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: +1 * p.X, Y: +1 * p.Y, Z: +1 * p.Z} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: +1 * p.X, Y: +1 * p.Z, Z: -1 * p.Y} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: +1 * p.X, Y: -1 * p.Y, Z: -1 * p.Z} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: +1 * p.X, Y: -1 * p.Z, Z: +1 * p.Y} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: +1 * p.Y, Y: -1 * p.X, Z: +1 * p.Z} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: +1 * p.Y, Y: +1 * p.Z, Z: +1 * p.X} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: +1 * p.Y, Y: +1 * p.X, Z: -1 * p.Z} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: +1 * p.Y, Y: -1 * p.Z, Z: -1 * p.X} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: -1 * p.X, Y: -1 * p.Y, Z: +1 * p.Z} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: -1 * p.X, Y: -1 * p.Z, Z: -1 * p.Y} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: -1 * p.X, Y: +1 * p.Y, Z: -1 * p.Z} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: -1 * p.X, Y: +1 * p.Z, Z: +1 * p.Y} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: -1 * p.Y, Y: +1 * p.X, Z: +1 * p.Z} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: -1 * p.Y, Y: -1 * p.Z, Z: +1 * p.X} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: -1 * p.Y, Y: -1 * p.X, Z: -1 * p.Z} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: -1 * p.Y, Y: +1 * p.Z, Z: -1 * p.X} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: +1 * p.Z, Y: +1 * p.Y, Z: -1 * p.X} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: +1 * p.Z, Y: +1 * p.X, Z: +1 * p.Y} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: +1 * p.Z, Y: -1 * p.Y, Z: +1 * p.X} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: +1 * p.Z, Y: -1 * p.X, Z: -1 * p.Y} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: -1 * p.Z, Y: -1 * p.Y, Z: -1 * p.X} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: -1 * p.Z, Y: -1 * p.X, Z: +1 * p.Y} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: -1 * p.Z, Y: +1 * p.Y, Z: +1 * p.X} },
	func(p puz.Point3D) puz.Point3D { return puz.Point3D{X: -1 * p.Z, Y: +1 * p.X, Z: -1 * p.Y} },
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
	Beacons      []puz.Point3D
	AllDistances puz.Set[int]
	Distances    []puz.Set[int]
}

func (s *Scanner) AddBeacon(p puz.Point3D) {
	var distances puz.Set[int]
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
	for _, line := range puz.InputToLines() {
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

		var beacon puz.Point3D
		fmt.Sscanf(line, "%d,%d,%d", &beacon.X, &beacon.Y, &beacon.Z)
		current.AddBeacon(beacon)
	}
	scanners = append(scanners, current)

	return scanners
}
