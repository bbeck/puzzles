package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib/in"
)

const N int = 100

var Arrows = []string{"R", "G", "B"}

func main() {
	var balloons = strings.Split(strings.Repeat(in.String(), N), "")

	var count = 0
	for len(balloons) > 0 {
		if Arrows[count%3] == balloons[0] && len(balloons)%2 == 0 {
			// We're going to pop 2 balloons
			other := len(balloons) / 2
			balloons = append(balloons[1:other], balloons[other+1:]...)
		} else {
			// We're going to pop 1 balloon
			balloons = balloons[1:]
		}
		count++

		if len(balloons) == 0 {
			break
		}
	}
	fmt.Println(count)
}
