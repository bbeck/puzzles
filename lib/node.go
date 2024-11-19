package lib

// ChildrenFunc is a function that is called to determine the children of a
// node.
type ChildrenFunc[T any] func(node T) []T

// GoalFunc is a function that is called the first time a node is visited.
// Its return value is interpreted as whether the search has reached the goal
// or not.  If the search reaches the goal it will terminate.
type GoalFunc[T any] func(node T) bool

// IDFunc converts a node into a comparable for identity checking.
type IDFunc[T any, I comparable] func(T) I

// BreadthFirstSearch performs a search starting at the provided root node and
// calls the isGoal function the first time each node is encountered.  The goal
// function determines whether the search should continue or not.  If the goal
// function returns true then the goal has been found and the search should
// terminate.  If it returns false then the search will continue on as long as
// there are remaining children present.
//
// The BreadthFirstSearch function returns two values.  The first is the path
// from the root node to the goal.  The second is a boolean indicating whether
// a path was found.
func BreadthFirstSearch[T comparable](root T, children ChildrenFunc[T], goal GoalFunc[T]) ([]T, bool) {
	return BreadthFirstSearchWithIdentity(root, children, goal, func(t T) T { return t })
}

// BreadthFirstSearchWithIdentity performs a breadth first search as described
// by BreadthFirstSearch, but allows an id function to be used to determine
// node identity.  This is useful for allowing only a portion of a node object
// to be used for comparison against other nodes.
func BreadthFirstSearchWithIdentity[T, I comparable](root T, children ChildrenFunc[T], goal GoalFunc[T], id IDFunc[T, I]) ([]T, bool) {
	queue := []T{root}
	parents := make(map[T]T)
	seen := SetFrom(id(root))

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if goal(current) {
			return path(current, parents), true
		}

		for _, child := range children(current) {
			if seen.Add(id(child)) {
				parents[child] = current
				queue = append(queue, child)
			}
		}
	}

	return nil, false
}

// CostFunc determines the cost of transitioning from one node to one of its
// child nodes.
type CostFunc[T any] func(from, to T) int

// HeuristicFunc provides an estimate of the cost to reach the goal node from a
// given node.
type HeuristicFunc[T any] func(node T) int

// AStarSearch utilizes the A* algorithm to find the shortest path from the
// start node to the goal node in a graph.  The search utilizes a heuristic
// function to aid in making the search run faster.  The heuristic function
// may underestimate the cost to reach the goal node at the expense of
// exploring more of the search space, but it should NOT overestimate the
// cost.  If the heuristic overestimates the cost to the goal, the search
// may not find the shortest path.
//
// The AStarSearch function returns three values.  The first is the path from
// the start node to the goal node.  The second is the cost of that path, and
// the third is a boolean indicating whether a path was found.
func AStarSearch[T comparable](start T, children ChildrenFunc[T], g GoalFunc[T], c CostFunc[T], h HeuristicFunc[T]) ([]T, int, bool) {
	return AStarSearchWithIdentity(start, children, g, c, h, func(t T) T { return t })
}

// AStarSearchWithIdentity performs an A* search as described by AStarSearch,
// but allows an id function to be used to determine node identity.  This is
// useful for allowing only a portion of a node object to be used for
// comparison against other nodes.
func AStarSearchWithIdentity[T, I comparable](start T, children ChildrenFunc[T], g GoalFunc[T], c CostFunc[T], h HeuristicFunc[T], id IDFunc[T, I]) ([]T, int, bool) {
	var frontier PriorityQueue[T]
	frontier.Push(start, 0)

	parents := make(map[T]T)

	costs := make(map[T]int)
	costs[start] = 0

	seen := SetFrom(id(start))

	for !frontier.Empty() {
		current := frontier.Pop()
		ccost := costs[current]
		seen.Add(id(current))

		if g(current) {
			return path(current, parents), ccost, true
		}

		for _, next := range children(current) {
			if seen.Contains(id(next)) {
				continue
			}

			ncost := c(current, next)
			if oldCost, ok := costs[next]; !ok || ccost+ncost < oldCost {
				costs[next] = ccost + ncost
				parents[next] = current
				frontier.Push(next, ccost+ncost+h(next))
			}
		}
	}

	return nil, 0, false
}

// path will build the path from a goal node all the way back to the start
// node (the node that has no parent) and return it.  The returned path is in
// the order from start to goal.
func path[T comparable](node T, parents map[T]T) []T {
	var path []T
	for {
		path = append([]T{node}, path...)
		parent, present := parents[node]
		if !present {
			break
		}
		node = parent
	}

	return path
}
