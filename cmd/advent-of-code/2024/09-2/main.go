package main

import (
	"fmt"
	"slices"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	files := InputToFiles()

	// Locate a file by its id.
	findById := func(id int) int {
		for i, f := range files {
			if f.ID == id {
				return i
			}
		}
		return -1
	}

	// Locate a file before the given offset that has a gap of at least the given
	// size.
	findByGap := func(offset, gap int) (int, bool) {
		for i, f := range files {
			if f.Gap >= gap && f.Offset+f.Len < offset {
				return i, true
			}
		}
		return -1, false
	}

	// Resort the files after their offsets have been changed.
	fixOrder := func() {
		slices.SortFunc(files, func(a, b File) int {
			return a.Offset - b.Offset
		})
	}

	for id := len(files) - 1; id > 0; id-- {
		move := findById(id)

		after, ok := findByGap(files[move].Offset, files[move].Len)
		if !ok {
			continue
		}

		files[move-1].Gap += files[move].Len + files[move].Gap
		files[move].Offset = files[after].Offset + files[after].Len
		files[move].Gap = files[after].Gap - files[move].Len
		files[after].Gap = 0
		fixOrder()
	}

	var sum int
	var offset int
	for _, f := range files {
		for n := 0; n < f.Len; n++ {
			sum += offset * f.ID
			offset++
		}
		offset += f.Gap
	}
	fmt.Println(sum)
}

type File struct {
	ID     int
	Len    int
	Offset int
	Gap    int
}

func InputToFiles() []File {
	var files []File
	var offset int
	for i, bs := range Chunks(InputToBytes(), 2) {
		var l, g int
		l = ParseInt(string(bs[0]))
		if len(bs) == 2 {
			g = ParseInt(string(bs[1]))
		} else {
			g = 0
		}
		files = append(files, File{
			ID:     i,
			Len:    l,
			Offset: offset,
			Gap:    g,
		})
		offset += l + g
	}

	return files
}
