package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	words, message := InputToMessage()

	var count int
	for _, word := range words {
		var start int
		for {
			index := strings.Index(message[start:], word)
			if index == -1 {
				break
			}

			start += index + 1
			count++
		}
	}

	fmt.Println(count)
}

func InputToMessage() ([]string, string) {
	lines := InputToLines()
	_, words, _ := strings.Cut(lines[0], ":")
	return strings.Split(words, ","), lines[len(lines)-1]
}
