package example

import "fmt"

type QNode[T any] struct {
	value T
	next  *QNode[T]
}

type Queue[T any] struct {
	length uint
	head   *QNode[T]
	tail   *QNode[T]
}

// add item to the queue
func (q *Queue[T]) Enqueue(item T) {
	node := &QNode[T]{value: item, next: nil}

	if q.length == 0 {
		q.head = node
		q.tail = node
		q.length++

		return
	}

	q.tail.next = node
	q.tail = node
	q.length++
}

// get the first item from the queue
func (q *Queue[T]) Dequeue() (T, error) {
	if q.length == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("can't dequeue an empty queue")
	}

	q.length -= 1
	head := q.head

	q.head = q.head.next

	if q.length == 0 {
		q.tail = nil
	}

	return head.value, nil
}

// get the first item from the queue without dequeueing
func (q *Queue[T]) Peek() T {
	return q.head.value
}

// size of the queue
func (q *Queue[T]) Size() uint {
	return q.length
}
