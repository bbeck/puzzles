package aoc

// This implementation is heavily influenced by the builtin deque type in
// Python which uses a linked list of blocks based approach.  It also
// supports an efficient rotate operation which allows the deque to be used
// as the underlying data structure for a ring in addition to the traditional
// stack and queue.
//
// https://github.com/python/cpython/blob/main/Modules/_collectionsmodule.c

// A Deque is a container of elements where elements can be added or removed
// from the front or back of the deque.
type Deque[T any] struct {
	len               int
	left, right       *Block[T]
	leftIdx, rightIdx int
}

const BlockSize = 64
const Center = (BlockSize - 1) / 2

type Block[T any] struct {
	left, right *Block[T]
	data        [BlockSize]T
}

// Empty determines if there are any elements in the deque.
func (d *Deque[T]) Empty() bool {
	return d.len == 0
}

// Len returns the number of elements in the deque.
func (d *Deque[T]) Len() int {
	return d.len
}

// PeekFront returns the element at the front of the deque but does not remove
// it.  If the deque is empty then the zero value is returned.
func (d *Deque[T]) PeekFront() T {
	d.initialize()
	return d.left.data[d.leftIdx]
}

// PeekBack returns the element at the back of the deque but does not remove
// it.  If the deque is empty then the zero value is returned.
func (d *Deque[T]) PeekBack() T {
	d.initialize()
	return d.right.data[d.rightIdx]
}

// PushFront adds an element to the front of the deque.
func (d *Deque[T]) PushFront(elem T) {
	d.initialize()

	// Create a new block, if necessary, to hold the new element.
	if d.leftIdx == 0 {
		block := &Block[T]{}
		block.right = d.left

		d.left.left = block
		d.left = block
		d.leftIdx = BlockSize
	}

	// Insert this element into the left-most block.
	d.leftIdx--
	d.left.data[d.leftIdx] = elem
	d.len++
}

// PushBack adds an element to the back of the deque.
func (d *Deque[T]) PushBack(elem T) {
	d.initialize()

	// Create a new block, if necessary, to hold the new element.
	if d.rightIdx == BlockSize-1 {
		block := &Block[T]{}
		block.left = d.right

		d.right.right = block
		d.right = block
		d.rightIdx = -1
	}

	// Insert this element into the right-most block.
	d.rightIdx++
	d.right.data[d.rightIdx] = elem
	d.len++
}

// PopFront returns the element from the front of the deque and removes it
// from the deque.  If the deque is empty then the zero value is returned.
func (d *Deque[T]) PopFront() T {
	d.initialize()

	var zero T
	if d.len == 0 {
		return zero
	}

	elem := d.left.data[d.leftIdx]
	d.left.data[d.leftIdx] = zero
	d.leftIdx++
	d.len--

	if d.leftIdx == BlockSize {
		if d.len > 0 {
			d.left = d.left.right
			d.left.left = nil
			d.leftIdx = 0
		} else {
			d.leftIdx = Center + 1
			d.rightIdx = Center
		}
	}

	return elem
}

// PopBack returns the element from the back of the deque and removes it
// from the deque.  If the deque is empty then the zero value is returned.
func (d *Deque[T]) PopBack() T {
	d.initialize()

	var zero T
	if d.len == 0 {
		return zero
	}

	elem := d.right.data[d.rightIdx]
	d.right.data[d.rightIdx] = zero
	d.rightIdx--
	d.len--

	if d.rightIdx < 0 {
		if d.len > 0 {
			d.right = d.right.left
			d.right.right = nil
			d.rightIdx = BlockSize - 1
		} else {
			d.leftIdx = Center + 1
			d.rightIdx = Center
		}
	}

	return elem
}

// Rotate will rotate the elements in the deque by the provided amount.  If n
// is positive then the rotation will be from back to front (clockwise).  If n
// is negative then the rotation will be from front to back (anticlockwise).
func (d *Deque[T]) Rotate(n int) {
	L := d.len

	if L <= 1 {
		return
	}

	if n > L/2 || n < -(L/2) {
		n %= L
		if n > L/2 {
			n -= L
		} else if n < -(L / 2) {
			n += L
		}
	}

	var b *Block[T]
	left := d.left
	right := d.right
	leftIdx := d.leftIdx
	rightIdx := d.rightIdx

	for n > 0 {
		if leftIdx == 0 {
			if b == nil {
				b = &Block[T]{}
			}

			b.right = left
			left.left = b
			left = b
			leftIdx = BlockSize
			b = nil
		}

		m := n
		if m > rightIdx+1 {
			m = rightIdx + 1
		}
		if m > leftIdx {
			m = leftIdx
		}
		rightIdx -= m
		leftIdx -= m
		n -= m
		for i := 0; m > 0; i++ {
			left.data[leftIdx+i] = right.data[rightIdx+i+1]
			m--
		}

		if rightIdx < 0 {
			b = right
			right = right.left
			rightIdx = BlockSize - 1
		}
	}

	for n < 0 {
		if rightIdx == BlockSize-1 {
			if b == nil {
				b = &Block[T]{}
			}

			b.left = right
			right.right = b
			right = b
			rightIdx = -1
			b = nil
		}

		m := -n
		if m > BlockSize-leftIdx {
			m = BlockSize - leftIdx
		}
		if m > BlockSize-rightIdx-1 {
			m = BlockSize - rightIdx - 1
		}
		leftBase := leftIdx
		rightBase := rightIdx
		leftIdx += m
		rightIdx += m
		n += m
		for i := 0; m > 0; i++ {
			right.data[rightBase+i+1] = left.data[leftBase+i]
			m--
		}

		if leftIdx == BlockSize {
			b = left
			left = left.right
			leftIdx = 0
		}
	}

	d.left = left
	d.right = right
	d.leftIdx = leftIdx
	d.rightIdx = rightIdx
}

// Entries returns the elements contained within the deque as a slice.
// This method doesn't modify the deque.
func (d *Deque[T]) Entries() []T {
	d.initialize()
	if d.len == 0 {
		return nil
	}

	entries := make([]T, 0, d.len)

	b := d.left
	i := d.leftIdx
	for {
		entries = append(entries, b.data[i])
		i++
		if b == d.right && i > d.rightIdx {
			break
		}

		if i == BlockSize {
			i = 0
			b = b.right
		}
	}

	return entries
}

func (d *Deque[T]) initialize() {
	if d.left == nil {
		block := &Block[T]{}
		d.left = block
		d.right = block
		d.leftIdx = Center + 1
		d.rightIdx = Center
	}
}
