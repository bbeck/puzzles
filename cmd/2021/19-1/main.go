package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"log"
)

func main() {
	scanners := InputToScanners()

	// Determine which scanners overlap.  The overlapping scanners will form a graph
	// that can be used to determine the path to take to convert coordinates of the
	// beacons from any scanner into the coordinate system of any other scanner.
	edges := NewMatrix(len(scanners))
	for i := 0; i < len(scanners); i++ {
		for j := i + 1; j < len(scanners); j++ {
			if ok, iBeacon, jBeacon := Overlaps(scanners[i], scanners[j]); ok {
				edges[i][j] = BuildTransform(scanners[i].Beacons, scanners[j].Beacons, iBeacon, jBeacon)
				edges[j][i] = BuildTransform(scanners[j].Beacons, scanners[i].Beacons, jBeacon, iBeacon)
			}
		}
	}

	// Now that we have the graph, determine the path we have to take to convert each
	// scanner's beacons into the coordinate system of the first scanner and use it to
	// build a transform.
	transforms := make([]Transform, len(scanners))
	transforms[0] = func(p Point3D) Point3D { return p }

	isGoal := func(node aoc.Node) bool {
		return node.(*ScannerNode).id == 0
	}
	for i := 1; i < len(scanners); i++ {
		path, ok := aoc.BreadthFirstSearch(&ScannerNode{id: i, edges: edges}, isGoal)
		if !ok {
			log.Fatalf("unable to determine path to scanner 0 for scanner %d", i)
		}

		var funcs []Transform
		for j := 0; j < len(path)-1; j++ {
			current := path[j].(*ScannerNode).id
			next := path[j+1].(*ScannerNode).id
			funcs = append(funcs, edges[current][next])
		}

		transforms[i] = func(p Point3D) Point3D {
			for i := 0; i < len(funcs); i++ {
				p = funcs[i](p)
			}
			return p
		}
	}

	// Finally, use each transform to convert every beacon into the coordinate
	// space of scanner 0 so that we can determine how many beacons there actually
	// are.
	beacons := aoc.NewSet()
	for i, scanner := range scanners {
		for _, b := range scanner.Beacons {
			beacons.Add(transforms[i](b))
		}
	}
	fmt.Println(beacons.Size())
}

func NewMatrix(n int) [][]Transform {
	var m [][]Transform
	for i := 0; i < n; i++ {
		m = append(m, make([]Transform, n))
	}
	return m
}

func Overlaps(a, b *Scanner) (bool, Point3D, Point3D) {
	// Find a pair of beacons across the scanners that share distances to other beacons.
	// If there are 12 or more distances that are the same then the beacons are considered
	// to be the same.
	for i := 0; i < len(a.Beacons); i++ {
		for j := 0; j < len(b.Beacons); j++ {
			if a.Distances[i].Intersect(b.Distances[j]).Size() >= 12 {
				return true, a.Beacons[i], b.Beacons[j]
			}
		}
	}
	return false, Point3D{}, Point3D{}
}

// BuildTransform takes two sets of beacons that are known to overlap and determines
// transform functions that will convert coordinates from a to b.
func BuildTransform(a, b []Point3D, aBeacon, bBeacon Point3D) Transform {
	// We don't yet know which transform should be used, so try them all and see which
	// one works.
	for _, transform := range Transforms {
		dx := bBeacon.X - transform(aBeacon).X
		dy := bBeacon.Y - transform(aBeacon).Y
		dz := bBeacon.Z - transform(aBeacon).Z
		convert := func(p Point3D) Point3D {
			p = transform(p)
			p.X += dx
			p.Y += dy
			p.Z += dz
			return p
		}

		// Transform and translate a beacons
		aSet := aoc.NewSet()
		for _, p := range a {
			aSet.Add(convert(p))
		}

		// Now count how many b's match.
		var count int
		for _, bBeacon := range b {
			if aSet.Contains(bBeacon) {
				count++
			}
		}

		if count >= 12 {
			// We have a match, return a function that converts from a's to b's.
			return func(p Point3D) Point3D {
				p = transform(p)
				p.X += dx
				p.Y += dy
				p.Z += dz
				return p
			}
		}
	}

	return nil
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

type Scanner struct {
	Beacons   []Point3D
	Distances []aoc.Set
}

func (s *Scanner) AddBeacon(p Point3D) {
	index := len(s.Beacons)

	s.Beacons = append(s.Beacons, p)
	s.Distances = append(s.Distances, aoc.NewSet())

	for i, b := range s.Beacons {
		distance := p.Distance2(b)
		s.Distances[i].Add(distance)
		s.Distances[index].Add(distance)
	}
}

type ScannerNode struct {
	id    int
	edges [][]Transform
}

func (sn *ScannerNode) ID() string {
	return fmt.Sprintf("%d", sn.id)
}

func (sn *ScannerNode) Children() []aoc.Node {
	var children []aoc.Node
	for i := 0; i < len(sn.edges); i++ {
		if sn.edges[sn.id][i] != nil {
			children = append(children, &ScannerNode{id: i, edges: sn.edges})
		}
	}
	return children
}

func InputToScanners() []*Scanner {
	var scanners []*Scanner

	var scanner = new(Scanner)
	for _, line := range aoc.InputToLines(2021, 19) {
		var id string
		if _, err := fmt.Sscanf(line, "--- scanner %s ---", &id); err == nil {
			scanners = append(scanners, scanner)
			continue
		}

		if line == "" {
			scanner = new(Scanner)
			continue
		}

		var x, y, z int
		if _, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z); err == nil {
			scanner.AddBeacon(Point3D{X: x, Y: y, Z: z})
			continue
		}
	}

	// Now that the scanners have been created determine the distance from each beacon to
	// each other beacon.
	for _, s := range scanners {
		for _, origin := range s.Beacons {
			distances := aoc.NewSet()
			for _, destination := range s.Beacons {
				distances.Add(origin.Distance2(destination))
			}

			s.Distances = append(s.Distances, distances)
		}
	}

	return scanners
}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) Distance2(other Point3D) int {
	dx := other.X - p.X
	dy := other.Y - p.Y
	dz := other.Z - p.Z
	return dx*dx + dy*dy + dz*dz
}
