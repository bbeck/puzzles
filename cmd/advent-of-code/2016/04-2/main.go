package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"strings"
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
	SectorID int32
	Checksum string
}

func (r Room) Decrypt() string {
	bs := []byte(r.Name)
	for i := 0; i < len(bs); i++ {
		if bs[i] == '-' {
			continue
		}

		bs[i] = 'a' + byte((int32(bs[i]-'a')+r.SectorID)%26)
	}

	return string(bs)
}

func InputToRooms() []Room {
	return lib.InputLinesTo(func(line string) Room {
		hyphen := strings.LastIndex(line, "-")
		bracket := strings.LastIndex(line, "[")

		return Room{
			Name:     line[:hyphen],
			SectorID: int32(lib.ParseInt(line[hyphen+1 : bracket])),
			Checksum: line[bracket+1 : len(line)-1],
		}
	})
}
