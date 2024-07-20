package example

type SNode[T any] struct {
	value T
	prev  *SNode[T]
}

type Stack[T any] struct {
	length uint
	head   *SNode[T]
}

func (s *Stack[T]) Push(item T) {
	node := &SNode[T]{value: item, prev: nil}

	if s.length == 0 {
		s.head = node
		s.length++
		return
	}

	s.length++
	node.prev = s.head
	s.head = node
}

func (s *Stack[T]) Pop() T {
	if s.length == 0 {
		return s.head.value
	}

	head := s.head
	s.head = s.head.prev
	s.length--
	return head.value
}

func (s *Stack[T]) Peek() T {
	return s.head.value
}

func (s *Stack[T]) Size() uint {
	return s.length
}
