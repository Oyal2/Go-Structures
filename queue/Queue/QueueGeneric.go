package Queue

import "github.com/oyal2/Go-Structures/list"

type Queue[T comparable] struct {
	list list.List[T]
}

func New[T comparable](list list.List[T]) *Queue[T] {
	return &Queue[T]{
		list: list,
	}
}

func (Q *Queue[T]) Enqueue(element T) error {
	err := Q.list.Insert(Q.list.Length(), element)
	return err
}

func (Q *Queue[T]) Front() (T, error) {
	return Q.list.Get(0)
}

func (Q *Queue[T]) Dequeue() (T, error) {
	return Q.list.Remove(0)
}

func (Q *Queue[T]) IsEmpty() bool {
	return Q.list.Length() == 0
}

func (Q *Queue[T]) Length() int {
	return Q.list.Length()
}

func (Q *Queue[T]) Get(index int) (T, error) {
	return Q.list.Get(index)
}

func (Q *Queue[T]) Remove(element T) (T, error) {
	return Q._list.Remove(Q._list.IndexOf(element))
}
