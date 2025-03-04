package main

import (
	"bytes"
	"fmt"
	"slices"

	"github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
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

func ReversePositions(x, y int) Operation {
	return func(bs []byte) []byte {
		for lo, hi := x, y; lo < hi; lo, hi = lo+1, hi-1 {
			bs[lo], bs[hi] = bs[hi], bs[lo]
		}
		return bs
	}
}

func MovePosition(x, y int) Operation {
	return func(bs []byte) []byte {
		c := bs[x]
		bs = slices.Delete(bs, x, x+1) // remove c

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
	return in.LinesToS(func(in in.Scanner[Operation]) Operation {
		switch {
		case in.HasPrefix("swap position"):
			var x, y int
			in.Scanf("swap position %d with position %d", &x, &y)
			return SwapPositions(x, y)

		case in.HasPrefix("swap letter"):
			var x, y string
			in.Scanf("swap letter %s with letter %s", &x, &y)
			return SwapLetters(x, y)

		case in.HasPrefix("rotate left"):
			var x int
			in.Scanf("rotate left %d", &x)
			return Rotate("left", x)

		case in.HasPrefix("rotate right"):
			var x int
			in.Scanf("rotate right %d", &x)
			return Rotate("right", x)

		case in.HasPrefix("rotate based"):
			var x string
			in.Scanf("rotate based on position of letter %s", &x)
			return RotateBasedOnPosition(x)

		case in.HasPrefix("reverse positions"):
			var x, y int
			in.Scanf("reverse positions %d through %d", &x, &y)
			return ReversePositions(x, y)

		case in.HasPrefix("move position"):
			var x, y int
			in.Scanf("move position %d to position %d", &x, &y)
			return MovePosition(x, y)

		default:
			panic("unknown operation")
		}
	})
}
