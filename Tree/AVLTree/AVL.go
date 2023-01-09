package AVLTree

import (
	"sync"

	"github.com/oyal2/Go-Structures/Tree/AVLTree/Node"
	"github.com/oyal2/Go-Structures/utils"
)

type AVLTree[T utils.Ordered] struct {
	root         *Node.Node[T]
	numberStored int64
	sync.RWMutex
}

func New[T utils.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{
		root:         nil,
		numberStored: 0,
	}
}

func (B *AVLTree[T]) Insert(element T) {
	B.Lock()
	defer B.Unlock()
	B.numberStored++
	if B.root == nil {
		B.root = Node.New(element)
	} else {
		B.root = B.root.Add(element)
	}
}

func (B *AVLTree[T]) Remove(element T) (valueReturn T) {
	B.Lock()
	defer B.Unlock()
	if B.root == nil {
		return valueReturn
	} else {
		B.root = B.root.Remove(element)
	}
	B.numberStored--
	return element
}

func (B *AVLTree[T]) Traverse() (arr []T) {
	B.RLock()
	defer B.RUnlock()
	if B.root != nil {
		return B.root.Inorder()
	}

	return arr
}

func (B *AVLTree[T]) Contains(element T) bool {
	B.RLock()
	defer B.RUnlock()
	if B.numberStored == 0 {
		return false
	}

	currNode := B.root

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
	B.RLock()
	defer B.RUnlock()
	return B.numberStored
}
