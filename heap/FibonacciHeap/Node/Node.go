package Node

import "github.com/oyal2/Go-Structures/utils"

type Node[T utils.Ordered] struct {
	Degree  int
	Parent  *Node[T]
	Right   *Node[T]
	Left    *Node[T]
	Child   *Node[T]
	Mark    bool
	Element T
}

func New[T utils.Ordered]() *Node[T] {
	return &Node[T]{}
}
