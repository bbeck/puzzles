package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	for _, room := range InputToRooms(2016, 4) {
		if room.IsReal() && room.DecryptID() == "northpole-object-storage" {
			fmt.Printf("sectore: %d\n", room.sector)
			break
		}
	}
}

type Room struct {
	id       string
	sector   int
	checksum string
}

func InputToRooms(year, day int) []Room {
	var regex = regexp.MustCompile(`^([a-z-]+)-([0-9]+)\[([a-z]+)]`)

	var rooms []Room
	for _, line := range aoc.InputToLines(year, day) {
		matches := regex.FindStringSubmatch(line)
		rooms = append(rooms, Room{matches[1], aoc.ParseInt(matches[2]), matches[3]})
	}

	return rooms
}

func (r Room) IsReal() bool {
	letters := make(map[string]int)
	for _, c := range strings.ReplaceAll(r.id, "-", "") {
		letters[string(c)]++
	}

	frequencies := make([]struct {
		letter string
		count  int
	}, 0)
	for letter, count := range letters {
		frequencies = append(frequencies, struct {
			letter string
			count  int
		}{
			letter: letter,
			count:  count,
		})
	}

	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[j].count < frequencies[i].count ||
			(frequencies[j].count == frequencies[i].count && frequencies[i].letter < frequencies[j].letter)
	})

	if len(frequencies) < 5 {
		return false
	}

	for i := 0; i < 5; i++ {
		if frequencies[i].letter != string(r.checksum[i]) {
			return false
		}
	}

	return true
}

func (r Room) DecryptID() string {
	bs := []byte(r.id)
	for i := 0; i < len(bs); i++ {
		if bs[i] != '-' {
			bs[i] = byte((int(bs[i]-'a')+r.sector)%26) + 'a'
		}
	}

	return string(bs)
}
