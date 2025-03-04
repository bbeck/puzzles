package main

import (
	"fmt"
	"sort"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var sum int
	for _, room := range InputToRooms() {
		if room.IsReal() {
			sum += room.SectorID
		}
	}

	fmt.Println(sum)
}

type Room struct {
	Name     string
	SectorID int
	Checksum string
}

func (r Room) IsReal() bool {
	var counter FrequencyCounter[rune]
	for _, c := range r.Name {
		if c == '-' {
			continue
		}
		counter.Add(c)
	}

	entries := counter.Entries()
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Count == entries[j].Count {
			return entries[i].Value < entries[j].Value
		}
		return entries[i].Count > entries[j].Count
	})

	var sb strings.Builder
	for _, entry := range entries[:5] {
		sb.WriteRune(entry.Value)
	}

	return sb.String() == r.Checksum
}

func InputToRooms() []Room {
	return in.LinesToS(func(in in.Scanner[Room]) Room {
		var letters, checksum string
		var number int
		in.Scanf("%s-%d[%s]", &letters, &number, &checksum)

		return Room{
			Name:     strings.ReplaceAll(letters, "-", ""),
			SectorID: number,
			Checksum: checksum,
		}
	})
}
