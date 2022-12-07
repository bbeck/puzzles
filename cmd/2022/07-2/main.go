package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	dir := InputToFilesystem()

	sizes := make(map[*Directory]int)
	dir.DFS(func(d *Directory) {
		size := aoc.Sum(aoc.GetMapValues(d.Files)...)
		for _, child := range d.Dirs {
			size += sizes[child]
		}
		sizes[d] = size
	})

	free := math.MaxInt
	for _, size := range sizes {
		if 40000000-sizes[dir]+size >= 0 {
			free = aoc.Min(free, size)
		}
	}
	fmt.Println(free)
}

type Directory struct {
	Name  string
	Dirs  map[string]*Directory
	Files map[string]int
}

func NewDirectory(name string) *Directory {
	return &Directory{
		Name:  name,
		Dirs:  make(map[string]*Directory),
		Files: make(map[string]int),
	}
}

func (d *Directory) DFS(fn func(*Directory)) {
	for _, child := range d.Dirs {
		child.DFS(fn)
	}
	fn(d)
}

func InputToFilesystem() *Directory {
	var lines aoc.Deque[string]
	for _, line := range aoc.InputToLines(2022, 7) {
		lines.PushBack(line)
	}

	root := NewDirectory("/")

	var path []string
	for !lines.Empty() {
		line := lines.PopFront()

		switch {
		case line == "$ cd /":
			path = nil

		case line == "$ cd ..":
			path = path[:len(path)-1]

		case strings.HasPrefix(line, "$ cd"):
			path = append(path, line[5:])

		case line == "$ ls":
			dir := GetDirectory(root, path)
			for !lines.Empty() {
				if line := lines.PeekFront(); line[0] == '$' {
					break
				}

				line := lines.PopFront()
				size, name, _ := strings.Cut(line, " ")
				if size != "dir" {
					dir.Files[name] = aoc.ParseInt(size)
				}
			}
		}
	}

	return root
}

func GetDirectory(root *Directory, path []string) *Directory {
	for _, part := range path {
		if _, ok := root.Dirs[part]; !ok {
			root.Dirs[part] = NewDirectory(part)
		}
		root = root.Dirs[part]
	}
	return root
}
