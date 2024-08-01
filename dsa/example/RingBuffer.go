package example

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyRingBuffer = errors.New("invalid attempt to pop, empty buffer")
	ErrInvalidIdx      = errors.New("invalid attempt to get, index out of range")
)

type RingBuffer[T any] struct {
	arr  []T
	head int
	tail int
	size int
}

func NewRingBuffer[T any](s int) *RingBuffer[T] {
	rb := &RingBuffer[T]{
		arr:  make([]T, s),
		head: 0,
		tail: -1,
		size: s,
	}

	return rb
}

// overwriting ring buffer

func (rb *RingBuffer[T]) Get(idx int) (T, error) {
	if idx >= rb.size {
		var empty T
		return empty, ErrInvalidIdx
	}

	actualIdx := (rb.head + idx) % rb.size
	return rb.arr[actualIdx], nil
}

func (rb *RingBuffer[T]) Push(item T) {
	rb.tail = (rb.tail + 1) % rb.size
	rb.arr[rb.tail] = item

	if rb.head == rb.tail {
		// fix the head position
		rb.head = (rb.head + 1) % rb.size
	}
}

func (rb *RingBuffer[T]) Pop() (T, error) {
	if rb.head == rb.tail || rb.tail == -1 {
		var empty T
		return empty, ErrEmptyRingBuffer
	}

	value := rb.arr[rb.head]
	rb.head = (rb.head + 1) % rb.size

	return value, nil
}

func (rb *RingBuffer[T]) Print() {
	fmt.Println("--------------------------")
	fmt.Println("head -> ", rb.head)
	fmt.Println("tail -> ", rb.tail)
	fmt.Println("--------------------------")
	for i, item := range rb.arr {
		fmt.Printf("index - %d * item - %+v\n", i, item)
	}
	fmt.Println("--------------------------")
}
