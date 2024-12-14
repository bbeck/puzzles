package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	W, H := 101, 103

	robots := InputToRobots()
	for n := 1; ; n++ {
		for i := range robots {
			robots[i].Px = (robots[i].Px + robots[i].Vx + W) % W
			robots[i].Py = (robots[i].Py + robots[i].Vy + H) % H
		}

		if IsChristmasTree(robots, W, H) {
			fmt.Println(n)
			break
		}
	}
}

func IsChristmasTree(robots []Robot, W, H int) bool {
	// The tree contains a box of thickness 1 around it which should be easy to
	// detect as it'll leave a pair of rows and columns with a large number of
	// robots in them.
	rows := make(map[int]int)
	cols := make(map[int]int)
	for _, r := range robots {
		rows[r.Py]++
		cols[r.Px]++
	}

	rs := slices.Sorted(maps.Values(rows))
	r1, r2 := rs[len(rs)-1], rs[len(rs)-2]

	cs := slices.Sorted(maps.Values(cols))
	c1, c2 := cs[len(cs)-1], cs[len(cs)-2]

	return r1 > W/4 && r2 > W/4 && c1 > H/4 && c2 > H/4
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
