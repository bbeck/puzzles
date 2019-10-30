package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	area := InputToArea(2018, 18)

	// With enough iterations we eventually end up in a repeating cycle of areas.
	// Start iterating through the areas keeping track of the progression until
	// we end up repeating ourselves.
	next := make(map[string]*Area)
	for tm := 0; tm < 1_000_000_000; tm++ {
		key := area.Key()
		if a, ok := next[key]; ok {
			area = a
			continue
		}

		area = area.Next()
		next[key] = area
	}

	var trees, lumberyards int
	for _, s := range area.acres {
		switch s {
		case TREES:
			trees++
		case LUMBERYARD:
			lumberyards++
		}
	}

	fmt.Printf("total resource value: %d\n", trees*lumberyards)
}

const (
	OPEN       = "."
	TREES      = "|"
	LUMBERYARD = "#"
)

type Area struct {
	acres  map[aoc.Point2D]string
	width  int
	height int
	key    string
}

func InputToArea(year, day int) *Area {
	acres := make(map[aoc.Point2D]string)
	width := 0
	height := 0

	for y, line := range aoc.InputToLines(year, day) {
		height = y + 1
		for x, c := range line {
			width = x + 1
			p := aoc.Point2D{X: x, Y: y}
			acres[p] = string(c)
		}
	}

	return &Area{
		acres:  acres,
		width:  width,
		height: height,
	}
}

func (a *Area) Key() string {
	if a.key == "" {
		h := md5.New()
		_, _ = io.WriteString(h, a.String())
		a.key = hex.EncodeToString(h.Sum(nil))
	}

	return a.key
}

func (a *Area) String() string {
	var builder strings.Builder
	for y := 0; y < a.height; y++ {
		for x := 0; x < a.width; x++ {
			builder.WriteString(a.acres[aoc.Point2D{X: x, Y: y}])
		}

		builder.WriteString("\n")
	}
	return builder.String()
}

func (a *Area) Next() *Area {
	// return the count of open, trees, and lumber yards surrounding the given
	// point.
	neighbors := func(p aoc.Point2D) (int, int, int) {
		ps := []aoc.Point2D{
			p.Up().Left(),
			p.Up(),
			p.Up().Right(),
			p.Left(),
			p.Right(),
			p.Down().Left(),
			p.Down(),
			p.Down().Right(),
		}

		var open, trees, lumberyards int
		for _, p := range ps {
			switch a.acres[p] {
			case OPEN:
				open++
			case TREES:
				trees++
			case LUMBERYARD:
				lumberyards++
			}
		}

		return open, trees, lumberyards
	}

	acres := make(map[aoc.Point2D]string)
	for y := 0; y < a.height; y++ {
		for x := 0; x < a.width; x++ {
			p := aoc.Point2D{X: x, Y: y}
			_, trees, lumberyards := neighbors(p)

			switch a.acres[p] {
			case OPEN:
				if trees >= 3 {
					acres[p] = TREES
				} else {
					acres[p] = OPEN
				}

			case TREES:
				if lumberyards >= 3 {
					acres[p] = LUMBERYARD
				} else {
					acres[p] = TREES
				}

			case LUMBERYARD:
				if trees >= 1 && lumberyards >= 1 {
					acres[p] = LUMBERYARD
				} else {
					acres[p] = OPEN
				}
			}
		}
	}

	return &Area{
		acres:  acres,
		width:  a.width,
		height: a.height,
	}
}
