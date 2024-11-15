package main

import (
	"errors"
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"github.com/bbeck/advent-of-code/puz/cpus"
	"math/rand"
	"regexp"
	"strings"
)

func main() {
	// The map is not symmetric -- there is at least one door that is directional
	// and doesn't have the same rooms on both sides.
	//
	// By exploring we've discovered several things:
	//
	//   1. Some doorways are not symmetric.  There is at least one door that
	//      doesn't return you to the prior room when you travel back through it.
	//
	//   2. There's a room with a pressure sensitive floor that you get trapped
	//      in if you don't have items weighing the correct amount.
	//
	//   3. The room prior to the pressure sensitive floor room is the security
	//      checkpoint.
	//
	//   4. There are 13 items available on the map.  5 of them end the run in
	//      some way when you pick them up.  So there are really only 8 items to
	//      care about.
	//
	// Since the map is quite small our approach will be to try all possible
	// combinations of items to pick up.
	var ignore puz.Set[string]
	ignore.Add("giant electromagnet")
	ignore.Add("escape pod")
	ignore.Add("molten lava")
	ignore.Add("photons")
	ignore.Add("infinite loop")

	var path []puz.Heading
	var items []string

	var lens []int
	for i := 0; i < 10; i++ {
		p, is := Explore(ignore)
		lens = append(lens, len(p))
		if len(path) == 0 || len(p) < len(path) {
			path = p
			items = is
		}
	}

	var code *string
	for k := 1; k < len(items) && code == nil; k++ {
		puz.EnumerateCombinations(len(items), k, func(ints []int) bool {
			var pickup puz.Set[string]
			for _, n := range ints {
				pickup.Add(items[n])
			}

			robot, current := NewRobot()
			for _, heading := range path {
				for item := range current.Items.Intersect(pickup) {
					robot.Take(item)
				}

				current = robot.Move(heading)
			}

			// Try moving into the pressure sensitive room
			current = robot.Move(puz.Right)

			if robot.Code != "" {
				code = &robot.Code
				return true
			}

			return false
		})
	}

	fmt.Println(*code)
}

func Explore(ignore puz.Set[string]) ([]puz.Heading, []string) {
	// We'll randomly walk through the map taking note of which item's we've
	// seen.  Once we end up in the security checkpoint after having seen all
	// 13 items then we'll return along that path.  We'll avoid the room with the
	// pressure sensitive floor as well since we'll get trapped in it without
	// the right items.  The pressure sensitive floor room is to the right of
	// the security checkpoint.
	var items puz.Set[string]
	var path []puz.Heading

	robot, current := NewRobot()
	for len(items) < 13 || current.ID != "Security Checkpoint" {
		items = items.Union(current.Items)

		doors := current.Doors
		if current.ID == "Security Checkpoint" {
			// The pressure sensitive floor is to the east.  Avoid it.
			doors = doors.DifferenceElems(puz.Right)
		}

		choices := doors.Entries()
		heading := choices[rand.Intn(len(choices))]
		path = append(path, heading)

		current = robot.Move(heading)
	}

	return path, items.Difference(ignore).Entries()
}

type Room struct {
	ID    string
	Doors puz.Set[puz.Heading]
	Items puz.Set[string]
}

type Robot struct {
	CPU     *cpus.IntcodeCPU
	Command chan int
	Rooms   chan Room
	Code    string
}

func NewRobot() (*Robot, Room) {
	command := make(chan int)
	rooms := make(chan Room)

	output := &strings.Builder{}

	var robot *Robot
	robot = &Robot{
		CPU: &cpus.IntcodeCPU{
			Memory: cpus.InputToIntcodeMemory(),
			Input: func() int {
				return <-command
			},
			Output: func(value int) {
				output.WriteByte(byte(value))

				if strings.HasSuffix(output.String(), "Command?") {
					if room, err := ParseRoom(output.String()); err == nil {
						rooms <- room
					}
					output.Reset()
				}

				if matches := CodeRegex.FindStringSubmatch(output.String()); len(matches) > 0 {
					robot.CPU.Stop()
					close(command)
					close(rooms)

					robot.Code = matches[1]
				}
			},
		},
		Command: command,
		Rooms:   rooms,
	}
	go robot.CPU.Execute()

	return robot, <-rooms
}

func (r *Robot) Move(heading puz.Heading) Room {
	var command string
	switch heading {
	case puz.Up:
		command = "north"
	case puz.Left:
		command = "west"
	case puz.Right:
		command = "east"
	case puz.Down:
		command = "south"
	}

	for _, c := range command {
		r.Command <- int(c)
	}
	r.Command <- 10

	return <-r.Rooms
}

func (r *Robot) Take(item string) {
	var command strings.Builder
	command.WriteString("take ")
	command.WriteString(item)

	for _, c := range command.String() {
		r.Command <- int(c)
	}
	r.Command <- 10
}

var IDRegex = regexp.MustCompile("== (.+) ==")
var ListRegex = regexp.MustCompile("- (.+)")
var CodeRegex = regexp.MustCompile("You should be able to get in by typing ([0-9]+) ")

func ParseRoom(s string) (Room, error) {
	if len(IDRegex.FindStringSubmatch(s)) == 0 {
		return Room{}, errors.New("no room id")
	}
	id := IDRegex.FindStringSubmatch(s)[1]

	var doors puz.Set[puz.Heading]
	var items puz.Set[string]
	for _, matches := range ListRegex.FindAllStringSubmatch(s, -1) {
		for _, match := range matches[1:] {
			switch match {
			case "north":
				doors.Add(puz.Up)
			case "west":
				doors.Add(puz.Left)
			case "east":
				doors.Add(puz.Right)
			case "south":
				doors.Add(puz.Down)
			default:
				items.Add(match)
			}
		}
	}

	return Room{ID: id, Doors: doors, Items: items}, nil
}
