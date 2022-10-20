package list

import (
	"github.com/pkg/errors"
)

type DoublyLinkedList struct {
	_tailNode     *Node
	_headNode     *Node
	_numberStored *int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	i := 0
	return &DoublyLinkedList{
		_numberStored: &i,
	}
}

func (Dl *DoublyLinkedList) Get(index int) (interface{}, error) {
	if 0 > index || index >= *Dl._numberStored {
		return nil, errors.New("the index is out of bounds")
	}
	if index == *Dl._numberStored-1 {
		//Retrieve the tail node value
		return Dl._tailNode._value, nil
	}

	if Dl.Length()-index < index {
		currentNode := Dl._tailNode
		for i := Dl.Length() - 1; i != index; i-- {
			currentNode = currentNode._prev
		}
		return currentNode._value, nil
	} else {
		currentNode := Dl._headNode
		for i := 0; i < index; i++ {
			currentNode = currentNode._next
		}
		return currentNode._value, nil
	}

}

func (Dl *DoublyLinkedList) Update(index int, element interface{}) error {
	if 0 > index || index >= *Dl._numberStored {
		return errors.New("the index is out of bounds")
	}
	if index == *Dl._numberStored-1 {
		Dl._tailNode._value = element
		return nil
	}

	if index == 0 {
		Dl._headNode._value = element
		return nil
	}

	currentNode := Dl._headNode
	for i := 0; i < index; i++ {
		currentNode = currentNode._next
	}
	currentNode._value = element

	return nil
}

func (Dl *DoublyLinkedList) Insert(index int, element interface{}) error {
	if 0 > index || index > *Dl._numberStored {
		return errors.New("the index is out of bounds")
	}
	newNode := &Node{
		_value: element,
	}
	//If the DL is empty
	if *Dl._numberStored == 0 {
		Dl._headNode = newNode
		Dl._tailNode = newNode
	} else {
		if index == 0 {
			newNode._next = Dl._headNode
			Dl._headNode._prev = newNode
			Dl._headNode = newNode
		} else if index == *Dl._numberStored {
			newNode._prev = Dl._tailNode
			Dl._tailNode._next = newNode
			Dl._tailNode = newNode
		} else if index >= *Dl._numberStored/2 {
			currentNode := Dl._tailNode
			for i := *Dl._numberStored - 1; index != i; i-- {
				currentNode = currentNode._prev
			}
			newNode._next = currentNode
			newNode._prev = currentNode._prev
			newNode._prev._next = newNode
			currentNode._prev = newNode
		} else {
			currentNode := Dl._headNode
			for i := 0; i < index-1; i++ {
				currentNode = currentNode._next
			}
			newNode._prev = currentNode
			newNode._next = currentNode._next
			newNode._next._prev = newNode
			currentNode._next = newNode
		}
	}
	*Dl._numberStored++
	return nil
}

func (Dl *DoublyLinkedList) Remove(index int) (interface{}, error) {
	if 0 > index || index >= *Dl._numberStored {
		return nil, errors.New("the index is out of bounds")
	}
	var result interface{}
	if *Dl._numberStored == 1 {
		result = Dl._headNode._value
		Dl._headNode = nil
		Dl._tailNode = nil
	} else {
		if index == 0 {
			result = Dl._headNode._value
			Dl._headNode = Dl._headNode._next
			Dl._headNode._prev = nil
		} else if index == *Dl._numberStored-1 {
			result = Dl._tailNode._value
			Dl._tailNode = Dl._tailNode._prev
			Dl._tailNode._next = nil
		} else if index >= *Dl._numberStored/2 {
			currentNode := Dl._tailNode
			for i := *Dl._numberStored - 1; index != i; i-- {
				currentNode = currentNode._prev
			}
			result = currentNode._value
			currentNode._prev._next = currentNode._next
			currentNode._next._prev = currentNode._prev
		} else {
			currentNode := Dl._headNode
			for i := 0; i != index; i++ {
				currentNode = currentNode._next
			}
			result = currentNode._value
			currentNode._prev._next = currentNode._next
			currentNode._next._prev = currentNode._prev
		}
	}
	*Dl._numberStored--
	return result, nil
}

func (Dl *DoublyLinkedList) Length() int { return *Dl._numberStored }
