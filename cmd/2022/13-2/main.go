package main

import (
	"encoding/json"
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"reflect"
	"sort"
)

func main() {
	two, six := []any{[]any{2.}}, []any{[]any{6.}}
	numbers := InputToNumbers()
	numbers = append(numbers, two)
	numbers = append(numbers, six)

	sort.Slice(numbers, func(i, j int) bool {
		return Compare(numbers[i], numbers[j]) < 0
	})

	var indices []int
	for i, number := range numbers {
		if reflect.DeepEqual(number, two) || reflect.DeepEqual(number, six) {
			indices = append(indices, i+1)
		}
	}
	fmt.Println(puz.Product(indices...))
}

func Compare(lhs, rhs any) int {
	switch {
	case IsInt(lhs) && IsInt(rhs):
		lhs, rhs := int(lhs.(float64)), int(rhs.(float64))
		return lhs - rhs

	case IsList(lhs) && IsList(rhs):
		lhs, rhs := lhs.([]any), rhs.([]any)
		for i := 0; i < puz.Min(len(lhs), len(rhs)); i++ {
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
	for _, line := range puz.InputToLines(2022, 13) {
		if len(line) == 0 {
			continue
		}

		var number any
		_ = json.Unmarshal([]byte(line), &number)
		numbers = append(numbers, number)
	}

	return numbers
}
