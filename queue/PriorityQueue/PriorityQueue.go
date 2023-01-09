package PriorityQueue

import (
	"github.com/oyal2/Go-Structures/heap"
	"github.com/oyal2/Go-Structures/utils"
)

type PriorityQueue[T utils.Ordered] struct {
	storage heap.Heap[T]
}

func New[T utils.Ordered](storage heap.Heap[T]) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		storage: storage,
	}
}

func (Q *PriorityQueue[T]) Enqueue(elements ...T) error {
	err := Q.storage.Insert(elements...)
	return err
}

func (Q *PriorityQueue[T]) Front() (T, error) {
	return Q.storage.Peek()
}

func (Q *PriorityQueue[T]) Dequeue() (T, error) {
	return Q.storage.Extract()
}

func (Q *PriorityQueue[T]) Contains(element T) bool {
	return Q.storage.Contains(element)
}

func (Q *PriorityQueue[T]) IsEmpty() bool {
	return Q.storage.IsEmpty()
}

func (Q *PriorityQueue[T]) Length() int {
	return Q.storage.Length()
}

func (Q *PriorityQueue[T]) Clear() {
	Q.storage.Clear()
}

func (Q *PriorityQueue[T]) Update(element, newElement T) error {
	return Q.storage.ChangeKey(element, newElement)
}
