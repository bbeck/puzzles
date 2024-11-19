package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"path/filepath"
	"strings"
)

func main() {
	sizes := make(map[string]int)

	var path []string
	for _, line := range lib.InputToLines() {
		words := strings.Fields(line)

		switch {
		case words[0] == "$" && words[1] == "cd" && words[2] == "/":
			path = nil

		case words[0] == "$" && words[1] == "cd" && words[2] == "..":
			path = path[:len(path)-1]

		case words[0] == "$" && words[1] == "cd":
			path = append(path, words[2])

		case words[0] != "$" && words[0] != "dir":
			size := lib.ParseInt(words[0])

			// Add this size to the current directory and all parents
			sizes[filepath.Join(path...)] += size
			for end := 0; end < len(path); end++ {
				sizes[filepath.Join(path[:end]...)] += size
			}
		}
	}

	var sum int
	for _, size := range sizes {
		if size < 100000 {
			sum += size
		}
	}
	fmt.Println(sum)
}
