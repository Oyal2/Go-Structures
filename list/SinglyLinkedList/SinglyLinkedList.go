package SinglyLinkedList

import (
	"errors"
	"github.com/oyal2/Go-Structures/list"
)

type SinglyLinkedList struct {
	_tailNode     *list.Node
	_headNode     *list.Node
	_numberStored int
}

func New() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (Sl *SinglyLinkedList) Get(index int) (interface{}, error) {
	if 0 > index || index >= Sl._numberStored {
		return nil, errors.New("the index is out of bounds")
	}
	if index == Sl._numberStored-1 {
		//Retrieve the tail node value
		return Sl._tailNode.Value, nil
	}

	currentNode := Sl._headNode
	for i := 0; i != index; i++ {
		currentNode = currentNode.Next
	}

	return currentNode.Value, nil
}

func (Sl *SinglyLinkedList) Update(index int, element interface{}) error {
	if 0 > index || index >= Sl._numberStored {
		return errors.New("the index is out of bounds")
	}
	if index == Sl._numberStored-1 {
		Sl._tailNode.Value = element
		return nil
	}

	if index == 0 {
		Sl._headNode.Value = element
		return nil
	}

	currentNode := Sl._headNode
	for i := 0; i != index; i++ {
		currentNode = currentNode.Next
	}
	currentNode.Value = element
	return nil
}

func (Sl *SinglyLinkedList) Insert(index int, element interface{}) error {
	if 0 > index || index > Sl._numberStored {
		return errors.New("the index is out of bounds")
	}
	newNode := list.Node{Value: element}
	//If the SL is empty
	if Sl._numberStored == 0 {
		Sl._headNode = &newNode
		Sl._tailNode = &newNode
		Sl._numberStored++
		return nil
	}

	if index == 0 {
		newNode.Next = Sl._headNode
		Sl._headNode = &newNode
		Sl._numberStored++
		return nil
	}

	if index == Sl._numberStored {
		//Insert at tail
		Sl._tailNode.Next = &newNode
		Sl._tailNode = &newNode
		Sl._numberStored++
		return nil
	}

	currentNode := Sl._headNode
	for i := 0; i != index; i++ {
		currentNode = currentNode.Next
	}
	newNode.Next = currentNode.Next
	currentNode.Next = &newNode

	Sl._numberStored++
	return nil
}

func (Sl *SinglyLinkedList) Add(elements ...interface{}) error {
	return Sl.Append(elements...)
}

func (Sl *SinglyLinkedList) Prepend(elements ...interface{}) error {
	for _, element := range elements {
		err := Sl.Insert(0, element)
		if err != nil {
			return err
		}
	}
	return nil
}

func (Sl *SinglyLinkedList) Append(elements ...interface{}) error {
	for _, element := range elements {
		err := Sl.Insert(Sl.Length(), element)
		if err != nil {
			return err
		}
	}
	return nil
}

func (Sl *SinglyLinkedList) Contains(element interface{}) bool {
	currentNode := Sl._headNode
	for currentNode != nil {
		if currentNode.Value == element {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

func (Sl *SinglyLinkedList) IndexOf(element interface{}) int {
	currentNode := Sl._headNode
	for i := 0; currentNode != nil; i++ {
		if currentNode.Value == element {
			return i
		}
		currentNode = currentNode.Next
	}
	return -1
}

func (Sl *SinglyLinkedList) Clear() {
	Sl._headNode = nil
	Sl._tailNode = nil
	Sl._numberStored = 0
}

func (Sl *SinglyLinkedList) Empty() bool {
	return Sl.Length() == 0
}

func (Sl *SinglyLinkedList) Remove(index int) (interface{}, error) {
	if 0 > index || index >= Sl._numberStored {
		return nil, errors.New("the index is out of bounds")
	}

	var result interface{}
	if Sl._numberStored == 1 {
		result = Sl._headNode.Value
		Sl._headNode = nil
		Sl._tailNode = nil
		Sl._numberStored = 0
		return result, nil
	}

	if index == 0 {
		result = Sl._headNode.Value
		Sl._headNode = Sl._headNode.Next
		Sl._numberStored--
		return result, nil
	}

	currentNode := Sl._headNode
	for e := 0; e != index-1; e, currentNode = e+1, currentNode.Next {
	}

	result = currentNode.Next.Value
	currentNode.Next = currentNode.Next.Next

	if index == Sl._numberStored-1 {
		Sl._tailNode = currentNode
		Sl._tailNode.Next = nil
	}

	Sl._numberStored--
	return result, nil

}

func (Sl *SinglyLinkedList) Length() int { return Sl._numberStored }
