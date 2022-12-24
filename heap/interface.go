package heap

import "github.com/oyal2/Go-Structures/utils"

type Heap[T utils.Ordered] interface {
	utils.Sort[T]
	Insert() T
	Extract(element T)
	Peek() (T, error)
	ChangeKey(element T, newElement T) error
	Contains(element T) bool
	Length() int
	IsEmpty() bool
	Clear()
}
