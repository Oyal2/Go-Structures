package Node

type Node[T any] struct {
	Children []*Node[T]
	Keys     []T
	Leaf     bool
	N        int
}

func New[T any](sizeNode int) *Node[T] {
	return &Node[T]{
		Children: make([]*Node[T], sizeNode),
		Keys:     make([]T, sizeNode-1),
		Leaf:     true,
		N:        0,
	}
}
