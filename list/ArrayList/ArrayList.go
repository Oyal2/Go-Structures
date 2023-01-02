package ArrayList

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/oyal2/Go-Structures/list"
)

type ArrayList[T comparable] struct {
	_storage []T
	_length  int
	capacity int
}

func New[T comparable](size int) *ArrayList[T] {
	return &ArrayList[T]{
		_storage: make([]T, size),
		_length:  0,
		capacity: size,
	}
}

func (A *ArrayList[T]) Get(index int) (T, error) {
	var empty T
	if index < 0 || index >= A._length {
		return empty, errors.New("the index is out of bounds")
	}

	return A._storage[index], nil
}

func (A *ArrayList[T]) Update(index int, element T) (T, error) {
	var empty T
	if index < 0 || index >= A._length {
		return empty, errors.New("the index is out of bounds")
	}
	empty = A._storage[index]
	A._storage[index] = element
	return empty, nil
}

func (A *ArrayList[T]) reserve(newCapacity int) error {
	//New capacity needs to be greater than the old one
	if newCapacity <= A.capacity {
		return errors.New("the reserve capacity is less than or equal to your original capacity")
	}

	oldStorage := A._storage
	A._storage = make([]T, newCapacity)

	A.capacity = newCapacity
	for index, element := range oldStorage {
		A._storage[index] = element
	}

	return nil
}

func (A *ArrayList[T]) Insert(index int, element T) error {
	if index < 0 || index > A._length {
		return fmt.Errorf("index [%d] is out of bounds", index)
	}

	if A._length == A.capacity {
		err := A.reserve(list.Max(2*A.capacity, 1))
		if err != nil {
			return err
		}
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

// Appends the specified element(s) to the end of this list
func (A *ArrayList[T]) Add(elements ...T) error {
	for _, element := range elements {
		err := A.Insert(A.Length(), element)
		if err != nil {
			return err
		}
	}
	return nil
}

// Removes all of the elements from this list.
func (A *ArrayList[T]) Clear() {
	A._storage = make([]T, A._length)
	A._length = 0
}

// Contains Returns true if this list contains the specified element.
func (A *ArrayList[T]) Contains(element T) bool {
	for _, item := range A._storage {
		if item == element {
			return true
		}
	}
	return false
}

// Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (A *ArrayList[T]) IndexOf(element T) int {
	for i, item := range A._storage {
		if item == element {
			return i
		}
	}
	return -1
}

func (A *ArrayList[T]) Remove(index int) (T, error) {
	var value, empty T
	if !(index < A._length && index >= 0) {
		return value, fmt.Errorf("index [%d] is out of bounds", index)
	}

	value = A._storage[index]
	A._storage[index] = empty
	if index == A._length-1 {
		A._length--
		return value, nil
	}

	for i := index; i < A._length-1; i++ {
		if A._storage[i+1] != empty {
			A._storage[i] = A._storage[i+1]
		} else {
			break
		}
	}

	A._length--

	return value, nil
}

func (A *ArrayList[T]) Length() int { return A._length }

func (A *ArrayList[T]) Display() (string, error) {
	b, err := json.Marshal(A._storage)
	return string(b), err
}
