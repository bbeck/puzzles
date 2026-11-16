package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToStringGrid2D()

	var start, end Point2D
	var stops = make(map[Point2D]int)
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "S" {
			start = p
		}
		if s == "E" {
			end = p
		}
		if s == "@" {
			stops[p] = len(stops)
		}
	})

	type State struct {
		Pose      Pose2D
		Collected BitSet // There are less than 64 stops
	}

	children := func(s State) []State {
		candidates := []Pose2D{
			s.Pose.Forward(),
			s.Pose.TurnRight().Forward(),
			s.Pose.TurnLeft().Forward(),
		}

		var children []State
		for _, child := range candidates {
			if !grid.InBoundsPoint(child.Point) || grid.GetPoint(child.Point) == "#" {
				continue
			}

			var collected = s.Collected
			if id, ok := stops[child.Point]; ok {
				collected = collected.Add(id)
			}

			children = append(children, State{Pose: child, Collected: collected})
		}
		return children
	}

	goal := func(s State) bool {
		return s.Pose.Point == end && s.Collected.Size() == len(stops)
	}

	steps, _ := BreadthFirstSearch(State{Pose: Pose2D{Point: start, Heading: Right}}, children, goal)
	fmt.Println(len(steps) - 1)
}
