package aoc

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
	bs, err := ioutil.ReadFile(InputFilename(year, day))
	if err != nil {
		log.Fatalf("unable to read input.txt: %+v", err)
	}

	return bs
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

// InputToInts reads the input file into a single integer.
func InputToInt(year, day int) int {
	for _, line := range InputToLines(year, day) {
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("unable to parse integer: %s", line)
		}

		return i
	}

	log.Fatal("unable to find integer in input file")
	return -1
}

// InputToInts reads the input file into a slice of integers.
func InputToInts(year, day int) []int {
	ints := make([]int, 0)
	for _, line := range InputToLines(year, day) {
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("unable to parse integer: %s", line)
		}

		ints = append(ints, i)
	}

	return ints
}
