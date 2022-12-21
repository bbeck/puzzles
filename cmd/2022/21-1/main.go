package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	values := InputToValues()
	fmt.Println(Eval("root", values))
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
	for _, line := range aoc.InputToLines(2022, 21) {
		line = strings.ReplaceAll(line, ":", "")
		fields := strings.Fields(line)

		var value Value
		if len(fields) == 2 {
			value = Value{IsConstant: true, Value: aoc.ParseInt(fields[1])}
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
