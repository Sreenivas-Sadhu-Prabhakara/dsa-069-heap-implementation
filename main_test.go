package main

import (
	"reflect"
	"testing"
)

func drain(h *Heap) []int {
	out := []int{}
	for h.Len() > 0 {
		v, _ := h.Pop()
		out = append(out, v)
	}
	return out
}

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap()
	for _, v := range []int{3, 1, 4, 1, 5, 9, 2, 6} {
		h.Push(v)
	}
	if top, _ := h.Peek(); top != 9 {
		t.Errorf("Peek = %d, want 9", top)
	}
	want := []int{9, 6, 5, 4, 3, 2, 1, 1}
	if got := drain(h); !reflect.DeepEqual(got, want) {
		t.Errorf("max drain = %v, want %v", got, want)
	}
}

func TestMinHeap(t *testing.T) {
	h := NewMinHeap()
	for _, v := range []int{3, 1, 4, 1, 5, 9, 2, 6} {
		h.Push(v)
	}
	want := []int{1, 1, 2, 3, 4, 5, 6, 9}
	if got := drain(h); !reflect.DeepEqual(got, want) {
		t.Errorf("min drain = %v, want %v", got, want)
	}
}

func TestEmpty(t *testing.T) {
	h := NewMinHeap()
	if _, ok := h.Pop(); ok {
		t.Error("Pop on empty should be not-ok")
	}
	if _, ok := h.Peek(); ok {
		t.Error("Peek on empty should be not-ok")
	}
	h.Push(42)
	if v, ok := h.Pop(); !ok || v != 42 {
		t.Errorf("single push/pop got (%d,%v)", v, ok)
	}
}
