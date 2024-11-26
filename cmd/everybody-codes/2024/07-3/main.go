package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	instructions := InputToInstructions()
	track := GetTrack()

	opponent := Run(instructions["A"], track)
	chs := []string{"+", "+", "+", "+", "+", "-", "-", "-", "=", "=", "="}

	var seen Set[string]
	var wins int
	EnumeratePermutations(len(chs), func(ns []int) bool {
		elems := make([]string, len(chs))
		for i, n := range ns {
			elems[i] = chs[n]
		}

		ins := strings.Join(elems, "")
		if seen.Add(ins) {
			score := Run(ins, track)
			if score > opponent {
				wins++
			}
		}

		return false
	})

	fmt.Println(wins)
}

func Run(instructions string, track string) int {
	var essence int
	power := 10

	for step := 0; step < 2024*len(track); step++ {
		ta := string(track[step%len(track)])
		ka := string(instructions[step%len(instructions)])

		var aa string
		if ta == "+" || ta == "-" {
			aa = ta
		} else {
			aa = ka
		}

		switch aa {
		case "+":
			power++
		case "-":
			power--
		}
		essence += power
	}

	return essence
}

func InputToInstructions() map[string]string {
	instructions := make(map[string]string)
	for _, line := range InputToLines() {
		lhs, rhs, _ := strings.Cut(line, ":")
		rhs = strings.ReplaceAll(rhs, ",", "")
		instructions[lhs] = rhs
	}

	return instructions
}

const Track = `
S+= +=-== +=++=     =+=+=--=    =-= ++=     +=-  =+=++=-+==+ =++=-=-=--
- + +   + =   =     =      =   == = - -     - =  =         =-=        -
= + + +-- =-= ==-==-= --++ +  == == = +     - =  =    ==++=    =++=-=++
+ + + =     +         =  + + == == ++ =     = =  ==   =   = =++=
= = + + +== +==     =++ == =+=  =  +  +==-=++ =   =++ --= + =
+ ==- = + =   = =+= =   =       ++--          +     =   = = =--= ==++==
=     ==- ==+-- = = = ++= +=--      ==+ ==--= +--+=-= ==- ==   =+=    =
-               = = = =   +  +  ==+ = = +   =        ++    =          -
-               = + + =   +  -  = + = = +   =        +     =          -
--==++++==+=+++-= =-= =-+-=  =+-= =-= =--   +=++=+++==     -=+=++==+++-
`

func GetTrack() string {
	lines := strings.Split(strings.TrimSpace(Track), "\n")

	var sb strings.Builder
	t := Turtle{Heading: Right}

	for {
		next := Step(t, lines)
		ch := rune(lines[next.Location.Y][next.Location.X])
		sb.WriteRune(ch)

		if ch == 'S' {
			break
		}

		t = next
	}
	return sb.String()
}

func Step(t Turtle, grid []string) Turtle {
	// forward
	{
		u := t
		u.Forward(1)
		x, y := u.Location.X, u.Location.Y

		if 0 <= y && y < len(grid) && 0 <= x && x < len(grid[y]) && grid[y][x] != ' ' {
			return u
		}
	}

	// right
	{
		u := t
		u.TurnRight()
		u.Forward(1)
		x, y := u.Location.X, u.Location.Y

		if 0 <= y && y < len(grid) && 0 <= x && x < len(grid[y]) && grid[y][x] != ' ' {
			return u
		}
	}

	// left
	{
		u := t
		u.TurnLeft()
		u.Forward(1)
		x, y := u.Location.X, u.Location.Y

		if 0 <= y && y < len(grid) && 0 <= x && x < len(grid[y]) && grid[y][x] != ' ' {
			return u
		}
	}

	return t // should never happen
}
