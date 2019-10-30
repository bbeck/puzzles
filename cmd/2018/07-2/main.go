package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	graph := InputToGraph(2018, 7)

	workers := make([]*Worker, 0)
	for n := 1; n <= 5; n++ {
		workers = append(workers, &Worker{id: n})
	}

	durations := make(map[string]int)
	for i, c := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		durations[string(c)] = 61 + i
	}

	todo := TopologicalSort(graph.Copy()) // the work items still left to do
	complete := make(map[string]bool)     // the work items that are complete
	var tm int
	for tm = 0; ; tm++ {
		// See if we're out of work to do and all workers are idle.
		var busy int
		for _, worker := range workers {
			if worker.node != "" {
				busy++
			}
		}

		if len(todo) == 0 && busy == 0 {
			break
		}

		// Move each worker along
		for _, worker := range workers {
			if worker.node == "" {
				continue
			}

			worker.duration--
			if worker.duration == 0 {
				complete[worker.node] = true
				worker.node = ""
			}
		}

		// Attempt to assign work to idle workers.
		for _, worker := range workers {
			// If there's no work left then move on, nothing to assign.
			if len(todo) == 0 {
				break
			}

			// See if this worker is already busy.
			if worker.node != "" {
				continue
			}

			// See if there is a piece of work that can be started.  We do this in
			// the worker loop so that we can assign multiple pieces of work to
			// different workers at the same time slice.
			for i, work := range todo {
				ready := true
				for _, parent := range graph.Parents(work) {
					if !complete[parent] {
						ready = false
					}
				}

				if !ready {
					continue
				}

				todo = append(todo[0:i], todo[i+1:]...)

				worker.node = work
				worker.duration = durations[work]
				break
			}
		}
	}

	fmt.Printf("duration: %d\n", tm-1)
}

type Worker struct {
	id       int
	node     string // the node in the graph they're currently working on
	duration int    // the amount of time that a worker has left before being complete
}

type Graph struct {
	vertices map[string]int
	matrix   [][]bool
}

func (g *Graph) Copy() *Graph {
	matrix := make([][]bool, len(g.matrix))
	for i := 0; i < len(matrix); i++ {
		matrix[i] = append([]bool{}, g.matrix[i]...)
	}

	return &Graph{
		vertices: g.vertices,
		matrix:   matrix,
	}
}

func (g *Graph) Parents(to string) []string {
	tid := g.vertices[to]

	var parents []string
	for from, fid := range g.vertices {
		if g.matrix[fid][tid] {
			parents = append(parents, from)
		}
	}

	return parents
}

func (g *Graph) Children(from string) []string {
	fid := g.vertices[from]

	var children []string
	for to, tid := range g.vertices {
		if g.matrix[fid][tid] {
			children = append(children, to)
		}
	}

	return children
}

// Compute the topological sort of the graph using Kahn's algorithm -- this
// method will modify the graph.
func TopologicalSort(g *Graph) []string {
	var L []string // the topological sort output

	var S []string // set of nodes that have no incoming edge
	for name := range g.vertices {
		if len(g.Parents(name)) == 0 {
			S = append(S, name)
		}
	}

	for len(S) > 0 {
		// ensure S is sorted
		sort.Strings(S)

		// remove a node n from S
		n := S[0]
		S = S[1:]

		// add n to tail of L
		L = append(L, n)

		// for each node m with an edge e from n to m
		for _, m := range g.Children(n) {
			// remove edge e from the graph
			g.matrix[g.vertices[n]][g.vertices[m]] = false

			// if m has no other incoming edges then
			if len(g.Parents(m)) == 0 {
				// insert m into S
				S = append(S, m)
			}
		}
	}

	return L
}

func InputToGraph(year, day int) *Graph {
	assigned := make(map[string]bool)  // each vertex that has been assigned an id
	vertices := make(map[string]int)   // the assigned id for each vertex
	edges := make(map[string][]string) // the mapping of edges from vertex to vertex

	for _, line := range aoc.InputToLines(year, day) {
		var from, to string
		if _, err := fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &from, &to); err != nil {
			log.Fatalf("unable to parse line: %s", line)
		}

		if !assigned[from] {
			vertices[from] = len(vertices)
			assigned[from] = true
		}
		if !assigned[to] {
			vertices[to] = len(vertices)
			assigned[to] = true
		}

		edges[from] = append(edges[from], to)
	}

	matrix := make([][]bool, len(vertices))
	for name, id := range vertices {
		matrix[id] = make([]bool, len(vertices))
		for _, child := range edges[name] {
			matrix[id][vertices[child]] = true
		}
	}

	return &Graph{
		vertices: vertices,
		matrix:   matrix,
	}
}
