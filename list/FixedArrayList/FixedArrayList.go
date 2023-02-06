package FixedArrayList

import (
	"errors"
	"fmt"

	"github.com/oyal2/Go-Structures/list"
	"github.com/oyal2/Go-Structures/utils"
)

type FixedArrayList[T utils.Ordered] struct {
	storage []T
	length  int
}

func New[T utils.Ordered](size int) FixedArrayList[T] {
	return FixedArrayList[T]{
		storage: make([]T, size),
		length:  0,
	}
}

func (A *FixedArrayList[T]) Get(index int) (T, error) {
	var empty T
	if list.Require(index, A.length) {
		return empty, errors.New("the index is out of bounds")
	}
	return A.storage[index], nil
}

func (A *FixedArrayList[T]) Update(index int, element T) (T, error) {
	var empty T
	if list.Require(index, A.length) {
		return empty, errors.New("the index is out of bounds")
	}
	empty = A.storage[index]
	A.storage[index] = element
	return empty, nil
}

func (A *FixedArrayList[T]) Insert(index int, element T) error {
	if index < 0 || index > A.length {
		return fmt.Errorf("index [%d] is out of bounds", index)
	}

	if cap(A.storage) == A.length {
		return errors.New("array capacity is full")
	}

	i := A.length - 1
	for i > index {
		A.storage[i+1] = A.storage[i]
		i--
	}

	A.storage[index] = element
	A.length++
	return nil
}

func (A *FixedArrayList[T]) Remove(index int) (value T, err error) {
	if !(index < A.length && index >= 0) {
		return value, fmt.Errorf("index [%d] is out of bounds", index)
	}

	value = A.storage[index]
	var empty T
	A.storage[index] = empty
	if index == A.length-1 {
		A.length--
		return
	}

	for i := index; i < A.length-1; i++ {
		if A.storage[i+1] != empty {
			A.storage[i] = A.storage[i+1]
		} else {
			break
		}
	}

	A.length--

	return
}

func (A *FixedArrayList[T]) Clear() {
	A.storage = make([]T, A.length)
	A.length = 0
}

func (A *FixedArrayList[T]) IndexOf(element T) int {
	for i, item := range A.storage {
		if item == element {
			return i
		}
	}

	return -1
}

func (A *FixedArrayList[T]) Length() int { return A.length }
