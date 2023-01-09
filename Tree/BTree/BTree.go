package BTree

import (
	"fmt"

	"github.com/oyal2/Go-Structures/Tree/BTree/Node"
)

type BTree[T any] struct {
	_root         *Node.Node[T]
	_numberStored int64
	M             int
	_comparator   func(a, b T) bool
}

func New[T any](comparator func(a, b T) bool, max int) *BTree[T] {
	if max < 3 {
		return nil
	}
	return &BTree[T]{
		_root:         Node.New[T](max),
		_numberStored: 0,
		M:             max,
		_comparator:   comparator,
	}
}

func (B *BTree[T]) Insert(element T) {
	B._numberStored++
	r := B._root
	if r.N == B.M-1 {
		s := Node.New[T](B.M)
		B._root = s
		s.Leaf = false
		s.N = 0
		s.Children[0] = r
		B.SplitChild(s, 0)
		B.NonFullInsert(s, element)
	} else {
		B.NonFullInsert(r, element)
	}

}

func (B *BTree[T]) SplitChild(node *Node.Node[T], position int) {
	var empty T
	y := node.Children[position]
	z := Node.New[T](B.M)
	z.Leaf = y.Leaf
	for i := (B.M / 2) + 1; i < (B.M)-1; i++ {
		z.Keys[i] = y.Keys[i]
	}

	if !y.Leaf {
		for i := B.M / 2; i < B.M; i++ {
			z.Children[i] = y.Children[i]
			y.Children[i] = nil
		}
	}

	for i := node.N; i >= position+1; i-- {
		node.Children[i+1] = node.Children[i]
	}

	node.Children[position+1] = z

	for i := node.N - 1; i >= position; i-- {
		node.Keys[i+1] = node.Keys[i]
	}
	node.Keys[position] = y.Keys[B.M/2]
	y.Keys[B.M/2] = empty
	y.N--

	node.N = node.N + 1
}

func (B *BTree[T]) NonFullInsert(node *Node.Node[T], element T) {
	i := node.N - 1
	if node.Leaf {

		for i >= 0 && B._comparator(element, node.Keys[i]) {
			fmt.Println(node.Keys[i], element)
			node.Keys[i+1] = node.Keys[i]
			i--
		}
		node.Keys[i+1] = element
		node.N = node.N + 1
	} else {
		for i >= 0 && B._comparator(element, node.Keys[i]) {
			i--
		}
		i++
		tempNode := node.Children[i]
		if tempNode.N == len(tempNode.Keys) {
			B.SplitChild(node, i)
			if B._comparator(node.Keys[i], element) {
				i++
			}
		}
		B.NonFullInsert(node.Children[i], element)
	}
}
