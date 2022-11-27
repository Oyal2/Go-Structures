package Node

import (
	"github.com/oyal2/Go-Structures/utils"
)

type Node[T utils.Ordered] struct {
	Children []*Node[T]
	Keys     []T
	Leaf     bool
	N        int
}

func New[T utils.Ordered](sizeNode int) *Node[T] {
	return &Node[T]{
		Children: make([]*Node[T], sizeNode),
		Keys:     make([]T, sizeNode-1),
		Leaf:     true,
		N:        0,
	}
}

func (Node *Node[T]) SplitChild(position int) {
	var empty T
	t := len(Node.Children)
	y := Node.Children[position]
	z := New[T](t)
	z.Leaf = y.Leaf
	for i := (t / 2) + 1; i < t-1; i++ {
		z.Keys[i] = y.Keys[i]
	}

	if !y.Leaf {
		for i := (t / 2); i < t; i++ {
			z.Children[i] = y.Children[i]
			y.Children[i] = nil
		}
	}

	for i := Node.N; i >= position+1; i-- {
		Node.Children[i+1] = Node.Children[i]
	}

	Node.Children[position+1] = z

	for i := Node.N - 1; i >= position; i-- {
		Node.Keys[i+1] = Node.Keys[i]
	}
	Node.Keys[position] = y.Keys[t/2]
	y.Keys[t/2] = empty
	y.N--

	Node.N = Node.N + 1
}

func (Node *Node[T]) NonFullInsert(element T) {
	i := Node.N - 1
	if Node.Leaf {
		for i >= 0 && element < Node.Keys[i] {
			Node.Keys[i+1] = Node.Keys[i]
			i--
		}
		Node.Keys[i+1] = element
		Node.N = Node.N + 1
	} else {
		for i >= 0 && element < Node.Keys[i] {
			i--
		}
		i++
		tempNode := Node.Children[i]
		if tempNode.N == len(tempNode.Keys) {
			Node.SplitChild(i)
			if element > Node.Keys[i] {
				i++
			}
		}
		Node.Children[i].NonFullInsert(element)
	}
}
