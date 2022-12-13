package main

import (
	"encoding/json"
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"reflect"
	"sort"
	"strings"
)

func main() {
	lines := aoc.InputToLines(2022, 13)

	two := []any{[]any{2.}}
	six := []any{[]any{6.}}

	var packets []any
	packets = append(packets, two)
	packets = append(packets, six)
	for i := 0; i < len(lines); i += 3 {
		lhs, rhs := ParseNumber(lines[i]), ParseNumber(lines[i+1])
		packets = append(packets, lhs)
		packets = append(packets, rhs)
	}

	sort.SliceStable(packets, func(i, j int) bool {
		a, _ := Compare(packets[i], packets[j], 0)
		return bool(a)
	})

	var indices []int
	for i, packet := range packets {
		if reflect.DeepEqual(two, packet) || reflect.DeepEqual(six, packet) {
			indices = append(indices, i+1)
		}
	}
	fmt.Println(indices[0] * indices[1])
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
	// fmt.Printf(Indent(indent)+"Compare %v vs %v\n", lhs, rhs)
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
