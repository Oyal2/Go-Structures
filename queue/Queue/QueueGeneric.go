package Queue

import "github.com/oyal2/Go-Structures/list"

type Queue[T comparable] struct {
	_list list.List[T]
}

func New[T comparable](list list.List[T]) *Queue[T] {
	return &Queue[T]{
		_list: list,
	}
}

func (Q *Queue[T]) Enqueue(element T) error {
	err := Q._list.Insert(Q._list.Length(), element)
	return err
}

func (Q *Queue[T]) Front() (T, error) {
	return Q._list.Get(0)
}

func (Q *Queue[T]) Dequeue() (T, error) {
	return Q._list.Remove(0)
}

func (Q *Queue[T]) IsEmpty() bool {
	return Q._list.Length() == 0
}

func (Q *Queue[T]) Length() int {
	return Q._list.Length()
}

func (Q *Queue[T]) Get(index int) (T, error) {
	return Q._list.Get(index)
}
