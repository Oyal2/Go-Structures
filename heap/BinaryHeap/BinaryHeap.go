package BinaryHeap

import (
	"errors"
	"sync"

	"github.com/oyal2/Go-Structures/utils"
)

type BinaryHeap[T utils.Ordered] struct {
	storage    []T
	size       int
	comparator func(a, b T) bool
	sync.RWMutex
}

func New[T utils.Ordered](comparator func(a, b T) bool) *BinaryHeap[T] {
	return &BinaryHeap[T]{
		storage:    make([]T, 1),
		size:       0,
		comparator: comparator,
	}
}

func (BH *BinaryHeap[T]) Insert(elements ...T) error {
	BH.Lock()
	defer BH.Unlock()
	for _, element := range elements {
		err := BH.insert_items(element)
		if err != nil {
			return err
		}
	}
	return nil
}

func (BH *BinaryHeap[T]) insert_items(element T) error {

	if BH.size == len(BH.storage)-1 {
		oldStorage := BH.storage
		BH.storage = make([]T, len(BH.storage)*2)
		copy(BH.storage, oldStorage)
	}
	BH.size++

	pos := BH.size
	for ; pos > 1 && BH.Less(element, BH.storage[pos/2]); pos = pos / 2 {
		BH.storage[pos] = BH.storage[pos/2]
	}

	BH.storage[pos] = element

	return nil
}

func (BH *BinaryHeap[T]) Extract() (T, error) {
	BH.Lock()
	defer BH.Unlock()
	var returnElement T
	var empty T
	if BH.size == 0 {
		return returnElement, errors.New("empty heap")
	}

	returnElement = BH.storage[1]
	settingElement := BH.storage[BH.size]
	pos := 1
	BH.storage[pos] = settingElement
	BH.storage[BH.size] = empty
	BH.size--

	if BH.size == 2 {
		if BH.Less(BH.storage[pos*2], BH.storage[pos]) {
			BH.Swap(pos, pos*2)
		}
	} else {
		for pos*2 < BH.size {
			if BH.Less(BH.storage[pos*2], BH.storage[pos*2+1]) {
				if BH.Less(BH.storage[pos*2], BH.storage[pos]) {
					BH.Swap(pos, pos*2)
				} else {
					break
				}
				pos = pos * 2
			} else {
				if BH.Less(BH.storage[pos*2+1], BH.storage[pos]) {
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
	return BH.comparator(a, b)
}

func (BH *BinaryHeap[T]) Swap(i, j int) {
	BH.storage[i], BH.storage[j] = BH.storage[j], BH.storage[i]
}

func (BH *BinaryHeap[T]) Contains(element T) bool {
	BH.RLock()
	defer BH.RUnlock()
	if BH.size == 0 {
		return false
	}

	for i := 1; i <= BH.size; i++ {
		if BH.storage[i] == element {
			return true
		}
	}
	return false
}

func (BH *BinaryHeap[T]) Peek() (T, error) {
	BH.RLock()
	defer BH.RUnlock()
	var returnValue T
	if BH.size == 0 {
		return returnValue, errors.New("empty heap")
	}

	return BH.storage[1], nil
}

func (BH *BinaryHeap[T]) ChangeKey(element T, newElement T) error {
	BH.Lock()
	defer BH.Unlock()
	if BH.size == 0 {
		return errors.New("empty heap")
	}
	pos := 1
	for ; pos <= BH.size; pos++ {
		if BH.storage[pos] == element {
			break
		}
	}

	if pos == BH.size+1 {
		return errors.New("element doesnt exist")
	}

	BH.storage[pos] = newElement

	if BH.size == 2 {
		if BH.Less(BH.storage[pos*2], BH.storage[pos]) {
			BH.Swap(pos, pos*2)
		}
	} else {
		if BH.Less(newElement, BH.storage[pos/2]) {
			for ; pos > 1 && BH.Less(element, BH.storage[pos/2]); pos = pos / 2 {
				BH.storage[pos] = BH.storage[pos/2]
				BH.Swap(pos, pos/2)
			}
		} else {
			for pos*2 < BH.size {
				if BH.Less(BH.storage[pos*2], BH.storage[pos*2+1]) {
					if BH.Less(BH.storage[pos*2], BH.storage[pos]) {
						BH.Swap(pos, pos*2)
					} else {
						break
					}
					pos = pos * 2
				} else {
					if BH.Less(BH.storage[pos*2+1], BH.storage[pos]) {
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
	BH.RLock()
	defer BH.RUnlock()
	return BH.size
}

func (BH *BinaryHeap[T]) IsEmpty() bool {
	BH.RLock()
	defer BH.RUnlock()
	return BH.size == 0
}

func (BH *BinaryHeap[T]) Clear() {
	BH.Lock()
	defer BH.Unlock()
	BH.storage = make([]T, 1)
	BH.size = 0
}
