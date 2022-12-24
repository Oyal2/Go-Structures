package BinaryHeap

import (
	"errors"

	"github.com/oyal2/Go-Structures/utils"
)

type BinaryHeap[T utils.Ordered] struct {
	_storage    []T
	_size       int
	_comparator func(a, b T) bool
}

func New[T utils.Ordered](comparator func(a, b T) bool) *BinaryHeap[T] {
	return &BinaryHeap[T]{
		_storage:    make([]T, 1),
		_size:       0,
		_comparator: comparator,
	}
}

func (BH *BinaryHeap[T]) Insert(elements ...T) error {
	for _, element := range elements {
		err := BH.insert_items(element)
		if err != nil {
			return err
		}
	}
	return nil
}

func (BH *BinaryHeap[T]) insert_items(element T) error {

	if BH._size == len(BH._storage)-1 {
		oldStorage := BH._storage
		BH._storage = make([]T, len(BH._storage)*2)
		copy(BH._storage, oldStorage)
	}
	BH._size++

	pos := BH._size
	for ; pos > 1 && BH.Less(element, BH._storage[pos/2]); pos = pos / 2 {
		BH._storage[pos] = BH._storage[pos/2]
	}

	BH._storage[pos] = element

	return nil
}

func (BH *BinaryHeap[T]) Extract() (T, error) {
	var returnElement T
	var empty T
	if BH._size == 0 {
		return returnElement, errors.New("empty heap")
	}

	returnElement = BH._storage[1]
	settingElement := BH._storage[BH._size]
	pos := 1
	BH._storage[pos] = settingElement
	BH._storage[BH._size] = empty
	BH._size--

	if BH._size == 2 {
		if BH.Less(BH._storage[pos*2], BH._storage[pos]) {
			BH.Swap(pos, pos*2)
		}
	} else {
		for pos*2 < BH._size {
			if BH.Less(BH._storage[pos*2], BH._storage[pos*2+1]) {
				if BH.Less(BH._storage[pos*2], BH._storage[pos]) {
					BH.Swap(pos, pos*2)
				} else {
					break
				}
				pos = pos * 2
			} else {
				if BH.Less(BH._storage[pos*2+1], BH._storage[pos]) {
					BH.Swap(pos, pos*2+1)
				} else {
					break
				}
				pos = pos*2 + 1
			}
		}
	}

	return returnElement, nil
}

func (BH *BinaryHeap[T]) Less(a, b T) bool {
	return BH._comparator(a, b)
}

func (BH *BinaryHeap[T]) Swap(i, j int) {
	BH._storage[i], BH._storage[j] = BH._storage[j], BH._storage[i]
}

func (BH *BinaryHeap[T]) Contains(element T) bool {
	if BH._size == 0 {
		return false
	}

	for i := 1; i <= BH._size; i++ {
		if BH._storage[i] == element {
			return true
		}
	}
	return false
}

func (BH *BinaryHeap[T]) Peek() (T, error) {
	var returnValue T
	if BH._size == 0 {
		return returnValue, errors.New("empty heap")
	}

	return BH._storage[1], nil
}

func (BH *BinaryHeap[T]) ChangeKey(element T, newElement T) error {
	if BH._size == 0 {
		return errors.New("empty heap")
	}
	pos := 1
	for ; pos <= BH._size; pos++ {
		if BH._storage[pos] == element {
			break
		}
	}

	if pos == BH._size+1 {
		return errors.New("element doesnt exist")
	}

	BH._storage[pos] = newElement

	if BH._size == 2 {
		if BH.Less(BH._storage[pos*2], BH._storage[pos]) {
			BH.Swap(pos, pos*2)
		}
	} else {
		if BH.Less(newElement, BH._storage[pos/2]) {
			for ; pos > 1 && BH.Less(element, BH._storage[pos/2]); pos = pos / 2 {
				BH._storage[pos] = BH._storage[pos/2]
				BH.Swap(pos, pos/2)
			}
		} else {
			for pos*2 < BH._size {
				if BH.Less(BH._storage[pos*2], BH._storage[pos*2+1]) {
					if BH.Less(BH._storage[pos*2], BH._storage[pos]) {
						BH.Swap(pos, pos*2)
					} else {
						break
					}
					pos = pos * 2
				} else {
					if BH.Less(BH._storage[pos*2+1], BH._storage[pos]) {
						BH.Swap(pos, pos*2+1)
					} else {
						break
					}
					pos = pos*2 + 1
				}
			}
		}
	}

	return nil
}

func (BH *BinaryHeap[T]) Length() int {
	return BH._size
}

func (BH *BinaryHeap[T]) IsEmpty() bool {
	return BH._size == 0
}

func (BH *BinaryHeap[T]) Clear() {
	BH._storage = make([]T, 1)
	BH._size = 0
}
