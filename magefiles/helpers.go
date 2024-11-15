package main

import (
	"fmt"
	"github.com/bitfield/script"
	"os"
	"strconv"
	"strings"
)

func LookupInt(key string) (int, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return -1, fmt.Errorf("missing %s environment variable", key)
	}

	return ParseInt(value)
}

func ParseInt(s string) (int, error) {
	n, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return -1, err
	}
	return int(n), nil
}

func Files(paths ...string) string {
	var files []string
	for _, path := range paths {
		fs, err := script.FindFiles(path).Slice()
		if err != nil {
			continue
		}

		files = append(files, fs...)
	}

	return strings.Join(files, "\n")
}
