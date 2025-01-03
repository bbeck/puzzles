package main

import (
	"bytes"
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"log"
	"strings"
)

func main() {
	ops := InputToOperations()
	scramble := func(password []byte) []byte {
		for _, op := range ops {
			password = op(password)
		}
		return password
	}

	// Instead of determining an inverse scramble, we can just try all possible
	// passwords and see which one scrambles into our goal.
	goal := []byte("fbgdceah")

	password := make([]byte, len(goal))
	lib.EnumeratePermutations(len(goal), func(perm []int) bool {
		for i, index := range perm {
			password[i] = goal[index]
		}

		return bytes.Equal(scramble(password), goal)
	})

	fmt.Println(string(password))
}

type Operation func(bs []byte) []byte

func SwapPositions(x, y int) Operation {
	return func(bs []byte) []byte {
		bs[x], bs[y] = bs[y], bs[x]
		return bs
	}
}

func SwapLetters(x, y string) Operation {
	return func(bs []byte) []byte {
		bs = bytes.ReplaceAll(bs, []byte(x), []byte("_"))
		bs = bytes.ReplaceAll(bs, []byte(y), []byte(x))
		bs = bytes.ReplaceAll(bs, []byte("_"), []byte(y))
		return bs
	}
}

func RotateBasedOnPosition(x string) Operation {
	return func(bs []byte) []byte {
		index := bytes.IndexByte(bs, x[0])
		if index >= 4 {
			index += 1
		}
		index++

		x := index % len(bs)
		return append(bs[len(bs)-x:], bs[:len(bs)-x]...)
	}
}

func Rotate(dir string, x int) Operation {
	if dir == "left" {
		return func(bs []byte) []byte {
			return append(bs[x:], bs[:x]...)
		}
	} else {
		return func(bs []byte) []byte {
			return append(bs[len(bs)-x:], bs[:len(bs)-x]...)
		}
	}
}

func Reverse(x, y int) Operation {
	return func(bs []byte) []byte {
		for lo, hi := x, y; lo < hi; lo, hi = lo+1, hi-1 {
			bs[lo], bs[hi] = bs[hi], bs[lo]
		}
		return bs
	}
}

func Move(x, y int) Operation {
	return func(bs []byte) []byte {
		c := bs[x]
		bs = append(bs[:x], bs[x+1:]...) // remove c

		var cs []byte
		cs = append(cs, bs[:y]...) // copy the first y characters
		cs = append(cs, c)         // insert c
		if len(bs) > y {
			cs = append(cs, bs[y:]...) // copy the remainder
		}

		return cs
	}
}

func InputToOperations() []Operation {
	return lib.InputLinesTo(func(line string) Operation {
		tokens := strings.Fields(line)
		if tokens[0] == "swap" && tokens[1] == "position" {
			x := lib.ParseInt(tokens[2])
			y := lib.ParseInt(tokens[5])
			return SwapPositions(x, y)
		}

		if tokens[0] == "swap" && tokens[1] == "letter" {
			x := tokens[2]
			y := tokens[5]
			return SwapLetters(x, y)
		}

		if tokens[0] == "rotate" && (tokens[1] == "left" || tokens[1] == "right") {
			x := lib.ParseInt(tokens[2])
			return Rotate(tokens[1], x)
		}

		if tokens[0] == "rotate" && tokens[1] == "based" {
			x := tokens[6]
			return RotateBasedOnPosition(x)
		}

		if tokens[0] == "reverse" {
			x := lib.ParseInt(tokens[2])
			y := lib.ParseInt(tokens[4])
			return Reverse(x, y)
		}
		if tokens[0] == "move" {
			x := lib.ParseInt(tokens[2])
			y := lib.ParseInt(tokens[5])
			return Move(x, y)
		}

		log.Fatalf("unable to parse line: %s", line)
		return nil
	})
}
