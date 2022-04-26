package list

import (
	"github.com/pkg/errors"
)

type Dnode struct {
	_value        interface{}
	_tailNode     *Dnode
	_headNode     *Dnode
	_next         *Dnode
	_prev         *Dnode
	_numberStored *int
}

func NewDoublyLinkedList() Dnode {
	i := 0
	return Dnode{
		_value:        nil,
		_tailNode:     nil,
		_headNode:     nil,
		_next:         nil,
		_prev:         nil,
		_numberStored: &i,
	}
}

func (Dl *Dnode) Get(index int) (interface{}, error) {
	if 0 > index || index >= *Dl._numberStored {
		return nil, errors.New("the index is out of bounds")
	}
	if index == *Dl._numberStored-1 {
		//Retrieve the tail node value
		return Dl._tailNode._value, nil
	}

	currentNode := Dl._headNode
	for i := 0; i < index; i++ {
		currentNode = currentNode._next
	}

	return currentNode._value, nil
}

func (Dl *Dnode) Update(index int, element interface{}) error {
	if 0 > index || index >= *Dl._numberStored {
		return errors.New("the index is out of bounds")
	}
	if index == *Dl._numberStored-1 {
		Dl._tailNode._value = element
		return nil
	}

	if index == 0 {
		Dl._value = element
	}

	currentNode := Dl._headNode
	for i := 0; i < index; i++ {
		currentNode = currentNode._next
	}
	currentNode._value = element

	return nil
}

func (Dl *Dnode) Insert(index int, element interface{}) error {
	if 0 > index || index > *Dl._numberStored {
		return errors.New("the index is out of bounds")
	}
	newNode := NewDoublyLinkedList()
	newNode._value = element
	//If the DL is empty
	if *Dl._numberStored == 0 {
		Dl._headNode = &newNode
		Dl._tailNode = &newNode
		Dl._value = element
		*Dl._numberStored++
		newNode._numberStored = Dl._numberStored
		return nil
	} else {
		if index == 0 {
			newNode._next = Dl._headNode
			Dl._headNode = &newNode
			Dl._headNode._prev = Dl._headNode
		} else if index < *Dl._numberStored {
			currentNode := Dl._headNode
			for i := 0; i < (index - 1); i++ {
				currentNode = currentNode._next
			}
			newNode._next = currentNode._next
			newNode._prev = currentNode
			currentNode._next = &newNode
		} else {
			//Insert at tail
			newNode._prev = Dl._tailNode
			newNode._next = nil
			newNode._headNode = Dl._headNode
			Dl._tailNode._next = &newNode
			Dl._tailNode = Dl._tailNode._next
			if 1 == Dl.Length() {
				Dl._next = Dl._tailNode
			}
		}
		*Dl._numberStored++
		newNode._numberStored = Dl._numberStored
	}
	return nil
}

func (Dl *Dnode) Remove(index int) (interface{}, error) {
	if 0 > index || index >= *Dl._numberStored {
		return nil, errors.New("the index is out of bounds")
	}

	if *Dl._numberStored == 1 {
		result := Dl._headNode._value
		Dl._headNode = nil
		Dl._tailNode = nil
		*Dl._numberStored--
		return result, nil
	} else {
		if index == 0 {
			result := Dl._headNode._value
			Dl._headNode = Dl._headNode._next
			Dl._headNode._prev = nil
			*Dl._numberStored--
			return result, nil
		} else if index == *Dl._numberStored-1 {
			result := Dl._tailNode._value
			Dl._tailNode = Dl._tailNode._prev
			Dl._tailNode._next = nil
			*Dl._numberStored--
			return result, nil
		} else {
			currentNode := Dl._headNode
			for i := 0; i < (index - 1); i++ {
				currentNode = currentNode._next
			}
			result := currentNode._next._value
			currentNode._next = currentNode._next._next
			currentNode._next._prev = currentNode
			*Dl._numberStored--
			return result, nil
		}
	}
}

func (Dl *Dnode) Length() int { return *Dl._numberStored }
