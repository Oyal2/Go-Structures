package ArrayList

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/oyal2/Go-Structures/list"
)

type ArrayList[T comparable] struct {
	storage []T
	length  int
	capacity int
}

func New[T comparable](size int) *ArrayList[T] {
	return &ArrayList[T]{
		storage: make([]T, size),
		length:  0,
		capacity: size,
	}
}

func (A *ArrayList[T]) Get(index int) (T, error) {
	var empty T
	if index < 0 || index >= A.length {
		return empty, errors.New("the index is out of bounds")
	}

	return A.storage[index], nil
}

func (A *ArrayList[T]) Update(index int, element T) (T, error) {
	var empty T
	if index < 0 || index >= A.length {
		return empty, errors.New("the index is out of bounds")
	}
	empty = A.storage[index]
	A.storage[index] = element
	return empty, nil
}

func (A *ArrayList[T]) reserve(newCapacity int) error {
	//New capacity needs to be greater than the old one
	if newCapacity <= A.capacity {
		return errors.New("the reserve capacity is less than or equal to your original capacity")
	}

	oldStorage := A.storage
	A.storage = make([]T, newCapacity)

	A.capacity = newCapacity
	for index, element := range oldStorage {
		A.storage[index] = element
	}

	return nil
}

func (A *ArrayList[T]) Insert(index int, element T) error {
	if index < 0 || index > A.length {
		return fmt.Errorf("index [%d] is out of bounds", index)
	}

	if A.length == A.capacity {
		err := A.reserve(list.Max(2*A.capacity, 1))
		if err != nil {
			return err
		}
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
	A.storage = make([]T, A.length)
	A.length = 0
}

// Contains Returns true if this list contains the specified element.
func (A *ArrayList[T]) Contains(element T) bool {
	for _, item := range A.storage {
		if item == element {
			return true
		}
	}
	return false
}

// Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (A *ArrayList[T]) IndexOf(element T) int {
	for i, item := range A.storage {
		if item == element {
			return i
		}
	}
	return -1
}

func (A *ArrayList[T]) Remove(index int) (T, error) {
	var value, empty T
	if !(index < A.length && index >= 0) {
		return value, fmt.Errorf("index [%d] is out of bounds", index)
	}

	value = A.storage[index]
	A.storage[index] = empty
	if index == A.length-1 {
		A.length--
		return value, nil
	}

	for i := index; i < A.length-1; i++ {
		if A.storage[i+1] != empty {
			A.storage[i] = A.storage[i+1]
		} else {
			break
		}
	}

	A.length--

	return value, nil
}

func (A *ArrayList[T]) Length() int { return A.length }

func (A *ArrayList[T]) Display() (string, error) {
	b, err := json.Marshal(A.storage)
	return string(b), err
}
