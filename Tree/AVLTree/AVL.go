package AVLTree

import (
	"github.com/oyal2/Go-Structures/Tree/AVLTree/Node"
	"github.com/oyal2/Go-Structures/utils"
)

type AVLTree[T utils.Ordered] struct {
	_root         *Node.Node[T]
	_numberStored int64
}

func New[T utils.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{
		_root:         nil,
		_numberStored: 0,
	}
}

func (B *AVLTree[T]) Insert(element T) {
	B._numberStored++
	if B._root == nil {
		B._root = Node.New(element)
	} else {
		B._root = B._root.Add(element)
	}
}

func (B *AVLTree[T]) Remove(element T) (valueReturn T) {
	if B._root == nil {
		return valueReturn
	} else {
		B._root = B._root.Remove(element)
	}
	B._numberStored--
	return element
}

func (B *AVLTree[T]) Traverse() (arr []T) {
	if B._root != nil {
		return B._root.Inorder()
	}

	return arr
}

func (B *AVLTree[T]) Contains(element T) bool {
	if B._numberStored == 0 {
		return false
	}

	currNode := B._root

	for currNode != nil {
		if element == currNode.Value {
			return true
		} else if element < currNode.Value {
			currNode = currNode.Left
		} else {
			currNode = currNode.Right
		}
	}

	return false
}

func (B *AVLTree[T]) Length() int64 {
	return B._numberStored
}
