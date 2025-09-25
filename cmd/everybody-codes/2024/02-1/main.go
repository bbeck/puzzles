package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib/in"
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
	_, words := in.Cut(":")
	in.Line()
	return strings.Split(words, ","), in.Line()
}
