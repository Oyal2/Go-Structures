package Stack

import "github.com/oyal2/Go-Structures/list"

type Stack[T comparable] struct {
	list list.List[T]
}

func New[T comparable](list list.List[T]) Stack[T] {
	return Stack[T]{
		list: list,
	}
}

func (S *Stack[T]) Push(element T) error {
	err := S.list.Insert(S.list.Length(), element)
	return err
}

func (S *Stack[T]) Peek() (T, error) {
	val, err := S.list.Get(S.list.Length() - 1)
	return val, err
}

func (S *Stack[T]) Pop() (T, error) {
	val, err := S.list.Remove(S.list.Length() - 1)
	return val, err
}

func (S *Stack[T]) IsEmpty() bool {
	return S.list.Length() == 0
}

func (S *Stack[T]) Length() int {
	return S.list.Length()
}

func (S *Stack[T]) Get(index int) (T, error) {
	val, err := S.list.Get(index)
	return val, err
}
