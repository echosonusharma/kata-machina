package example

import (
	"errors"
)

type DNode[T comparable] struct {
	value T
	next  *DNode[T]
	prev  *DNode[T]
}

type DoublyLinkedList[T comparable] struct {
	length uint
	head   *DNode[T]
	tail   *DNode[T]
}

func (dll *DoublyLinkedList[T]) Get(index uint) (T, error) {
	if index >= dll.length {
		var empty T
		return empty, errors.New("index out of range")
	}

	var curr *DNode[T] = dll.getNodeAt(index)
	return curr.value, nil
}

func (dll *DoublyLinkedList[T]) Append(item T) {
	node := &DNode[T]{
		value: item,
		next:  nil,
		prev:  nil,
	}

	if dll.length == 0 {
		dll.length++
		dll.head, dll.tail = node, node
		return
	}

	node.prev = dll.tail
	dll.tail.next = node
	dll.tail = node
	dll.length++
}

func (dll *DoublyLinkedList[T]) Prepend(item T) {
	node := &DNode[T]{
		value: item,
		next:  nil,
		prev:  nil,
	}

	dll.length++
	if dll.length == 0 {
		dll.head = node
		dll.tail = node
		return
	}

	node.next = dll.head
	dll.head.prev = node
	dll.head = node
}

func (dll *DoublyLinkedList[T]) Remove(item T) (T, error) {
	var curr *DNode[T] = dll.head

	for i := 0; curr != nil && i < int(dll.length); i++ {
		if curr.value == item {
			break
		}

		curr = curr.next
	}

	if curr == nil {
		var emp T
		return emp, errors.New("item not found")
	}

	return dll.removeNode(curr), nil
}

func (dll *DoublyLinkedList[T]) InsertAt(item T, index uint) error {
	if index > dll.length {
		return errors.New("index out of range")
	}

	dll.length++
	if index == dll.length {
		dll.Append(item)
		return nil
	} else if index == 0 {
		dll.Prepend(item)
		return nil
	}

	node := &DNode[T]{
		value: item,
		next:  nil,
		prev:  nil,
	}

	var curr *DNode[T] = dll.getNodeAt(index)

	node.next = curr
	node.prev = curr.prev

	curr.prev.next = node
	curr.prev = node

	return nil
}

func (dll *DoublyLinkedList[T]) RemoveAt(index uint) {
	var curr *DNode[T] = dll.head

	for i := 0; curr != nil && i < int(index); i++ {
		curr = curr.next
	}

	if curr == nil {
		return
	}

	dll.removeNode(curr)
}

func (dll *DoublyLinkedList[T]) getNodeAt(index uint) *DNode[T] {
	var node *DNode[T] = dll.head

	for i := 0; node != nil && i < int(index); i++ {
		node = node.next
	}

	return node
}

func (dll *DoublyLinkedList[T]) removeNode(node *DNode[T]) T {
	dll.length--

	if dll.length == 1 {
		val := dll.head.value
		dll.head, dll.tail = nil, nil
		return val
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == dll.head {
		dll.head = node.next
	}

	if node == dll.tail {
		dll.tail = node.prev
	}

	node.next, node.prev = nil, nil

	return node.value
}

// func (ll *DoublyLinkedList[T]) debug() {
// 	node := ll.head
// 	for node != nil {
// 		fmt.Printf("%v <-> ", node.value)
// 		node = node.next
// 	}
// 	fmt.Println("")
// }
