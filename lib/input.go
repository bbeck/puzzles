package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var CallerDirectoryRegexp = regexp.MustCompile(
	`.*/([^/]+)/([0-9]{4})/([0-9]{2})[-]([0-9])/main.go$`,
)

// InputFilename determines the input file for a specific day's part.
func InputFilename() string {
	// Determine the site, year, day and part from the caller.  In all cases I've
	// tested this info has been present in the caller entry, it may be at a
	// different level depending on how many method calls happened first.
	var m []string
	for n := 0; n < 10; n++ {
		_, caller, _, _ := runtime.Caller(n)

		m = CallerDirectoryRegexp.FindStringSubmatch(caller)
		if m != nil {
			break
		}
	}

	if m == nil {
		panic("unable to find site, year, day and part from caller")
	}

	site, year, day, part := m[1], m[2], m[3], m[4]

	// Determine the path to the input.txt relative to the working directory.
	wd, _ := os.Getwd()

	switch site {
	case "advent-of-code":
		// For Advent of Code there's a single input shared by all parts.  The
		// input always resides in the part 1 directory.
		if strings.HasSuffix(wd, fmt.Sprintf("/cmd/%s/%s/%s-%s", site, year, day, part)) {
			// We're in the problem's directory.
			return fmt.Sprintf("../%s-1/input.txt", day)
		}

		// Assume we're in the root of the project.
		return fmt.Sprintf("cmd/%s/%s/%s-1/input.txt", site, year, day)

	case "everybody-codes":
		// For Everybody Codes there's a separate input for each part.
		if strings.HasSuffix(wd, fmt.Sprintf("/cmd/%s/%s/%s-%s", site, year, day, part)) {
			return "input.txt"
		}

		// Assume we're in the root of the project.
		return fmt.Sprintf("cmd/%s/%s/%s-%s/input.txt", site, year, day, part)
	}

	log.Fatalf("unable to locate input.txt")
	return ""
}

// InputToBytes reads the entire input file into a slice of bytes.
func InputToBytes() []byte {
	bs, err := os.ReadFile(InputFilename())
	if err != nil {
		log.Fatalf("unable to read input.txt: %+v", err)
	}

	return bytes.TrimSpace(bs)
}

// InputToString reads the entire input file into a string.
func InputToString() string {
	return string(InputToBytes())
}

// InputToLines reads the input file into a slice of strings with each string
// representing a line of the file.  The newline character is not included.
func InputToLines() []string {
	file, err := os.Open(InputFilename())
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
func InputLinesTo[T any](parse func(string) T) []T {
	var ts []T
	for _, line := range InputToLines() {
		ts = append(ts, parse(line))
	}

	return ts
}

// InputToInt reads the input file into a single integer.
func InputToInt() int {
	return InputToInts()[0]
}

// InputToInts reads the input file into a slice of integers.
func InputToInts() []int {
	return InputLinesTo(func(s string) int {
		n, _ := strconv.Atoi(s)
		return n
	})
}

// InputToGrid2D builds a Grid2D instance from the input using the provided
// function to determine the value for each cell in the grid.
func InputToGrid2D[T any](fn func(int, int, string) T) Grid2D[T] {
	lines := InputToLines()

	grid := NewGrid2D[T](len(lines[0]), len(lines))
	for y, line := range lines {
		for x, c := range line {
			grid.Set(x, y, fn(x, y, string(c)))
		}
	}

	return grid
}

// InputToStringGrid2D builds a Grid2D[string] instance from the input.
func InputToStringGrid2D() Grid2D[string] {
	return InputToGrid2D(func(_ int, _ int, s string) string {
		return s
	})
}

// InputToIntGrid2D builds a Grid2D[int] instance from the input.
func InputToIntGrid2D() Grid2D[int] {
	return InputToGrid2D(func(_ int, _ int, s string) int {
		return ParseInt(s)
	})
}
