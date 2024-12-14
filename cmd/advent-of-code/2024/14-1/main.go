package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	W, H := 101, 103

	robots := InputToRobots()
	for n := 0; n < 100; n++ {
		for i := range robots {
			robots[i].Px = (robots[i].Px + robots[i].Vx + W) % W
			robots[i].Py = (robots[i].Py + robots[i].Vy + H) % H
		}
	}

	var q1, q2, q3, q4 int
	for _, r := range robots {
		switch {
		case r.Px < W/2 && r.Py < H/2:
			q1++
		case r.Px > W/2 && r.Py < H/2:
			q2++
		case r.Px < W/2 && r.Py > H/2:
			q3++
		case r.Px > W/2 && r.Py > H/2:
			q4++
		}
	}

	fmt.Println(q1 * q2 * q3 * q4)
}

type Robot struct {
	Px, Py int
	Vx, Vy int
}

func InputToRobots() []Robot {
	return InputLinesTo(func(s string) Robot {
		s = strings.ReplaceAll(s, ",", " ")
		s = strings.ReplaceAll(s, "=", " ")
		fields := strings.Fields(s)

		return Robot{
			Px: ParseInt(fields[1]),
			Py: ParseInt(fields[2]),
			Vx: ParseInt(fields[4]),
			Vy: ParseInt(fields[5]),
		}
	})
}
