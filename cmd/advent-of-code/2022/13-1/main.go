package main

import (
	"encoding/json"
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	numbers := InputToNumbers()

	var sum int
	for i := 0; i < len(numbers); i += 2 {
		cmp := Compare(numbers[i], numbers[i+1])
		if cmp < 0 {
			sum += 1 + (i+1)/2
		}
	}
	fmt.Println(sum)
}

func Compare(lhs, rhs any) int {
	switch {
	case IsInt(lhs) && IsInt(rhs):
		lhs, rhs := int(lhs.(float64)), int(rhs.(float64))
		return lhs - rhs

	case IsList(lhs) && IsList(rhs):
		lhs, rhs := lhs.([]any), rhs.([]any)
		for i := 0; i < lib.Min(len(lhs), len(rhs)); i++ {
			cmp := Compare(lhs[i], rhs[i])
			if cmp != 0 {
				return cmp
			}
		}
		return len(lhs) - len(rhs)

	case IsInt(lhs):
		return Compare([]any{lhs}, rhs)

	case IsInt(rhs):
		return Compare(lhs, []any{rhs})
	}

	return 0 // This should never happen
}

func IsInt(n any) bool {
	_, ok := n.(float64)
	return ok
}
func IsList(n any) bool {
	_, ok := n.([]any)
	return ok
}

func InputToNumbers() []any {
	var numbers []any
	for _, line := range lib.InputToLines() {
		if len(line) == 0 {
			continue
		}

		var number any
		_ = json.Unmarshal([]byte(line), &number)
		numbers = append(numbers, number)
	}

	return numbers
}
