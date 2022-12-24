package PriorityQueue

import (
	"github.com/oyal2/Go-Structures/heap"
	"github.com/oyal2/Go-Structures/utils"
)

type PriorityQueue[T utils.Ordered] struct {
	_storage heap.Heap[T]
}

func New[T utils.Ordered](storage heap.Heap[T]) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		_storage: storage,
	}
}

func (Q *PriorityQueue[T]) Enqueue(elements ...T) error {
	err := Q._storage.Insert(elements...)
	return err
}

func (Q *PriorityQueue[T]) Front() (T, error) {
	return Q._storage.Peek()
}

func (Q *PriorityQueue[T]) Dequeue() (T, error) {
	return Q._storage.Extract()
}

func (Q *PriorityQueue[T]) Contains(element T) bool {
	return Q._storage.Contains(element)
}

func (Q *PriorityQueue[T]) IsEmpty() bool {
	return Q._storage.IsEmpty()
}

func (Q *PriorityQueue[T]) Length() int {
	return Q._storage.Length()
}

func (Q *PriorityQueue[T]) Clear() {
	Q._storage.Clear()
}

func (Q *PriorityQueue[T]) Update(element, newElement T) error {
	return Q._storage.ChangeKey(element, newElement)
}
