package main

import "fmt"

// Heap is a binary heap over ints whose ordering is defined by before:
// before(a, b) true means a belongs above b. NewMaxHeap and NewMinHeap
// supply the two usual orderings.
type Heap struct {
	data   []int
	before func(a, b int) bool
}

// NewMaxHeap returns a heap whose Pop yields the largest value first.
func NewMaxHeap() *Heap { return &Heap{before: func(a, b int) bool { return a > b }} }

// NewMinHeap returns a heap whose Pop yields the smallest value first.
func NewMinHeap() *Heap { return &Heap{before: func(a, b int) bool { return a < b }} }

// Len reports how many values the heap holds.
func (h *Heap) Len() int { return len(h.data) }

// Peek returns the top value without removing it.
func (h *Heap) Peek() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
}

// Push adds v, sifting it up until its parent outranks it.
func (h *Heap) Push(v int) {
	h.data = append(h.data, v)
	i := len(h.data) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if !h.before(h.data[i], h.data[parent]) {
			break
		}
		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent
	}
}

// Pop removes and returns the top value, sifting the displaced last element
// down to its place.
func (h *Heap) Pop() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	top := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	i := 0
	for {
		l, r := 2*i+1, 2*i+2
		best := i
		if l < len(h.data) && h.before(h.data[l], h.data[best]) {
			best = l
		}
		if r < len(h.data) && h.before(h.data[r], h.data[best]) {
			best = r
		}
		if best == i {
			break
		}
		h.data[i], h.data[best] = h.data[best], h.data[i]
		i = best
	}
	return top, true
}

func main() {
	h := NewMaxHeap()
	for _, v := range []int{3, 1, 4, 1, 5, 9, 2, 6} {
		h.Push(v)
	}
	for h.Len() > 0 {
		v, _ := h.Pop()
		fmt.Print(v, " ")
	}
	fmt.Println()
}
