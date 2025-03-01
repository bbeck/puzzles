package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var seen Set[Point2D] = SetFrom(Origin2D)

	var santa, robot Point2D
	for n, dir := range InputToHeadings() {
		if n%2 == 0 {
			santa = santa.Move(dir)
		} else {
			robot = robot.Move(dir)
		}
		seen.Add(santa, robot)
	}

	fmt.Println(len(seen))
}

func InputToHeadings() []Heading {
	var headings []Heading
	for in.HasNext() {
		switch in.Byte() {
		case '^':
			headings = append(headings, Up)
		case '<':
			headings = append(headings, Left)
		case '>':
			headings = append(headings, Right)
		case 'v':
			headings = append(headings, Down)
		}
	}

	return headings
}
