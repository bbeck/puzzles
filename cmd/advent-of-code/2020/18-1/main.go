package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"strings"
)

func main() {
	var sum int
	for _, line := range lib.InputToLines() {
		sum += Evaluate(line)
	}
	fmt.Println(sum)
}

func Evaluate(s string) int {
	var helper func(Node) int
	helper = func(node Node) int {
		switch node.Kind {
		case "number":
			return lib.ParseInt(node.Value)
		case "+":
			return helper(*node.Left) + helper(*node.Right)
		case "*":
			return helper(*node.Left) * helper(*node.Right)
		default:
			return 0
		}
	}

	node := Parse(Tokenize(s), 0)
	return helper(node)
}

func Parse(tokens *Tokens, lbp int) Node {
	// This implements a simple Pratt parser for operator precedence.
	left := nud(tokens)
	for tokens.Len() > 0 && lbp < bp(tokens.Peek()) {
		left = led(tokens, left)
	}
	return left
}

func nud(tokens *Tokens) Node {
	switch token := tokens.Pop(); token {
	case "(":
		node := Parse(tokens, 0)
		tokens.Pop() // right parenthesis
		return node

	default: // number
		return Node{Kind: "number", Value: token}
	}
}

func led(tokens *Tokens, lhs Node) Node {
	op := tokens.Pop()
	rhs := Parse(tokens, bp(op))
	return Node{Kind: op, Left: &lhs, Right: &rhs}
}

func bp(op string) int {
	if op == "+" || op == "*" {
		return 1
	}
	return 0
}

type Node struct {
	Kind        string
	Value       string
	Left, Right *Node
}

func Tokenize(s string) *Tokens {
	s = strings.ReplaceAll(s, "(", " ( ")
	s = strings.ReplaceAll(s, ")", " ) ")
	return &Tokens{Fields: strings.Fields(s)}
}

type Tokens struct {
	Fields []string
	Index  int
}

func (t *Tokens) Peek() string {
	return t.Fields[t.Index]
}

func (t *Tokens) Pop() string {
	t.Index++
	return t.Fields[t.Index-1]
}

func (t *Tokens) Len() int {
	return len(t.Fields) - t.Index
}
