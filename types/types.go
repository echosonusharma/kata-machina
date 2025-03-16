package types

type Node struct {
	Value string
	Next  *Node
}

type LinkedList struct {
	Length uint
	Head   *Node
}
