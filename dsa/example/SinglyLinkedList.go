package example

import (
	"errors"
	"fmt"
)

type Node[S comparable] struct {
	value S
	next  *Node[S]
}

type LinkedList[S comparable] struct {
	length uint
	head   *Node[S]
}

func (ll *LinkedList[S]) Get(index uint) (S, error) {
	if index >= ll.length {
		var empty S
		return empty, fmt.Errorf("index %d - out of range", index)
	}

	var currNode *Node[S] = ll.head

	for i := 0; currNode != nil && i < int(index); i++ {
		currNode = currNode.next
	}

	return currNode.value, nil
}

func (ll *LinkedList[S]) Append(item S) {
	node := &Node[S]{value: item, next: nil}

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

func (ll *LinkedList[S]) Prepend(item S) {
	node := &Node[S]{value: item, next: nil}

	if ll.length == 0 {
		ll.head = node
	} else {
		node.next = ll.head
		ll.head = node
	}

	ll.length++
}

func (ll *LinkedList[S]) Remove(item S) (S, error) {
	if ll.length == 0 {
		var emp S
		return emp, errors.New("item not found")
	}

	currNode := ll.head

	if currNode.value == item {
		ll.head = ll.head.next
		return ll.head.value, nil
	}

	for currNode.next != nil {
		if currNode.next.value == item {
			break
		}

		currNode = currNode.next
	}

	if currNode == nil {
		var emp S
		return emp, errors.New("item not found")
	}

	currNode.next = currNode.next.next
	ll.length--

	return currNode.value, nil
}

func (ll *LinkedList[S]) InsertAt(item S, index uint) error {
	if index > ll.length {
		return fmt.Errorf("index %d - out of range", index)
	}

	if index == 0 {
		ll.Prepend(item)
		return nil
	} else if index == ll.length {
		ll.Append(item)
		return nil
	}

	node := &Node[S]{value: item}
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

func (ll *LinkedList[S]) RemoveAt(index uint) error {

	if index >= ll.length {
		return fmt.Errorf("index %d - out of range", index)
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

	currNode.next = currNode.next.next
	ll.length--

	return nil
}

// func (ll *LinkedList[S]) debug() {
// 	node := ll.head
// 	for node != nil {
// 		fmt.Printf("%v -> ", node.value)
// 		node = node.next
// 	}
// 	fmt.Println("")
// }
