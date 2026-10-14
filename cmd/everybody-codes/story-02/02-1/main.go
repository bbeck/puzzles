package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib/in"
)

var Arrows = []string{"R", "G", "B"}

func main() {
	var balloons = strings.Split(in.String(), "")

	var count = 0
	for len(balloons) > 0 {
		for len(balloons) > 0 {
			var first string
			first, balloons = balloons[0], balloons[1:]
			if first != Arrows[count%3] {
				break
			}
		}

		count++
	}
	fmt.Println(count)
}
