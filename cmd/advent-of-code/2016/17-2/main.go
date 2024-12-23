package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

var Vault = lib.Point2D{X: 3, Y: 3}

func main() {
	start := InputToRoom()

	var longest string
	goal := func(r Room) bool {
		if r.Coordinate == Vault {
			longest = r.Passcode[len(start.Passcode):]
		}

		return false
	}

	lib.BreadthFirstSearch(start, Children, goal)
	fmt.Println(len(longest))
}

func Children(r Room) []Room {
	if r.Coordinate == Vault {
		return nil
	}

	open := map[byte]bool{'b': true, 'c': true, 'd': true, 'e': true, 'f': true}
	hash := Hash(r.Passcode)

	var children []Room
	if r.Coordinate.Y > 0 && open[hash[0]] {
		children = append(children, Room{
			Coordinate: r.Coordinate.Up(),
			Passcode:   r.Passcode + "U",
		})
	}
	if r.Coordinate.Y < 3 && open[hash[1]] {
		children = append(children, Room{
			Coordinate: r.Coordinate.Down(),
			Passcode:   r.Passcode + "D",
		})
	}
	if r.Coordinate.X > 0 && open[hash[2]] {
		children = append(children, Room{
			Coordinate: r.Coordinate.Left(),
			Passcode:   r.Passcode + "L",
		})
	}
	if r.Coordinate.X < 3 && open[hash[3]] {
		children = append(children, Room{
			Coordinate: r.Coordinate.Right(),
			Passcode:   r.Passcode + "R",
		})
	}

	return children
}

func Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

type Room struct {
	Coordinate lib.Point2D
	Passcode   string
}

func InputToRoom() Room {
	return Room{
		Passcode: lib.InputToString(),
	}
}
