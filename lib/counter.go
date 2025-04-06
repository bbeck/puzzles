package lib

import (
	"container/heap"
)

type FrequencyCounter[T comparable] struct {
	// Map of entries by value for fast lookup when updating a value's counter.
	entries map[T]*FrequencyCount[T]

	// Heap of entries with the root being the most frequent entry.
	heap *fcHeap[T]
}

// AddWithCount adds a value to the frequency counter with a specific count.
// If the value already exists in the frequency counter then it's count will be
// incremented by the count.
func (fc *FrequencyCounter[T]) AddWithCount(value T, count int) {
	// Lazily initialize the entries map.
	if fc.entries == nil {
		fc.entries = make(map[T]*FrequencyCount[T])
	}
	if fc.heap == nil {
		var h fcHeap[T]
		fc.heap = &h
	}

	// First check if this value is already in the frequency counter, if it is
	// then just update its count and repair the heap.
	if entry, found := fc.entries[value]; found {
		entry.Count += count
		heap.Fix(fc.heap, entry.index)
		return
	}

	// This value wasn't in the frequency counter, create a new entry for it
	// and add it to the heap.
	entry := &FrequencyCount[T]{
		Value: value,
		Count: count,
		index: fc.heap.Len(),
	}
	fc.entries[value] = entry
	heap.Push(fc.heap, entry)
}

// Add adds a value to the frequency counter.  If the value already exists in
// the frequency counter then it's count will be incremented by 1.
func (fc *FrequencyCounter[T]) Add(value T) {
	fc.AddWithCount(value, 1)
}

// GetCount returns the count of a specific value within the frequency counter.
func (fc *FrequencyCounter[T]) GetCount(value T) int {
	if fc.entries == nil {
		fc.entries = make(map[T]*FrequencyCount[T])
	}

	if entry, found := fc.entries[value]; found {
		return entry.Count
	}
	return 0
}

// Entries returns the entries within the frequency counter in order of
// frequency from most frequent to least frequent.
func (fc *FrequencyCounter[T]) Entries() []FrequencyCount[T] {
	if fc.heap == nil {
		return nil
	}

	entries := make([]FrequencyCount[T], 0, fc.heap.Len())
	for dup := fc.heap.Copy(); dup.Len() > 0; {
		entry := heap.Pop(&dup).(*FrequencyCount[T])
		entries = append(entries, *entry)
	}

	return entries
}

type FrequencyCount[T any] struct {
	Value T
	Count int

	// Index in the heap array that this entry appears.
	index int
}

type fcHeap[T any] []*FrequencyCount[T]

func (h fcHeap[T]) Len() int {
	return len(h)
}
func (h fcHeap[T]) Less(i, j int) bool {
	// We use greater than here so that the root of the heap is the most
	// frequent entry.
	return h[i].Count > h[j].Count
}

func (h fcHeap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *fcHeap[T]) Push(x any) {
	entry := x.(*FrequencyCount[T])
	entry.index = len(*h)
	*h = append(*h, entry)
}

func (h *fcHeap[T]) Pop() any {
	n := len(*h)

	entry := (*h)[n-1]
	entry.index = -1

	(*h)[n-1] = nil
	*h = (*h)[0 : n-1]

	return entry
}

func (h fcHeap[T]) Copy() fcHeap[T] {
	dst := make(fcHeap[T], h.Len())
	copy(dst, h)
	return dst
}
