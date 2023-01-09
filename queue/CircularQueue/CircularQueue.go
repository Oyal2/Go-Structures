package CircularQueue

import (
	"errors"

	"github.com/oyal2/Go-Structures/list"
)

type CircularQueue[T comparable] struct {
	storage list.List[T]
	max     int
	front   int
	rear    int
	size    int
}

func New[T comparable](storage list.List[T], max int) *CircularQueue[T] {
	cq := &CircularQueue[T]{
		storage: storage,
		max:     max,
		front:   -1,
		rear:    -1,
	}

	var empty T
	for i := 0; i < max; i++ {
		cq.storage.Insert(i, empty)
	}

	return cq
}

func (CQ *CircularQueue[T]) Enqueue(elements ...T) error {
	if len(elements)+CQ.Length() > CQ.max {
		return errors.New("queue is full")
	}

	for _, element := range elements {

		if CQ.front == -1 {
			CQ.front = 0
		}
		CQ.rear = (CQ.rear + 1) % CQ.max

		_, err := CQ.storage.Update(CQ.rear, element)
		if err != nil {
			return err
		}

		CQ.size++
	}

	return nil
}

func (CQ *CircularQueue[T]) Front() (T, error) {
	var empty T
	if CQ.Length() == 0 {
		return empty, errors.New("queue is empty")
	}

	return CQ.storage.Get(CQ.front)
}

func (CQ *CircularQueue[T]) Dequeue() (T, error) {
	var empty T
	if CQ.size == 0 {
		return empty, errors.New("queue is empty")
	}

	element, err := CQ.storage.Get(CQ.front)
	if err != nil {
		return empty, err
	}

	if CQ.front == CQ.rear {
		CQ.front = -1
		CQ.rear = -1
	} else {
		CQ.front = (CQ.front + 1) % CQ.max
	}

	CQ.size--

	return element, nil
}

func (CQ *CircularQueue[T]) IsEmpty() bool {
	return CQ.Length() == 0
}

func (CQ *CircularQueue[T]) Length() int {
	return CQ.size
}

func (CQ *CircularQueue[T]) Clear() {
	CQ.storage.Clear()
	CQ.front, CQ.rear = -1, -1
	CQ.size = 0

	var empty T
	for i := 0; i < CQ.max; i++ {
		CQ.storage.Insert(i, empty)
	}

}
