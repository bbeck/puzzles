package aoc

type PriorityQueue struct {
	nodes      []interface{}
	priorities []int
}

// NewPriorityQueue creates a new, empty priority queue.
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		nodes:      make([]interface{}, 0),
		priorities: make([]int, 0),
	}
}

// Empty returns true if the priority queue is empty.
func (q *PriorityQueue) Empty() bool {
	return len(q.nodes) == 0
}

// Push adds a new value into the priority queue with the specified priority.
func (q *PriorityQueue) Push(node interface{}, priority int) {
	q.nodes = append(q.nodes, node)
	q.priorities = append(q.priorities, priority)
	q.up(len(q.nodes) - 1)
}

// Pop returns the value in the priority queue with the minimum priority.
func (q *PriorityQueue) Pop() interface{} {
	index := len(q.nodes) - 1
	q.swap(0, index)
	q.down(0, index)

	node := q.nodes[index]
	q.nodes[index] = nil
	q.nodes = q.nodes[:index]
	q.priorities = q.priorities[:index]

	return node
}

func (q *PriorityQueue) swap(i, j int) {
	q.nodes[i], q.nodes[j] = q.nodes[j], q.nodes[i]
	q.priorities[i], q.priorities[j] = q.priorities[j], q.priorities[i]
}

func (q *PriorityQueue) up(index int) {
	for {
		parent := (index - 1) / 2 // parent
		if parent == index || q.priorities[index] > q.priorities[parent] {
			break
		}

		q.swap(index, parent)
		index = parent
	}
}

func (q *PriorityQueue) down(i0, n int) {
	index := i0
	for {
		left := 2*index + 1
		if left >= n || left < 0 {
			break
		}

		smaller := left
		if right := left + 1; right < n && q.priorities[right] < q.priorities[left] {
			smaller = right
		}

		if q.priorities[smaller] > q.priorities[index] {
			break
		}

		q.swap(index, smaller)
		index = smaller
	}
}
