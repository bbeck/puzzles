package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/bbeck/advent-of-code/aoc"
)

const (
	_ int = iota
	Empty
	Wall
	Goal
)

func main() {
	grid := Explore(2_000_000)

	var goal aoc.Point2D
	for k, v := range grid {
		if v == Goal {
			goal = k
		}
	}

	// Index all of the open grid cells into our list of vertices.
	var vs []aoc.Point2D
	indices := make(map[aoc.Point2D]int)
	for k, v := range grid {
		if v == Empty || v == Goal {
			indices[k] = len(vs)
			vs = append(vs, k)
		}
	}

	// Helper to determine the neighbors of a given vertex, operating on indices.
	neighbors := func(i int) []int {
		v := vs[i]

		var ns []int
		if grid[v.Up()] != Wall {
			ns = append(ns, indices[v.Up()])
		}
		if grid[v.Down()] != Wall {
			ns = append(ns, indices[v.Down()])
		}
		if grid[v.Left()] != Wall {
			ns = append(ns, indices[v.Left()])
		}
		if grid[v.Right()] != Wall {
			ns = append(ns, indices[v.Right()])
		}

		return ns
	}

	// We'll use the Floyd-Warshall algorithm to determine the all pairs shortest
	// path to every empty grid cell from the goal location.
	var dist = make([][]int, len(vs))
	for i := 0; i < len(vs); i++ {
		dist[i] = make([]int, len(vs))
		for j := 0; j < len(vs); j++ {
			dist[i][j] = len(vs) * len(vs)
		}
	}

	for i := 0; i < len(vs); i++ {
		for _, j := range neighbors(i) {
			dist[i][j] = 1
		}
	}
	for i := 0; i < len(vs); i++ {
		dist[i][i] = 0
	}

	for k := 0; k < len(vs); k++ {
		for i := 0; i < len(vs); i++ {
			for j := 0; j < len(vs); j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	// Now figure out the furthest distance from the goal.
	var duration int
	gi := indices[goal]
	for vi := range vs {
		if dist[gi][vi] > duration {
			duration = dist[gi][vi]
		}
	}

	fmt.Println("duration:", duration)
}

func Explore(steps int) map[aoc.Point2D]int {
	grid := make(map[aoc.Point2D]int)

	// The currently location we're at as well as the next location we'll be at
	// if our move succeeds.
	var location, next aoc.Point2D

	cpu := &CPU{
		memory: InputToMemory(2019, 15),
	}

	cpu.input = func(int) int {
		dir := rand.Intn(4) + 1
		switch dir {
		case 1:
			next = location.Up()
		case 2:
			next = location.Down()
		case 3:
			next = location.Left()
		case 4:
			next = location.Right()
		default:
			log.Fatalf("invalid direction chosen: %d", dir)
		}

		return dir
	}

	cpu.output = func(value int) {
		switch value {
		case 0:
			grid[next] = Wall
		case 1:
			grid[next] = Empty
			location = next
		case 2:
			grid[next] = Goal
			location = next
		}

		// If we've captured the amount of data that we've wanted to we can
		// terminate the program and return our data.
		steps--
		if steps <= 0 {
			cpu.Stop()
		}
	}

	cpu.Execute()
	return grid
}
