package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	var count int
	for _, report := range InputToReports() {
		if IsSafe(report) {
			count++
		}
	}
	fmt.Println(count)
}

func IsSafe(report []int) bool {
	for remove := -1; remove < len(report); remove++ {
		var fwd []int
		if remove == -1 {
			fwd = report
		} else {
			fwd = Without(report, remove)
		}
		bwd := Reversed(fwd)

		isSafeFwd, isSafeBwd := true, true
		for n := 1; (isSafeFwd || isSafeBwd) && n < len(fwd); n++ {
			if delta := fwd[n] - fwd[n-1]; delta < 1 || delta > 3 {
				isSafeFwd = false
			}
			if delta := bwd[n] - bwd[n-1]; delta < 1 || delta > 3 {
				isSafeBwd = false
			}
		}

		if isSafeFwd || isSafeBwd {
			return true
		}
	}

	return false
}

func Without[T any](s []T, index int) []T {
	c := make([]T, len(s))
	copy(c, s)
	return append(c[:index], c[index+1:]...)
}

func InputToReports() [][]int {
	return InputLinesTo(func(s string) []int {
		var ns []int
		for _, f := range strings.Fields(s) {
			ns = append(ns, ParseInt(f))
		}
		return ns
	})
}
