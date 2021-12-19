package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	scanners := InputToScanners()
	fmt.Println("num scanners:", len(scanners))

outer:
	for len(scanners) > 1 {
		fmt.Println("there are", len(scanners), "scanners remaining")
		for i := 0; i < len(scanners); i++ {
			for j := i + 1; j < len(scanners); j++ {
				if ok, transform, correction := Overlap(scanners[i], scanners[j]); ok {
					// We've found two scanners that overlap.  Merge the beacons from scanner j into
					// scanner i, being sure to translate them to scanner i's coordinates first.
					for _, e := range scanners[j].Beacons.Entries() {
						b := e.(Point3D)
						scanners[i].Beacons.Add(correction(transform(b)))
					}

					for _, l := range scanners[j].ScannerLocations {
						scanners[i].ScannerLocations = append(scanners[i].ScannerLocations, correction(transform(l)))
					}

					// Remove scanner j
					scanners[j] = scanners[len(scanners)-1]
					scanners = scanners[:len(scanners)-1]
					continue outer
				}
			}
		}
	}

	fmt.Println("# beacons:", scanners[0].Beacons.Size())

	locations := scanners[0].ScannerLocations
	locations = append(locations, Point3D{})

	fmt.Println("locations:")
	for _, l := range scanners[0].ScannerLocations {
		fmt.Println("  ", l)
	}

	var max int
	for i := 0; i < len(locations); i++ {
		for j := i + 1; j < len(locations); j++ {
			max = aoc.MaxInt(max, locations[i].ManhattanDistance(locations[j]))
		}
	}
	fmt.Println(max)
}

func Overlap(a, b Scanner) (bool, Transform, Transform) {
	for _, transform := range Transforms {
		bBeacons := b.GetTransformedBeacons(transform)

		// Give every pair of beacons (one from A and one from B) the chance to be the same
		// beacon in different coordinate systems.  This will cause us to generate a
		// correction factor for B's beacons.  If the beacons are truly the same, then at
		// least 12 other beacons in B should become beacons of A.
		for _, aEntry := range a.Beacons.Entries() {
			aBeacon := aEntry.(Point3D)

			for _, bEntry := range bBeacons.Entries() {
				bBeacon := bEntry.(Point3D)
				dx := bBeacon.X - aBeacon.X
				dy := bBeacon.Y - aBeacon.Y
				dz := bBeacon.Z - aBeacon.Z
				correction := func(p Point3D) Point3D {
					return Point3D{p.X - dx, p.Y - dy, p.Z - dz}
				}

				matches := CountMatches(a.Beacons, bBeacons, correction)
				if matches >= 12 {
					return true, transform, correction
				}
			}
		}
	}
	return false, nil, nil
}

func CountMatches(as, bs aoc.Set, transform Transform) int {
	bSet := aoc.NewSet()
	for _, b := range bs.Entries() {
		bSet.Add(transform(b.(Point3D)))
	}

	return as.Intersect(bSet).Size()
}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) ManhattanDistance(q Point3D) int {
	dx := aoc.AbsInt(p.X - q.X)
	dy := aoc.AbsInt(p.Y - q.Y)
	dz := aoc.AbsInt(p.Z - q.Z)

	return dx + dy + dz
}

type Scanner struct {
	ID               string
	Beacons          aoc.Set
	ScannerLocations []Point3D
}

type Transform func(Point3D) Point3D

var Transforms = []Transform{
	// http://www.euclideanspace.com/maths/algebra/matrix/transforms/examples/index.htm
	func(p Point3D) Point3D { return Point3D{p.X, p.Y, p.Z} },
	func(p Point3D) Point3D { return Point3D{p.X, p.Z, -p.Y} },
	func(p Point3D) Point3D { return Point3D{p.X, -p.Y, -p.Z} },
	func(p Point3D) Point3D { return Point3D{p.X, -p.Z, p.Y} },
	func(p Point3D) Point3D { return Point3D{p.Y, -p.X, p.Z} },
	func(p Point3D) Point3D { return Point3D{p.Y, p.Z, p.X} },
	func(p Point3D) Point3D { return Point3D{p.Y, p.X, -p.Z} },
	func(p Point3D) Point3D { return Point3D{p.Y, -p.Z, -p.X} },
	func(p Point3D) Point3D { return Point3D{-p.X, -p.Y, p.Z} },
	func(p Point3D) Point3D { return Point3D{-p.X, -p.Z, -p.Y} },
	func(p Point3D) Point3D { return Point3D{-p.X, p.Y, -p.Z} },
	func(p Point3D) Point3D { return Point3D{-p.X, p.Z, p.Y} },
	func(p Point3D) Point3D { return Point3D{-p.Y, p.X, p.Z} },
	func(p Point3D) Point3D { return Point3D{-p.Y, -p.Z, p.X} },
	func(p Point3D) Point3D { return Point3D{-p.Y, -p.X, -p.Z} },
	func(p Point3D) Point3D { return Point3D{-p.Y, p.Z, -p.X} },
	func(p Point3D) Point3D { return Point3D{p.Z, p.Y, -p.X} },
	func(p Point3D) Point3D { return Point3D{p.Z, p.X, p.Y} },
	func(p Point3D) Point3D { return Point3D{p.Z, -p.Y, p.X} },
	func(p Point3D) Point3D { return Point3D{p.Z, -p.X, -p.Y} },
	func(p Point3D) Point3D { return Point3D{-p.Z, -p.Y, -p.X} },
	func(p Point3D) Point3D { return Point3D{-p.Z, -p.X, p.Y} },
	func(p Point3D) Point3D { return Point3D{-p.Z, p.Y, p.X} },
	func(p Point3D) Point3D { return Point3D{-p.Z, p.X, -p.Y} },
}

func (s Scanner) GetTransformedBeacons(t Transform) aoc.Set {
	bs := aoc.NewSet()
	for _, b := range s.Beacons.Entries() {
		bs.Add(t(b.(Point3D)))
	}
	return bs
}

func InputToScanners() []Scanner {
	var scanners []Scanner

	var scanner = Scanner{Beacons: aoc.NewSet()}
	for _, line := range aoc.InputToLines(2021, 19) {
		var id string
		if _, err := fmt.Sscanf(line, "--- scanner %s ---", &id); err == nil {
			scanner.ID = id
			scanners = append(scanners, scanner)
			continue
		}

		if line == "" {
			scanner = Scanner{Beacons: aoc.NewSet(), ScannerLocations: []Point3D{{}}}
			continue
		}

		var x, y, z int
		if _, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z); err == nil {
			scanner.Beacons.Add(Point3D{X: x, Y: y, Z: z})
			continue
		}
	}

	return scanners
}
