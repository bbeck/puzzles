package main

import (
	"encoding/json"
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	lines := aoc.InputToLines(2022, 13)

	var sum int
	index := 1
	for i := 0; i < len(lines); i += 3 {
		lhs, rhs := ParseNumber(lines[i]), ParseNumber(lines[i+1])
		if a, _ := Compare(lhs, rhs, 0); a {
			sum += index
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
		fmt.Println()
		index++
		// break
	}
	fmt.Println(sum)
}

func Indent(n int) string {
	var sb strings.Builder
	for n > 0 {
		sb.WriteString("  ")
		n--
	}
	return sb.String()
}

type Answer bool
type KeepGoing bool

func Compare(lhs, rhs any, indent int) (Answer, KeepGoing) {
	fmt.Printf(Indent(indent)+"Compare %v vs %v\n", lhs, rhs)
	_, isLeftInt := lhs.(float64)
	_, isRightInt := rhs.(float64)
	if isLeftInt && isRightInt {
		if lhs.(float64) == rhs.(float64) {
			return Answer(false), KeepGoing(true)
		}
		return Answer(lhs.(float64) < rhs.(float64)), KeepGoing(false)
	}

	_, isLeftList := lhs.([]any)
	_, isRightList := rhs.([]any)
	if isLeftList && isRightList {
		left := lhs.([]any)
		right := rhs.([]any)
		for i := 0; i < aoc.Min(len(left), len(right)); i++ {
			a, kg := Compare(left[i], right[i], indent+1)
			if !kg {
				return a, kg
			}
		}
		if len(left) < len(right) {
			return Answer(true), KeepGoing(false)
		} else if len(left) == len(right) {
			return Answer(true), KeepGoing(true)
		}
		return Answer(false), KeepGoing(false)
	}

	if isLeftInt {
		return Compare([]any{lhs}, rhs, indent+1)
	}
	return Compare(lhs, []any{rhs}, indent+1)
}

func ParseNumber(input string) any {
	var f any
	json.Unmarshal([]byte(input), &f)
	return f
}
