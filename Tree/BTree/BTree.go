package BTree

import (
	"github.com/oyal2/Go-Structures/Tree/BTree/Node"
	"github.com/oyal2/Go-Structures/utils"
)

type BTree[T utils.Ordered] struct {
	_root         *Node.Node[T]
	_numberStored int64
	M             int
}

func New[T utils.Ordered](max int) *BTree[T] {
	if max < 3 {
		return nil
	}
	return &BTree[T]{
		_root:         Node.New[T](max),
		_numberStored: 0,
		M:             max,
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
		s.SplitChild(0)
		s.NonFullInsert(element)
	} else {
		r.NonFullInsert(element)
	}

}
