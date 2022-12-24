package heap

import "github.com/oyal2/Go-Structures/utils"

type Heap[T utils.Ordered] interface {
	utils.Sort[T]
	Insert(element T) error
	Extract() (T, error)
	Peek() (T, error)
	ChangeKey(element T, newElement T) error
	Contains(element T) bool
	Length() int
	IsEmpty() bool
	Clear()
}
