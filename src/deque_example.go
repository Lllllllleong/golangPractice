package main

import (
	"container/list"
	"fmt"
)

func dequeExample() {
	// Create a new double-ended queue
	deque := list.New()

	// Add elements to the front (like a stack)
	deque.PushFront("first")
	deque.PushFront("second")

	// Add elements to the back (like a queue)
	deque.PushBack("third")
	deque.PushBack("fourth")

	fmt.Println("Deque contents (front to back):")
	// Iterate from front to back
	for e := deque.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()

	// Remove from front
	if front := deque.Front(); front != nil {
		fmt.Printf("Removed from front: %v\n", front.Value)
		deque.Remove(front)
	}

	// Remove from back
	if back := deque.Back(); back != nil {
		fmt.Printf("Removed from back: %v\n", back.Value)
		deque.Remove(back)
	}

	fmt.Println("After removals:")
	for e := deque.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()

	// Check size
	fmt.Printf("Deque length: %d\n", deque.Len())
}

// Alternative: Simple slice-based deque implementation
type SliceDeque[T any] struct {
	items []T
}

func NewSliceDeque[T any]() *SliceDeque[T] {
	return &SliceDeque[T]{items: make([]T, 0)}
}

func (d *SliceDeque[T]) PushFront(item T) {
	d.items = append([]T{item}, d.items...)
}

func (d *SliceDeque[T]) PushBack(item T) {
	d.items = append(d.items, item)
}

func (d *SliceDeque[T]) PopFront() (T, bool) {
	var zero T
	if len(d.items) == 0 {
		return zero, false
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item, true
}

func (d *SliceDeque[T]) PopBack() (T, bool) {
	var zero T
	if len(d.items) == 0 {
		return zero, false
	}
	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item, true
}

func (d *SliceDeque[T]) Len() int {
	return len(d.items)
}

func (d *SliceDeque[T]) IsEmpty() bool {
	return len(d.items) == 0
}

func sliceDequeExample() {
	fmt.Println("\n=== Slice-based Deque Example ===")

	deque := NewSliceDeque[string]()

	// Add to both ends
	deque.PushFront("front1")
	deque.PushBack("back1")
	deque.PushFront("front2")
	deque.PushBack("back2")

	fmt.Printf("Deque length: %d\n", deque.Len())

	// Remove from both ends
	if front, ok := deque.PopFront(); ok {
		fmt.Printf("Popped from front: %v\n", front)
	}

	if back, ok := deque.PopBack(); ok {
		fmt.Printf("Popped from back: %v\n", back)
	}

	fmt.Printf("Remaining length: %d\n", deque.Len())
}
