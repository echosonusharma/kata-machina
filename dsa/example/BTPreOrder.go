package example

type BinaryNode[T comparable] struct {
	Value T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}

func walk[T comparable](curr *BinaryNode[T], path *[]T) {
	if curr == nil {
		return
	}

	// pre
	*path = append(*path, curr.Value)
	// recurse
	walk(curr.Left, path)
	walk(curr.Right, path)
	// post
}

func (bt *BinaryNode[T]) PreOrderSearch() []T {
	path := []T{}
	walk(bt, &path)

	return path
}
