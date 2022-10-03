package queue

import "github.com/oyal2/Go-Structures/list"

type Queue struct {
	_list list.List
}

func NewQueueGeneric(list list.List) Queue {
	return Queue{
		_list: list,
	}
}

func (Q *Queue) Enqueue(element interface{}) error {
	err := Q._list.Insert(Q._list.Length(), element)
	return err
}

func (Q *Queue) Front() (interface{},error) {
	val,err := Q._list.Get(0)
	return val,err
}

func (Q *Queue) Dequeue() (interface{},error) {
	val,err := Q._list.Remove(0)
	return val,err
}

func (Q *Queue) isEmpty() bool {
	return Q._list.Length() == 0
}

func (Q *Queue) Length() int {
	return Q._list.Length()
}

func (Q *Queue) Get(index int) (interface{},error) {
	val,err := Q._list.Get(index)
	return val,err
}