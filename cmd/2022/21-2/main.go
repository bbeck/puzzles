package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	// Through experimentation a few things were observed.  I'm not sure if these
	// are generally true for all inputs, or just true for mine.  I suspect they
	// are true for all inputs except which side they happen on.
	//
	// 1. Only the LHS of the `root` expression depended on `humn`.
	// 2. The LHS of the `root` expression is inversely proportional to `humn`.
	//
	// Because of this we can treat the RHS of the `root` expression as a target
	// and use a binary search to find the correct value of `humn`.  This should
	// require a logarithmic number of evaluations so the overall speed of the
	// eval function won't matter.
	values := InputToValues()
	lhs, rhs := values["root"].Left, values["root"].Right

	target := Eval(rhs, values)
	humn := sort.Search(10*target, func(humn int) bool {
		values["humn"] = Value{IsConstant: true, Value: humn}
		output := Eval(lhs, values)
		return output <= target
	})

	fmt.Println(humn)
}

func Eval(s string, values map[string]Value) int {
	value := values[s]
	if value.IsConstant {
		return value.Value
	}

	switch value.Op {
	case "+":
		return Eval(value.Left, values) + Eval(value.Right, values)
	case "-":
		return Eval(value.Left, values) - Eval(value.Right, values)
	case "*":
		return Eval(value.Left, values) * Eval(value.Right, values)
	case "/":
		return Eval(value.Left, values) / Eval(value.Right, values)
	default:
		return 0
	}
}

type Value struct {
	IsConstant bool
	Value      int

	IsExpression    bool
	Left, Right, Op string
}

func InputToValues() map[string]Value {
	values := make(map[string]Value)
	for _, line := range puz.InputToLines(2022, 21) {
		line = strings.ReplaceAll(line, ":", "")
		fields := strings.Fields(line)

		var value Value
		if len(fields) == 2 {
			value = Value{IsConstant: true, Value: puz.ParseInt(fields[1])}
		} else {
			value = Value{
				IsExpression: true,
				Left:         fields[1],
				Op:           fields[2],
				Right:        fields[3],
			}
		}

		values[fields[0]] = value
	}

	return values
}
