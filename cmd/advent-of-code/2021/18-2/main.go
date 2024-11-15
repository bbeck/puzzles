package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	ns := InputToNumbers()

	var best int
	for i := 0; i < len(ns); i++ {
		for j := i + 1; j < len(ns); j++ {
			best = puz.Max(best, Magnitude(Add(ns[i], ns[j])), Magnitude(Add(ns[j], ns[i])))
		}
	}
	fmt.Println(best)
}

func Magnitude(n *Number) int {
	if n == nil {
		return 0
	}
	if n.Left == nil && n.Right == nil {
		return n.Value
	}
	return 3*Magnitude(n.Left) + 2*Magnitude(n.Right)
}

func Add(a, b *Number) *Number {
	return Reduce(&Number{
		Left:  a.Copy(),
		Right: b.Copy(),
	})
}

func Reduce(n *Number) *Number {
	for {
		var changed bool
		if changed = Explode(n); changed {
			continue
		}

		if changed = Split(n); changed {
			continue
		}

		break
	}

	return n
}

func Explode(root *Number) bool {
	var ordering []*Number

	var traverse func(*Number)
	traverse = func(n *Number) {
		if n.Left == nil && n.Right == nil {
			ordering = append(ordering, n)
		}
		if n.Left != nil {
			traverse(n.Left)
		}
		if n.Right != nil {
			traverse(n.Right)
		}
	}
	traverse(root)

	var neighbors func(*Number) (*Number, *Number)
	neighbors = func(n *Number) (*Number, *Number) {
		for i := 0; i < len(ordering); i++ {
			if ordering[i] == n {
				var left, right *Number
				if i > 0 {
					left = ordering[i-1]
				}
				if i < len(ordering)-1 {
					right = ordering[i+1]
				}
				return left, right
			}
		}
		return nil, nil
	}

	var explode func(n *Number, depth int) bool
	explode = func(n *Number, depth int) bool {
		if n.Left == nil && n.Right == nil {
			return false
		}

		if depth < 4 {
			if changed := explode(n.Left, depth+1); changed {
				return true
			}

			if changed := explode(n.Right, depth+1); changed {
				return true
			}

			return false
		}

		if left, _ := neighbors(n.Left); left != nil {
			left.Value += n.Left.Value
		}

		if _, right := neighbors(n.Right); right != nil {
			right.Value += n.Right.Value
		}

		n.Value = 0
		n.Left = nil
		n.Right = nil
		return true
	}

	return explode(root, 0)
}

func Split(n *Number) bool {
	if n.Value >= 10 {
		n.Left = &Number{Value: n.Value / 2}
		n.Right = &Number{Value: n.Value/2 + n.Value%2}
		n.Value = 0
		return true
	}

	return (n.Left != nil && Split(n.Left)) || (n.Right != nil && Split(n.Right))
}

type Number struct {
	// Leaf
	Value int

	// Pair
	Left  *Number
	Right *Number
}

func (n *Number) Copy() *Number {
	c := &Number{Value: n.Value}
	if n.Left != nil {
		c.Left = n.Left.Copy()
	}
	if n.Right != nil {
		c.Right = n.Right.Copy()
	}
	return c
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

		return input, &Number{Left: left, Right: right}
	}

	// Left
	value, input := puz.ParseInt(input[0:1]), input[1:]
	return input, &Number{Value: value}
}

func InputToNumbers() []*Number {
	return puz.InputLinesTo(func(line string) *Number {
		_, n := ParseNumber(line)
		return n
	})
}
