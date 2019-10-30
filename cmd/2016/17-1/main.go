package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	start := Room{
		location: aoc.Point2D{0, 0},
		passcode: aoc.InputToString(2016, 17),
	}

	goal := aoc.Point2D{3, 3}

	path, found := aoc.BreadthFirstSearch(start, func(node aoc.Node) bool {
		return node.(Room).location == goal
	})

	if !found {
		log.Fatal("no path found")
	}

	passcode := path[len(path)-1].(Room).passcode
	fmt.Printf("passcode: %s\n", passcode[len(start.passcode):])
}

type Room struct {
	location aoc.Point2D
	passcode string
}

func (r Room) ID() string {
	return r.passcode
}

func (r Room) Children() []aoc.Node {
	h := md5.New()
	_, _ = io.WriteString(h, r.passcode)
	sum := hex.EncodeToString(h.Sum(nil))

	open := func(c byte) bool {
		return c == 'b' || c == 'c' || c == 'd' || c == 'e' || c == 'f'
	}

	var children []aoc.Node

	if r.location.Y > 0 && open(sum[0]) {
		children = append(children, Room{r.location.Up(), r.passcode + "U"})
	}

	if r.location.Y < 3 && open(sum[1]) {
		children = append(children, Room{r.location.Down(), r.passcode + "D"})
	}

	if r.location.X > 0 && open(sum[2]) {
		children = append(children, Room{r.location.Left(), r.passcode + "L"})
	}

	if r.location.X < 3 && open(sum[3]) {
		children = append(children, Room{r.location.Right(), r.passcode + "R"})
	}

	return children
}
