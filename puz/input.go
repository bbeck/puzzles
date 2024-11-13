package puz

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

// InputFilename determines the input file for a specific day's part.
func InputFilename(year, day int) string {
	return fmt.Sprintf("cmd/%d/%02d-1/input.txt", year, day)
}

// InputToBytes reads the entire input file into a slice of bytes.
func InputToBytes(year, day int) []byte {
	bs, err := os.ReadFile(InputFilename(year, day))
	if err != nil {
		log.Fatalf("unable to read input.txt: %+v", err)
	}

	return bytes.TrimSpace(bs)
}

// InputToString reads the entire input file into a string.
func InputToString(year, day int) string {
	return string(InputToBytes(year, day))
}

// InputToLines reads the input file into a slice of strings with each string
// representing a line of the file.  The newline character is not included.
func InputToLines(year, day int) []string {
	file, err := os.Open(InputFilename(year, day))
	if err != nil {
		log.Fatalf("unable to open input.txt: %+v", err)
	}
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error while reading input.txt: %+v", err)
	}

	return lines
}

// InputLinesTo transforms each line of the input file into an instance
// returned by a transform function.  The instances are returned in a
// slice in the same order as they appear in the file.
func InputLinesTo[T any](year, day int, parse func(string) T) []T {
	var ts []T
	for _, line := range InputToLines(year, day) {
		ts = append(ts, parse(line))
	}

	return ts
}

// InputToInt reads the input file into a single integer.
func InputToInt(year, day int) int {
	return InputToInts(year, day)[0]
}

// InputToInts reads the input file into a slice of integers.
func InputToInts(year, day int) []int {
	return InputLinesTo(year, day, func(s string) int {
		n, _ := strconv.Atoi(s)
		return n
	})
}

// InputToGrid2D builds a Grid2D instance from the input using the provided
// function to determine the value for each cell in the grid.
func InputToGrid2D[T any](year, day int, fn func(int, int, string) T) Grid2D[T] {
	lines := InputToLines(year, day)

	grid := NewGrid2D[T](len(lines[0]), len(lines))
	for y, line := range lines {
		for x, c := range line {
			grid.Set(x, y, fn(x, y, string(c)))
		}
	}

	return grid
}

// InputToStringGrid2D builds a Grid2D[string] instance from the input.
func InputToStringGrid2D(year, day int) Grid2D[string] {
	return InputToGrid2D(year, day, func(_ int, _ int, s string) string {
		return s
	})
}

// InputToIntGrid2D builds a Grid2D[int] instance from the input.
func InputToIntGrid2D(year, day int) Grid2D[int] {
	return InputToGrid2D(year, day, func(_ int, _ int, s string) int {
		return ParseInt(s)
	})
}
