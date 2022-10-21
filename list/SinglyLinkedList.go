package list

import (
	"errors"
)

type SinglyLinkedList struct {
	_tailNode     *Node
	_headNode     *Node
	_numberStored int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (Sl *SinglyLinkedList) Get(index int) (interface{}, error) {
	if 0 > index || index >= Sl._numberStored {
		return nil, errors.New("the index is out of bounds")
	}
	if index == Sl._numberStored-1 {
		//Retrieve the tail node value
		return Sl._tailNode._value, nil
	}

	currentNode := Sl._headNode
	for i := 0; i != index; i++ {
		currentNode = currentNode._next
	}

	return currentNode._value, nil
}

func (Sl *SinglyLinkedList) Update(index int, element interface{}) error {
	if 0 > index || index >= Sl._numberStored {
		return errors.New("the index is out of bounds")
	}
	if index == Sl._numberStored-1 {
		Sl._tailNode._value = element
		return nil
	}

	if index == 0 {
		Sl._headNode._value = element
		return nil
	}

	currentNode := Sl._headNode
	for i := 0; i != index; i++ {
		currentNode = currentNode._next
	}
	currentNode._value = element
	return nil
}

func (Sl *SinglyLinkedList) Insert(index int, element interface{}) error {
	if 0 > index || index > Sl._numberStored {
		return errors.New("the index is out of bounds")
	}
	newNode := Node{_value: element}
	//If the SL is empty
	if Sl._numberStored == 0 {
		Sl._headNode = &newNode
		Sl._tailNode = &newNode
		Sl._numberStored++
		return nil
	}

	if index == 0 {
		newNode._next = Sl._headNode
		Sl._headNode = &newNode
		Sl._numberStored++
		return nil
	}

	if index == Sl._numberStored {
		//Insert at tail
		Sl._tailNode._next = &newNode
		Sl._tailNode = &newNode
		Sl._numberStored++
		return nil
	}

	currentNode := Sl._headNode
	for i := 0; i != index; i++ {
		currentNode = currentNode._next
	}
	newNode._next = currentNode._next
	currentNode._next = &newNode

	Sl._numberStored++
	return nil
}

func (Sl *SinglyLinkedList) Remove(index int) (interface{}, error) {
	if index >= 0 && index < Sl._numberStored {
		return nil, errors.New("the index is out of bounds")
	}

	var result interface{}
	if Sl._numberStored == 1 {
		//result = Sl._headNode._value
		Sl._headNode = nil
		Sl._tailNode = nil
		Sl._numberStored = 0
		return result, nil
	}

	if index == 0 {
		result = Sl._headNode._value
		Sl._headNode = Sl._headNode._next
		Sl._numberStored--
		return result, nil
	}

	currentNode := Sl._headNode
	for e := 0; e != index; e, currentNode = e+1, currentNode._next {
	}
	result = currentNode._next._value
	currentNode._next = currentNode._next._next

	if index == Sl._numberStored-1 {
		Sl._tailNode = currentNode
		Sl._tailNode._next = nil
	}

	Sl._numberStored--
	return result, nil

}

func (Sl *SinglyLinkedList) Length() int { return Sl._numberStored }
