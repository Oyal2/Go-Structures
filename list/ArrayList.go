package list

import (
	"fmt"
	"github.com/pkg/errors"
)

var require = func(a, b int) bool { return a >= b }
var max = func(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type ArrayList struct {
	_storage []interface{}
	_length  int
	capacity int
}

func newArrayList(size int) ArrayList {
	return ArrayList{
		_storage: make([]interface{}, size),
		_length:  0,
		capacity: size,
	}
}

func (A *ArrayList) Get(index int) (interface{}, error) {
	if require(index, A._length) {
		return nil, errors.New("the index is out of bounds")
	}
	return A._storage[index], nil
}

func (A *ArrayList) Update(index int, element interface{}) error {
	if require(index, A._length) {
		return errors.New("the index is out of bounds")
	}
	A._storage[index] = element
	return nil
}

func (A *ArrayList) Reserve(newCapacity int) error {
	if newCapacity >= A.capacity {
		return errors.New("the reserve capacity is less than your original capacity")
	}

	oldStorage := A._storage
	A._storage = make([]interface{}, newCapacity)

	A.capacity = newCapacity
	for index, element := range oldStorage {
		A._storage[index] = element
	}

	return nil
}

func (A *ArrayList) Insert(index int, element interface{}) error {
	if index < 0 || index > A._length {
		return fmt.Errorf("index [%d] is out of bounds", index)
	}

	if A._length == A.capacity {
		A.Reserve(max(2*A.capacity, 1))
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

func (A *ArrayList) Remove(index int) (interface{}, error) {
	if index < 0 || index > A._length {
		return nil, fmt.Errorf("index [%d] is out of bounds", index)
	}

	val := A._storage[index]
	i := index
	for i < A._length {
		A._storage[i] = A._storage[i+1]
		i++
	}
	A._length--
	return val, nil
}

func (A *ArrayList) Length() int { return A._length }
