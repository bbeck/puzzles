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

func Run(instructions []string, track []string) int {
	NI, NT := len(instructions), len(track)

	var essence int
	for power, step := 10, 0; step < 10*NT; step++ {
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
	for _, line := range InputToLines() {
		line = strings.ReplaceAll(line, ":", " ")
		line = strings.ReplaceAll(line, ",", " ")
		fields := strings.Fields(line)

		instructions[fields[0]] = fields[1:]
	}

	return instructions
}

func GetTrack() []string {
	s := `
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
	s = strings.ReplaceAll(s, "S", "=")
	s = strings.TrimSpace(s)
	grid := strings.Split(s, "\n")

	W, H := len(grid[0]), len(grid)

	var track []string

	// top
	for x := 1; x < W; x++ {
		track = append(track, grid[0][x:x+1])
	}
	// right
	for y := 1; y < H-1; y++ {
		track = append(track, grid[y][W-1:W])
	}
	// bottom
	for x := W - 1; x >= 0; x-- {
		track = append(track, grid[H-1][x:x+1])
	}
	// left
	for y := H - 2; y >= 0; y-- {
		track = append(track, grid[y][0:1])
	}
	return track
}
