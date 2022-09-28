package aoc

// PriorityQueue represents a min heap.  The zero value for PriorityQueue is
// an empty heap ready to use.
type PriorityQueue[T any] struct {
	values     []T
	priorities []int
}

// Empty returns true if the priority queue is empty.
func (q *PriorityQueue[T]) Empty() bool {
	return len(q.values) == 0
}

// Push adds a new value into the priority queue with the specified priority.
func (q *PriorityQueue[T]) Push(value T, priority int) {
	q.values = append(q.values, value)
	q.priorities = append(q.priorities, priority)
	q.up(len(q.values) - 1)
}

// Pop returns the value in the priority queue with the minimum priority.
func (q *PriorityQueue[T]) Pop() T {
	index := len(q.values) - 1
	q.swap(0, index)
	q.down(0, index)

	value := q.values[index]
	q.values = q.values[:index]
	q.priorities = q.priorities[:index]

	return value
}

func (q *PriorityQueue[T]) swap(i, j int) {
	q.values[i], q.values[j] = q.values[j], q.values[i]
	q.priorities[i], q.priorities[j] = q.priorities[j], q.priorities[i]
}

func (q *PriorityQueue[T]) up(index int) {
	for {
		parent := (index - 1) / 2 // parent
		if parent == index || q.priorities[index] > q.priorities[parent] {
			break
		}

		q.swap(index, parent)
		index = parent
	}
}

func (q *PriorityQueue[T]) down(i0, n int) {
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
