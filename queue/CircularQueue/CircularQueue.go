package CircularQueue

import (
	"errors"

	"github.com/oyal2/Go-Structures/list"
)

type CircularQueue[T comparable] struct {
	_storage list.List[T]
	_max     int
	_front   int
	_rear    int
	_size    int
}

func New[T comparable](storage list.List[T], max int) *CircularQueue[T] {
	cq := &CircularQueue[T]{
		_storage: storage,
		_max:     max,
		_front:   -1,
		_rear:    -1,
	}

	var empty T
	for i := 0; i < max; i++ {
		cq._storage.Insert(i, empty)
	}

	return cq
}

func (CQ *CircularQueue[T]) Enqueue(elements ...T) error {
	if len(elements)+CQ.Length() > CQ._max {
		return errors.New("queue is full")
	}

	for _, element := range elements {

		if CQ._front == -1 {
			CQ._front = 0
		}
		CQ._rear = (CQ._rear + 1) % CQ._max

		_, err := CQ._storage.Update(CQ._rear, element)
		if err != nil {
			return err
		}

		CQ._size++
	}

	return nil
}

func (CQ *CircularQueue[T]) Front() (T, error) {
	var empty T
	if CQ.Length() == 0 {
		return empty, errors.New("queue is empty")
	}

	return CQ._storage.Get(CQ._front)
}

func (CQ *CircularQueue[T]) Dequeue() (T, error) {
	var empty T
	if CQ._size == 0 {
		return empty, errors.New("queue is empty")
	}

	element, err := CQ._storage.Get(CQ._front)
	if err != nil {
		return empty, err
	}

	if CQ._front == CQ._rear {
		CQ._front = -1
		CQ._rear = -1
	} else {
		CQ._front = (CQ._front + 1) % CQ._max
	}

	CQ._size--

	return element, nil
}

func (CQ *CircularQueue[T]) IsEmpty() bool {
	return CQ.Length() == 0
}

func (CQ *CircularQueue[T]) Length() int {
	return CQ._size
}

func (CQ *CircularQueue[T]) Clear() {
	CQ._storage.Clear()
	CQ._front, CQ._rear = -1, -1
	CQ._size = 0

	var empty T
	for i := 0; i < CQ._max; i++ {
		CQ._storage.Insert(i, empty)
	}

}
