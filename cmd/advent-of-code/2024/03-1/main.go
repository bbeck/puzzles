package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`mul[(]([0-9]+),([0-9]+)[)]`)

func main() {
	var sum int
	for _, m := range regex.FindAllString(InputToString(), -1) {
		m = strings.ReplaceAll(m, "(", " ")
		m = strings.ReplaceAll(m, ")", " ")
		m = strings.ReplaceAll(m, ",", " ")

		fields := strings.Fields(m)
		sum += ParseInt(fields[1]) * ParseInt(fields[2])
	}
	fmt.Println(sum)
}
