package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var sum int
	for _, line := range aoc.InputToLines(2020, 18) {
		sum += Evaluate(line)
	}
	fmt.Println(sum)
}

var Precedence = map[string]int{
	LParen: 0,
	Plus:   1,
	Times:  1,
}

func Evaluate(input string) int {
	lexer := NewLexer(input)

	operands := aoc.NewStack()
	operators := aoc.NewStack()

	for token := lexer.Next(); token.Kind != EOF; token = lexer.Next() {
		switch token.Kind {
		case Number:
			operands.Push(aoc.ParseInt(token.Literal))

		case Plus:
			for !operators.Empty() && Precedence[operators.Peek().(string)] >= Precedence[token.Kind] {
				Consume(operands, operators)
			}
			operators.Push(token.Kind)

		case Times:
			for !operators.Empty() && Precedence[operators.Peek().(string)] >= Precedence[token.Kind] {
				Consume(operands, operators)
			}
			operators.Push(token.Kind)

		case LParen:
			operators.Push(token.Kind)

		case RParen:
			for !operators.Empty() && operators.Peek().(string) != "(" {
				Consume(operands, operators)
			}
			operators.Pop()
		}
	}

	for !operators.Empty() {
		Consume(operands, operators)
	}

	return operands.Pop().(int)
}

func Consume(operands, operators *aoc.Stack) {
	rhs := operands.Pop().(int)
	lhs := operands.Pop().(int)

	switch operators.Pop() {
	case Plus:
		operands.Push(lhs + rhs)
	case Times:
		operands.Push(lhs * rhs)
	}
}

const (
	EOF    string = "EOF"
	Number        = "Number"
	Plus          = "+"
	Times         = "*"
	LParen        = "("
	RParen        = ")"
)

type Token struct {
	Kind    string
	Literal string
}

// Lexer is a simple tokenizer that offers the ability to peek one token into
// the future.
type Lexer struct {
	scanner *bufio.Scanner
	next    *Token
}

func NewLexer(input string) *Lexer {
	// For simplicity we'll just use a bufio.Scanner to parse our string since it
	// has a tokenization algorithm that breaks things into whitespace separated
	// words.  The only caveat here is that in our input parentheses aren't
	// necessarily surrounded by whitespace so won't be tokenized properly.  In
	// order to address this prior to handing the input off to the scanner we'll
	// replace them to ensure they have surrounding whitespace.
	input = strings.ReplaceAll(input, "(", " ( ")
	input = strings.ReplaceAll(input, ")", " ) ")

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	return &Lexer{scanner: scanner}
}

// Next consumes and returns the next token from the input.
func (l *Lexer) Next() Token {
	if l.next != nil {
		token := *l.next
		l.next = nil
		return token
	}

	if !l.scanner.Scan() {
		return Token{Kind: EOF}
	}

	switch literal := l.scanner.Text(); literal {
	case "(":
		return Token{Kind: LParen, Literal: literal}
	case ")":
		return Token{Kind: RParen, Literal: literal}
	case "+":
		return Token{Kind: Plus, Literal: literal}
	case "*":
		return Token{Kind: Times, Literal: literal}
	default:
		// Only other option is a number
		return Token{Kind: Number, Literal: literal}
	}
}

// Peek returns the next token from the input without consuming it.
func (l *Lexer) Peek() Token {
	if l.next == nil {
		token := l.Next()
		l.next = &token
	}

	return *l.next
}
