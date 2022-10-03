package queue

import "github.com/oyal2/Go-Structures/list"

type Stack struct {
	_list list.List
}

func NewStackGeneric(list list.List) Stack {
	return Stack{
		_list: list,
	}
}

func (S *Stack) Push(element interface{}) error {
	err := S._list.Insert(S._list.Length(), element)
	return err
}

func (S *Stack) Top() (interface{},error) {
	val,err := S._list.Get(S._list.Length() - 1)
	return val,err
}

func (S *Stack) Pop() (interface{},error) {
	val,err := S._list.Remove(S._list.Length() - 1)
	return val,err
}

func (S *Stack) isEmpty() bool {
	return S._list.Length() == 0
}

func (S *Stack) Length() int {
	return S._list.Length()
}

func (S *Stack) Get(index int) (interface{},error) {
	val,err := S._list.Get(index)
	return val,err
}