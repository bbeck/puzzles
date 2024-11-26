package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"sort"
	"strings"
)

func main() {
	instructions := InputToInstructions()
	track := GetTrack()

	essences := make(map[string]int)
	for knight := range instructions {
		essences[knight] = Run(instructions[knight], track)
	}

	knights := Keys(essences)
	sort.Slice(knights, func(i, j int) bool {
		return essences[knights[i]] > essences[knights[j]]
	})

	fmt.Println(strings.Join(knights, ""))
}

func Run(instructions string, track string) int {
	var essence int
	power := 10

	for step := 0; step < 10*len(track); step++ {
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
S-=++=-==++=++=-=+=-=+=+=--=-=++=-==++=-+=-=+=-=+=+=++=-+==++=++=-=-=--
-                                                                     -
=                                                                     =
+                                                                     +
=                                                                     +
+                                                                     =
=                                                                     =
-                                                                     -
--==++++==+=+++-=+=-=+=-+-=+-=+-=+=-=+=--=+++=++=+++==++==--=+=++==+++-
`

func GetTrack() string {
	lines := strings.Split(strings.TrimSpace(Track), "\n")
	W, H := len(lines[0]), len(lines)

	var sb strings.Builder
	// top
	for x := 1; x < W; x++ {
		sb.WriteRune(rune(lines[0][x]))
	}
	// right
	for y := 1; y < H-1; y++ {
		sb.WriteRune(rune(lines[y][W-1]))
	}
	// bottom
	for x := W - 1; x >= 0; x-- {
		sb.WriteRune(rune(lines[H-1][x]))
	}
	// left
	for y := H - 2; y >= 0; y-- {
		sb.WriteRune(rune(lines[y][0]))
	}
	return sb.String()
}
