package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ns := InputToNumbers()

	sum := ns[0]
	for i := 1; i < len(ns); i++ {
		sum = Add(sum, ns[i])
	}
	fmt.Println(Magnitude(sum))
}

func Add(a, b *Number) *Number {
	// To add two snailfish numbers, form a pair from the left and right parameters
	// of the addition operator.
	sum := &Number{
		Kind:  "pair",
		Left:  a,
		Right: b,
	}

	// After adding the numbers must be reduced.  To reduce a snailfish number, you
	// must repeatedly do the first action in this list that applies to the
	// snailfish number:
	//
	// - If any pair is nested inside four pairs, the leftmost such pair explodes.
	// - If any regular number is 10 or greater, the leftmost such regular number splits.
	for {
		if changed, s := Explode(sum); changed {
			sum = s
			continue
		}
		if changed, s := Split(sum); changed {
			sum = s
			continue
		}

		break
	}

	return sum
}

func Magnitude(x *Number) int {
	if x.Kind == "regular" {
		return x.Value
	}
	return 3*Magnitude(x.Left) + 2*Magnitude(x.Right)
}

func Explode(x *Number) (bool, *Number) {
	// Helper that finds the leftmost regular number within x and adds the regular number n to it.
	var addLeft func(x, n *Number) *Number
	addLeft = func(x, n *Number) *Number {
		if n == nil {
			return x
		}

		if x.Kind == "regular" {
			return &Number{Kind: "regular", Value: x.Value + n.Value}
		}

		return &Number{Kind: "pair", Left: addLeft(x.Left, n), Right: x.Right}
	}

	// Helper that finds the rightmost regular number within x and adds the regular number n to it.
	var addRight func(x, n *Number) *Number
	addRight = func(x, n *Number) *Number {
		if n == nil {
			return x
		}

		if x.Kind == "regular" {
			return &Number{Kind: "regular", Value: x.Value + n.Value}
		}

		return &Number{Kind: "pair", Left: x.Left, Right: addRight(x.Right, n)}
	}

	var explode func(x *Number, n int) (bool, *Number, *Number, *Number)
	explode = func(x *Number, n int) (bool, *Number, *Number, *Number) {
		// If we reach a single value before getting to a depth of 4, then there's
		// nothing to explode.  Return no change.
		if x.Kind == "regular" {
			return false, nil, x, nil
		}

		// We've reached the desired depth, explode this number.  We'll replace
		// the current pair with a 0 and return the regular number at the left and
		// right to be dealt with at the previous layer of the recursion.
		if n == 0 {
			return true, x.Left, &Number{Kind: "regular", Value: 0}, x.Right
		}

		// Try exploding the leftmost pair first.
		if changed, left, nx, right := explode(x.Left, n-1); changed {
			return true, left, &Number{Kind: "pair", Left: nx, Right: addLeft(x.Right, right)}, nil
		}

		// Try exploding the rightmost pair.
		if changed, left, nx, right := explode(x.Right, n-1); changed {
			return true, nil, &Number{Kind: "pair", Left: addRight(x.Left, left), Right: nx}, right
		}

		// We weren't able to explode, return no change.
		return false, nil, x, nil
	}

	// Any values that need to be added to the leftmost or rightmost regular node that
	// propagate here can be ignored.  If they made it this far it means there were no
	// regular nodes to add them to.
	changed, _, x, _ := explode(x, 4)
	return changed, x
}

func Split(x *Number) (bool, *Number) {
	// Only regular numbers with a value of 10 or greater split.  To split a regular number,
	// replace it with a pair; the left element of the pair should be the regular number
	// divided by two and rounded down, while the right element of the pair should be the
	// regular number divided by two and rounded up.
	if x.Kind == "regular" {
		if x.Value < 10 {
			return false, x
		}

		left := &Number{Kind: "regular", Value: x.Value / 2}
		right := &Number{Kind: "regular", Value: (x.Value + 1) / 2}
		return true, &Number{Kind: "pair", Left: left, Right: right}
	}

	// This is a pair, try to recursively split their sub-numbers.  We'll only split one
	// at a time because addition says that once you perform a split operation you have
	// to then check for explodes that can be done.
	if changed, left := Split(x.Left); changed {
		return true, &Number{Kind: "pair", Left: left, Right: x.Right}
	}

	if changed, right := Split(x.Right); changed {
		return true, &Number{Kind: "pair", Left: x.Left, Right: right}
	}

	// We were unable to split this number or any of its sub-numbers.
	return false, x
}

type Number struct {
	Kind string

	// Leaf
	Value int

	// Pair
	Left  *Number
	Right *Number
}

func (n *Number) String() string {
	if n.Kind == "regular" {
		return fmt.Sprintf("%d", n.Value)
	}
	return fmt.Sprintf("[%s,%s]", n.Left, n.Right)
}

func ParseNumber(input string) (string, *Number) {
	if len(input) == 0 {
		return input, nil
	}

	// Pair
	if input[0] == '[' {
		input = input[1:] // skip [
		input, left := ParseNumber(input)
		input = input[1:] // skip ,
		input, right := ParseNumber(input)
		input = input[1:] // skip ]

		return input, &Number{Kind: "pair", Left: left, Right: right}
	}

	// Regular
	value, input := aoc.ParseInt(input[0:1]), input[1:]
	return input, &Number{Kind: "regular", Value: value}
}

func InputToNumbers() []*Number {
	var ns []*Number
	for _, line := range aoc.InputToLines(2021, 18) {
		_, n := ParseNumber(line)
		ns = append(ns, n)
	}
	return ns
}
