package Node

import (
	"github.com/oyal2/Go-Structures/utils"
)

type Node[T utils.Ordered] struct {
	Value  T
	Height int64
	Parent *Node[T]
	Right  *Node[T]
	Left   *Node[T]
}

func New[T utils.Ordered](value T) *Node[T] {
	return &Node[T]{
		Value:  value,
		Height: 0,
	}
}

func (Node *Node[T]) ComputeHeight() {
	var height int64 = -1
	if Node.Left != nil {
		height = utils.Max(height, Node.Left.Height)
	}

	if Node.Right != nil {
		height = utils.Max(height, Node.Right.Height)
	}

	Node.Height = height + 1
}

func (Node *Node[T]) HeightDifference() int64 {
	var left int64 = 0
	var right int64 = 0

	if Node.Left != nil {
		left = 1 + Node.Left.Height
	}

	if Node.Right != nil {
		right = 1 + Node.Right.Height
	}

	return left - right
}

func (Node *Node[T]) RotateRight() *Node[T] {
	newRoot := Node.Left
	grandson := newRoot.Right
	Node.Left = grandson
	newRoot.Right = Node

	Node.ComputeHeight()
	return newRoot
}

func (Node *Node[T]) RotateRightLeft() *Node[T] {
	child := Node.Right
	newRoot := child.Left
	grand1 := newRoot.Left
	grand2 := newRoot.Right
	child.Left = grand2
	Node.Right = grand1

	newRoot.Left = Node
	newRoot.Right = child

	child.ComputeHeight()
	Node.ComputeHeight()
	return newRoot
}

func (Node *Node[T]) RotateLeft() *Node[T] {
	newRoot := Node.Right
	grandson := newRoot.Left
	Node.Right = grandson
	newRoot.Left = Node

	Node.ComputeHeight()
	return newRoot
}

func (Node *Node[T]) rotateLeftRight() *Node[T] {
	child := Node.Left
	newRoot := child.Right
	grand1 := newRoot.Right
	grand2 := newRoot.Left
	child.Right = grand2
	Node.Left = grand1

	newRoot.Right = Node
	newRoot.Left = child

	child.ComputeHeight()
	Node.ComputeHeight()
	return newRoot
}

func (Node *Node[T]) Add(value T) *Node[T] {
	newRoot := Node
	if value <= Node.Value {
		Node.Left = Node.addToSubTree(Node.Left, value)
		if Node.HeightDifference() == 2 {
			if value <= Node.Left.Value {
				newRoot = Node.RotateRight()
			} else {
				newRoot = Node.rotateLeftRight()
			}
		}
	} else {
		Node.Right = Node.addToSubTree(Node.Right, value)
		if Node.HeightDifference() == -2 {
			if value > Node.Right.Value {
				newRoot = Node.RotateLeft()
			} else {
				newRoot = Node.RotateRightLeft()
			}
		}
	}
	newRoot.ComputeHeight()
	return newRoot
}

func (Node *Node[T]) addToSubTree(parent *Node[T], value T) *Node[T] {
	if parent == nil {
		return New(value)
	}

	parent = parent.Add(value)
	return parent

}

func (Node *Node[T]) Inorder() (arr []T) {
	if Node.Left != nil {
		arr = append(arr, Node.Left.Inorder()...)
	}

	arr = append(arr, Node.Value)

	if Node.Right != nil {
		arr = append(arr, Node.Right.Inorder()...)
	}

	return arr
}

func (Node *Node[T]) Remove(value T) *Node[T] {
	newRoot := Node
	if value == Node.Value {
		if Node.Left == nil {
			return Node.Right
		}

		child := Node.Left
		for child.Right != nil {
			child = child.Right
		}
		childKey := child.Value
		Node.Left = Node.RemoveFromParent(Node.Left, childKey)
		Node.Value = childKey
	} else if value < Node.Value {
		Node.Left = Node.RemoveFromParent(Node.Left, value)
		if Node.HeightDifference() == -2 {
			if Node.Left.HeightDifference() <= 0 {
				newRoot = Node.RotateLeft()
			} else {
				newRoot = Node.RotateRightLeft()
			}
		}
	} else {
		Node.Right = Node.RemoveFromParent(Node.Right, value)
		if Node.HeightDifference() == 2 {
			if Node.Left.HeightDifference() >= 0 {
				newRoot = Node.RotateRight()
			} else {
				newRoot = Node.rotateLeftRight()
			}
		}
	}
	newRoot.ComputeHeight()
	return newRoot
}

func (Node *Node[T]) RemoveFromParent(parent *Node[T], value T) *Node[T] {
	if parent != nil {
		return parent.Remove(value)
	}
	return nil
}
