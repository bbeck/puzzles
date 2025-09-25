package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	instructions := InputToInstructions()
	track := GetTrack()

	opponent := Run(instructions["A"], track)

	var wins int
	for ins := range UniquePermutations(strings.Split("+++++---===", "")) {
		if score := Run(ins, track); score > opponent {
			wins++
		}
	}
	fmt.Println(wins)
}

func Run(instructions []string, track []string) int {
	NI, NT := len(instructions), len(track)

	// The track has a length of 340 and the instructions have a length of 11.
	// The gcd(340, 11) = 1, so in theory we should cycle every 11 laps around
	// the track.  Therefore to figure out the winner, we only need to measure
	// who is winning after 11 laps instead of 2024.
	var essence int
	for power, step := 10, 0; step < 11*NT; step++ {
		switch ti := track[step%NT]; {
		case ti == "+" || (ti == "=" && instructions[step%NI] == "+"):
			power++
		case ti == "-" || (ti == "=" && instructions[step%NI] == "-"):
			power--
		}

		essence += power
	}

	return essence
}

func InputToInstructions() map[string][]string {
	instructions := make(map[string][]string)
	for in.HasNext() {
		lhs, rhs := in.Cut(":")
		instructions[lhs] = strings.Split(rhs, ",")
	}

	return instructions
}

func GetTrack() []string {
	s := `
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
	s = strings.ReplaceAll(s, "S", "=")
	s = strings.TrimSpace(s)
	grid := strings.Split(s, "\n")

	t := Turtle{Heading: Right}
	var start = t.Location

	var track []string
	for {
		t = Step(t, grid)
		track = append(track, string(grid[t.Location.Y][t.Location.X]))

		if t.Location == start {
			break
		}
	}
	return track
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
