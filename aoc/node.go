package aoc

// Node is a type that represents a node that can be searched from.  A node
// must have a unique identifier -- as the ID() method is used to define unique
// identity for a node.  Children() returns each of the reachable children for
// a node.
type Node interface {
	ID() string
	Children() []Node
}

// GoalFunc is a function that is called the first time each node is
// visited.  Its return value is interpreted as whether or not the search has
// reached the goal or not.  If the search reaches the goal it will terminate.
type GoalFunc func(node Node) bool

// BreadthFirstSearch performs a search starting at the provided root node and
// calls the visit function the first time each node is encountered.  The visit
// function determines whether the search should continue or not.  If the visit
// function returns true then the goal has been found and the search should
// terminate.  If it returns false then the search will continue on as long as
// there are remaining children present.
//
// The BreadthFirstSearch function returns two values.  The first is the path
// from the root node to the goal.  The second is a boolean indicating whether
// a path was found.
func BreadthFirstSearch(root Node, isGoal GoalFunc) ([]Node, bool) {
	queue := []Node{root}
	parents := make(map[string]Node)
	seen := map[string]bool{
		root.ID(): true,
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if isGoal(current) {
			path := path(current, parents)
			return path, true
		}

		for _, child := range current.Children() {
			if !seen[child.ID()] {
				seen[child.ID()] = true
				parents[child.ID()] = current
				queue = append(queue, child)
			}
		}
	}

	return nil, false
}

// CostFunc determines the cost of transitioning from one node to one of its
// child nodes.
type CostFunc func(from, to Node) int

// HeuristicFunc provides an estimate of the cost to reach the goal node from a
// given node.
type HeuristicFunc func(node Node) int

// AStarSearch utilizes the A* algorithm to find the shortest path from the
// start node to the goal node in a graph.  The search utilizes a heuristic
// function to aid in making the search run faster.
//
// The AStarSearch function returns three values.  The first is the path from
// the start node to the goal node.  The second is the cost of that path, and
// the third is a boolean indicating whether a path was found.
func AStarSearch(start Node, isGoal GoalFunc, c CostFunc, h HeuristicFunc) ([]Node, int, bool) {
	frontier := NewPriorityQueue()
	frontier.Push(start, 0)

	parents := make(map[string]Node)

	costs := make(map[string]int)
	costs[start.ID()] = 0

	nodes := make(map[string]Node)
	nodes[start.ID()] = start

	for !frontier.Empty() {
		current := frontier.Pop().(Node)
		cid := current.ID()
		ccost := costs[cid]
		nodes[cid] = current

		if isGoal(current) {
			return path(current, parents), costs[cid], true
		}

		for _, next := range current.Children() {
			ncost := c(current, next)
			nid := next.ID()
			if _, ok := nodes[nid]; ok {
				continue
			}

			if oldCost, ok := costs[nid]; !ok || ccost+ncost < oldCost {
				costs[nid] = ccost + ncost
				parents[nid] = current
				frontier.Push(next, ccost+ncost+h(next))
			}
		}
	}

	return nil, 0, false
}

// path will build the path from a goal node all the way back to the start
// node (the node that has no parent) and return it.  The returned path is in
// the order from start to goal.
func path(node Node, parents map[string]Node) []Node {
	var path []Node
	for node != nil {
		path = append([]Node{node}, path...)
		node = parents[node.ID()]
	}

	return path
}
