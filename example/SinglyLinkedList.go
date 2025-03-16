package example

import (
	"errors"
	"fmt"

	t "github.com/echosonusharma/kata-machina/types"
)

func LlGet(ll *t.LinkedList, index uint) (string, error) {
	if index >= ll.Length {
		return "", fmt.Errorf("index %d - out of range", index)
	}

	var currNode *t.Node = ll.Head

	for i := 0; currNode != nil && i < int(index); i++ {
		currNode = currNode.Next
	}

	return currNode.Value, nil
}

func LlAppend(ll *t.LinkedList, item string) {
	node := &t.Node{Value: item, Next: nil}

	if ll.Length == 0 {
		ll.Head = node
	} else {
		curNode := ll.Head

		for curNode.Next != nil {
			curNode = curNode.Next
		}

		curNode.Next = node
	}

	ll.Length++
}

func LlPrepend(ll *t.LinkedList, item string) {
	node := &t.Node{Value: item, Next: nil}

	if ll.Length == 0 {
		ll.Head = node
	} else {
		node.Next = ll.Head
		ll.Head = node
	}

	ll.Length++
}

func LlRemove(ll *t.LinkedList, item string) (string, error) {
	if ll.Length == 0 {
		return "", errors.New("item not found")
	}

	currNode := ll.Head

	if currNode.Value == item {
		ll.Head = ll.Head.Next
		return ll.Head.Value, nil
	}

	for currNode.Next != nil {
		if currNode.Next.Value == item {
			break
		}

		currNode = currNode.Next
	}

	if currNode == nil {
		return "", errors.New("item not found")
	}

	currNode.Next = currNode.Next.Next
	ll.Length--

	return currNode.Value, nil
}

func LlInsertAt(ll *t.LinkedList, item string, index uint) error {
	if index > ll.Length {
		return fmt.Errorf("index %d - out of range", index)
	}

	if index == 0 {
		LlPrepend(ll, item)
		return nil
	} else if index == ll.Length {
		LlAppend(ll, item)
		return nil
	}

	node := &t.Node{Value: item}
	currNode := ll.Head

	var i uint
	for i = 0; i < (index-1) && currNode.Next != nil; i++ {
		currNode = currNode.Next
	}

	if (i + 1) != index {
		return fmt.Errorf("failed to insert item at index %d", index)
	}

	node.Next = currNode.Next
	currNode.Next = node
	ll.Length++

	return nil
}

func LlRemoveAt(ll *t.LinkedList, index uint) error {
	if index >= ll.Length {
		return fmt.Errorf("index %d - out of range", index)
	}

	if index == 0 {
		ll.Head = ll.Head.Next
		ll.Length--
		return nil
	}

	currNode := ll.Head

	var i uint
	for i = 0; i < (index-1) && currNode.Next != nil; i++ {
		currNode = currNode.Next
	}

	if (i + 1) != index {
		return fmt.Errorf("failed to remove item at index %d", index)
	}

	currNode.Next = currNode.Next.Next
	ll.Length--

	return nil
}

func LlPrintList(ll *t.LinkedList) {
	node := ll.Head
	for node != nil {
		fmt.Printf("%v -> ", node.Value)
		node = node.Next
	}
	fmt.Println("")
}
