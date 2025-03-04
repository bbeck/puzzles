package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var storage Room
	for _, room := range InputToRooms() {
		if room.Decrypt() == "northpole-object-storage" {
			storage = room
			break
		}
	}

	fmt.Println(storage.SectorID)
}

type Room struct {
	Name     string
	SectorID int
	Checksum string
}

func (r Room) Decrypt() string {
	bs := []byte(r.Name)
	for i := range bs {
		if bs[i] == '-' {
			continue
		}

		bs[i] = 'a' + byte((int(bs[i]-'a')+r.SectorID)%26)
	}

	return string(bs)
}

func InputToRooms() []Room {
	return in.LinesToS(func(in in.Scanner[Room]) Room {
		var letters, checksum string
		var number int
		in.Scanf("%s-%d[%s]", &letters, &number, &checksum)

		return Room{
			Name:     letters,
			SectorID: number,
			Checksum: checksum,
		}
	})
}
