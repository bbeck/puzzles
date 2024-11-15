package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

var Vault = puz.Point2D{X: 3, Y: 3}

func main() {
	start := InputToRoom()

	goal := func(r Room) bool {
		return r.Coordinate == Vault
	}

	cost := func(from, to Room) int {
		return 1
	}

	heuristic := func(r Room) int {
		return Vault.ManhattanDistance(r.Coordinate)
	}

	path, _, found := puz.AStarSearch(start, Children, goal, cost, heuristic)
	if !found {
		fmt.Println("no path found")
		return
	}

	passcode := path[len(path)-1].Passcode
	shortest := passcode[len(start.Passcode):]
	fmt.Println(shortest)
}

func Children(r Room) []Room {
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
	Coordinate puz.Point2D
	Passcode   string
}

func InputToRoom() Room {
	return Room{
		Passcode: puz.InputToString(2016, 17),
	}
}
