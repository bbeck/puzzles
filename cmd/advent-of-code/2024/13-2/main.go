package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	var sum int
	for _, m := range InputToMachines() {
		if a, b, ok := Count(m); ok {
			sum += 3*a + b
		}
	}
	fmt.Println(sum)
}

func Count(m Machine) (int, int, bool) {
	// Any solution must satisfy a system of equations:
	//   a*ax + b*bx = px
	//   a*ay + b*by = py
	//
	// Performing some algebra results in:
	//   a = (by*px - bx*py)/(ax*by - ay*bx)
	//   b = (px - a*ax)/bx
	//
	// Thus if there's a solution it will be a single one that satisfies the
	// above equations.
	a := (m.By*m.Px - m.Bx*m.Py) / (m.Ax*m.By - m.Ay*m.Bx)
	b := (m.Px - a*m.Ax) / m.Bx

	// Verify that the solution works.
	if a >= 0 && b >= 0 && a*m.Ax+b*m.Bx == m.Px && a*m.Ay+b*m.By == m.Py {
		return a, b, true
	}

	return -1, -1, false
}

type Machine struct {
	Ax, Ay int
	Bx, By int
	Px, Py int
}

func InputToMachines() []Machine {
	lines := InputToLines()

	parse := func(s string) (int, int) {
		s = strings.ReplaceAll(s, "+", " ")
		s = strings.ReplaceAll(s, "=", " ")
		s = strings.ReplaceAll(s, ",", "")
		fields := strings.Fields(s)
		N := len(fields)

		return ParseInt(fields[N-3]), ParseInt(fields[N-1])
	}

	var machines []Machine
	for n := 0; n < len(lines); n += 4 {
		var m Machine
		m.Ax, m.Ay = parse(lines[n+0])
		m.Bx, m.By = parse(lines[n+1])
		m.Px, m.Py = parse(lines[n+2])
		m.Px, m.Py = m.Px+10000000000000, m.Py+10000000000000

		machines = append(machines, m)
	}
	return machines
}
