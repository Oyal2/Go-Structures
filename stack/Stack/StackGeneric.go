package Stack

import "github.com/oyal2/Go-Structures/list"

type Stack[T comparable] struct {
	_list list.List[T]
}

func New[T comparable](list list.List[T]) Stack[T] {
	return Stack[T]{
		_list: list,
	}
}

func (S *Stack[T]) Push(element T) error {
	err := S._list.Insert(S._list.Length(), element)
	return err
}

func (S *Stack[T]) Peek() (T, error) {
	val, err := S._list.Get(S._list.Length() - 1)
	return val, err
}

func (S *Stack[T]) Pop() (T, error) {
	val, err := S._list.Remove(S._list.Length() - 1)
	return val, err
}

func (S *Stack[T]) IsEmpty() bool {
	return S._list.Length() == 0
}

func (S *Stack[T]) Length() int {
	return S._list.Length()
}

func (S *Stack[T]) Get(index int) (T, error) {
	val, err := S._list.Get(index)
	return val, err
}
