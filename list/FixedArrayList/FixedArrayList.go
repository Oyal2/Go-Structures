package FixedArrayList

import (
	"errors"
	"fmt"

	"github.com/oyal2/Go-Structures/list"
	"github.com/oyal2/Go-Structures/utils"
)

type FixedArrayList[T utils.Ordered] struct {
	_storage []T
	_length  int
}

func New[T utils.Ordered](size int) FixedArrayList[T] {
	return FixedArrayList[T]{
		_storage: make([]T, size),
		_length:  0,
	}
}

func (A *FixedArrayList[T]) Get(index int) (T, error) {
	var empty T
	if list.Require(index, A._length) {
		return empty, errors.New("the index is out of bounds")
	}
	return A._storage[index], nil
}

func (A *FixedArrayList[T]) Update(index int, element T) error {
	if list.Require(index, A._length) {
		return errors.New("the index is out of bounds")
	}
	A._storage[index] = element
	return nil
}

func (A *FixedArrayList[T]) Insert(index int, element T) error {
	if index < 0 || index > A._length {
		return fmt.Errorf("index [%d] is out of bounds", index)
	}

	if cap(A._storage) == A._length {
		return errors.New("array capacity is full")
	}

	i := A._length - 1
	for i > index {
		A._storage[i+1] = A._storage[i]
		i--
	}

	A._storage[index] = element
	A._length++
	return nil
}

func (A *FixedArrayList[T]) Remove(index int) (value T, err error) {
	if !(index < A._length && index >= 0) {
		return value, fmt.Errorf("index [%d] is out of bounds", index)
	}

	value = A._storage[index]
	var empty T
	A._storage[index] = empty
	if index == A._length-1 {
		A._length--
		return
	}

	for i := index; i < A._length-1; i++ {
		if A._storage[i+1] != empty {
			A._storage[i] = A._storage[i+1]
		} else {
			break
		}
	}

	A._length--

	return
}

func (A *FixedArrayList[T]) Clear() {
	A._storage = make([]T, A._length)
	A._length = 0
}

func (A *FixedArrayList[T]) Length() int { return A._length }
