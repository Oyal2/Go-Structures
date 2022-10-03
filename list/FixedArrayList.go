package list

import (
	"errors"
	"fmt"
)

type FixedArrayList struct {
	_storage []interface{}
	_length  int
}

func NewFixedArrayList(size int) FixedArrayList {
	return FixedArrayList{
		_storage: make([]interface{},size),
		_length:  0,
	}
}

func (A *FixedArrayList) Get(index int) (interface{}, error) {
	if Require(index, A._length) {
		return nil, errors.New("the index is out of bounds")
	}
	return A._storage[index], nil
}

func (A *FixedArrayList) Set(index int, element interface{}) error {
	if Require(index, A._length) {
		return errors.New("the index is out of bounds")
	}
	A._storage[index] = element
	return nil
}

func (A *FixedArrayList) Insert(index int, element interface{}) error {
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

func (A *FixedArrayList) Length() int { return A._length }