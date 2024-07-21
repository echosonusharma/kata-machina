package example

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	length uint
	head   *Node[T]
}

func (ll *LinkedList[T]) Get(index uint) (T, error) {
	if index >= ll.length {
		return ll.head.value, fmt.Errorf("invalid index - %d", index)
	}

	currNode := ll.head

	var i uint = 0
	for i != index {
		currNode = currNode.next
		i++
	}

	return currNode.value, nil
}

func (ll *LinkedList[T]) Append(item T) {
	node := &Node[T]{value: item, next: nil}

	if ll.length == 0 {
		ll.head = node
	} else {
		curNode := ll.head

		for curNode.next != nil {
			curNode = curNode.next
		}

		curNode.next = node
	}

	ll.length++
}

func (ll *LinkedList[T]) Prepend(item T) {
	node := &Node[T]{value: item, next: nil}

	if ll.length == 0 {
		ll.head = node
	} else {
		node.next = ll.head
		ll.head = node
	}

	ll.length++
}

func (ll *LinkedList[T]) Remove() error {
	if ll.length == 0 {
		return errors.New("invalid attempt to remove")
	}

	if ll.head.next == nil {
		ll.head = nil
		ll.length--
		return nil
	}

	curNode := ll.head
	for curNode.next != nil && curNode.next.next != nil {
		curNode = curNode.next
	}

	curNode.next = nil
	ll.length--

	return nil
}

func (ll *LinkedList[T]) InsertAt(item T, index uint) error {
	if index > ll.length {
		return nil
	}
	node := &Node[T]{value: item}

	if index == 0 {
		node.next = ll.head
		ll.head = node
		ll.length++
		return nil
	}

	currNode := ll.head

	var i uint
	for i = 0; i < (index-1) && currNode.next != nil; i++ {
		currNode = currNode.next
	}

	if (i + 1) != index {
		return fmt.Errorf("failed to insert item at index %d", index)
	}

	node.next = currNode.next
	currNode.next = node
	ll.length++

	return nil
}

func (ll *LinkedList[T]) RemoveAt(index uint) error {
	if index >= ll.length {
		return nil
	}

	if index == 0 {
		ll.head = ll.head.next
		ll.length--
		return nil
	}

	currNode := ll.head

	var i uint
	for i = 0; i < (index-1) && currNode.next != nil; i++ {
		currNode = currNode.next
	}

	if (i + 1) != index {
		return fmt.Errorf("failed to remove item at index %d", index)
	}

	currNode = currNode.next.next
	ll.length--

	return nil
}
