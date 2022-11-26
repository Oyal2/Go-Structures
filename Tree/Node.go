package Tree

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
		Value: value,
	}
}

func (selfNode *Node[T]) Balance() *Node[T] {
	diff := selfNode.Parent.HeightDifference()
	if diff <= -1 {
		if Height(selfNode.Right) > Height(selfNode.Left) {
			return selfNode.Parent.RotateLeft()
		} else {
			newSet := selfNode.RotateRight()
			return newSet.Parent.RotateLeft()
		}

	} else if diff >= 1 {
		if Height(selfNode.Left) > Height(selfNode.Right) {
			return selfNode.Parent.RotateRight()
		} else {
			newSet := selfNode.RotateLeft()
			return newSet.Parent.RotateRight()
		}
	}
	return selfNode
}

func (ParentNode *Node[T]) RotateRight() *Node[T] {
	var x, y, z, t2 *Node[T]
	z = ParentNode
	y = z.Left
	if y != nil {
		x = y.Left
		t2 = y.Right
	}

	z.Left = t2
	y.Right = z
	y.Parent = z.Parent
	if y.Parent != nil {
		if y.Parent.Left == z {
			y.Parent.Left = y
		} else {
			y.Parent.Right = y
		}
	}
	z.Parent = y

	z.Height++
	y.Height--
	if x != nil {
		x.Height--
		if x.Left != nil {
			x.Left.Height--
		}
	}
	return y
}

func (ParentNode *Node[T]) RotateLeft() *Node[T] {
	var x, y, z, t2 *Node[T]
	z = ParentNode
	y = z.Right
	if y != nil {
		x = y.Right
		t2 = y.Left
	}

	z.Right = t2
	y.Left = z
	y.Parent = z.Parent
	//Parent value 2 right is 1
	if y.Parent != nil {
		if y.Parent.Left == z {
			y.Parent.Left = y
		} else {
			y.Parent.Right = y
		}
	}
	z.Parent = y

	z.Height++
	y.Height--
	if x != nil {
		x.Height--
		if x.Right != nil {
			x.Right.Height--
		}
	}
	return y
}

func (Node *Node[T]) HeightDifference() int64 {
	if Node == nil {
		return 0
	}
	var left int64 = 0
	var right int64 = 0

	if Node.Left != nil {
		left = Node.Left.Height
	}

	if Node.Right != nil {
		right = Node.Right.Height
	}

	return left - right
}

func (Node *Node[T]) CalcNodeHeight() int64 {
	if Node.Parent != nil {
		return Node.Parent.Height + 1
	}

	return 0
}

func Height[T utils.Ordered](Node *Node[T]) int64 {
	if Node == nil {
		return -1
	}

	return Node.Height
}

func (Node *Node[T]) FindReplacement() *Node[T] {
	tempNode := Node.Right

	if tempNode.Left == nil {
		Node.Value = tempNode.Value
		Node.Right = tempNode.Right
	} else {
		for tempNode.Left != nil {
			tempNode = tempNode.Left
		}

		Node.Value = tempNode.Value
		tempNode.Parent.Left = nil
		tempNode.Parent = tempNode.Parent.Balance()
	}
	return tempNode
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
