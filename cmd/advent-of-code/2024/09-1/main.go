package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
)

func main() {
	files := InputToFiles()

	take := func(upto int) (int, int) {
		for i := len(files) - 1; i >= 0; i-- {
			if files[i].Len > 0 {
				n := Min(files[i].Len, upto)
				files[i].Len -= n
				return files[i].ID, n
			}
		}
		return 0, 0
	}

	var offset, sum int
	for i := 0; i < len(files); i++ {
		// Steal from the end until we reach the offset of the current file
		for files[i].Len > 0 && offset < files[i].Offset {
			value, n := take(files[i].Offset - offset)
			sum += value * (n*offset + n*(n-1)/2)
			offset += n
		}

		// Offset is now up to the current file, include it in its entirety
		value, n := files[i].ID, files[i].Len
		sum += value * (n*offset + n*(n-1)/2)
		offset += n
		files[i].Len = 0
	}
	fmt.Println(sum)
}

type File struct {
	ID     int
	Len    int
	Offset int
}

func InputToFiles() []File {
	var files []File
	var offset int
	for i, ch := range InputToString() {
		if i%2 == 0 {
			files = append(files, File{
				ID:     i / 2,
				Len:    ParseInt(string(ch)),
				Offset: offset,
			})
		}
		offset += ParseInt(string(ch))
	}

	return files
}
