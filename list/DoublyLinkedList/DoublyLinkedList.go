package DoublyLinkedList

import (
	"github.com/oyal2/Go-Structures/list"
	"github.com/pkg/errors"
)

type DoublyLinkedList struct {
	_tailNode     *list.Node
	_middleNode   *list.Node
	_headNode     *list.Node
	_numberStored int
}

func New() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (Dl *DoublyLinkedList) Get(index int) (interface{}, error) {
	if 0 > index || index >= Dl._numberStored {
		return nil, errors.New("the index is out of bounds")
	}
	if index == Dl._numberStored-1 {
		//Retrieve the tail node value
		return Dl._tailNode.Value, nil
	}

	if Dl.Length()-index < index {
		currentNode := Dl._tailNode
		for i := Dl.Length() - 1; i != index; i-- {
			currentNode = currentNode.Prev
		}
		return currentNode.Value, nil
	} else {
		currentNode := Dl._headNode
		for i := 0; i < index; i++ {
			currentNode = currentNode.Next
		}
		return currentNode.Value, nil
	}

}

func (Dl *DoublyLinkedList) Update(index int, element interface{}) error {
	if 0 > index || index >= Dl._numberStored {
		return errors.New("the index is out of bounds")
	}
	if index == Dl._numberStored-1 {
		Dl._tailNode.Value = element
		return nil
	}

	if index == 0 {
		Dl._headNode.Value = element
		return nil
	}

	currentNode := Dl._headNode
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}
	currentNode.Value = element

	return nil
}

func (Dl *DoublyLinkedList) Insert(index int, element interface{}) error {
	if 0 > index || index > Dl._numberStored {
		return errors.New("the index is out of bounds")
	}
	newNode := &list.Node{
		Value: element,
	}
	//If the DL is empty
	if Dl._numberStored == 0 {
		Dl._headNode = newNode
		Dl._middleNode = newNode
		Dl._tailNode = newNode
	} else {
		if index == 0 {
			newNode.Next = Dl._headNode
			Dl._headNode.Prev = newNode
			Dl._headNode = newNode
		} else if index == Dl._numberStored {
			newNode.Prev = Dl._tailNode
			Dl._tailNode.Next = newNode
			Dl._tailNode = newNode
		} else if index >= Dl._numberStored/2 {
			currentNode := Dl._tailNode
			for i := Dl._numberStored - 1; index != i; i-- {
				currentNode = currentNode.Prev
			}
			newNode.Next = currentNode
			newNode.Prev = currentNode.Prev
			newNode.Prev.Next = newNode
			currentNode.Prev = newNode
		} else {
			currentNode := Dl._headNode
			for i := 0; i < index-1; i++ {
				currentNode = currentNode.Next
			}
			newNode.Prev = currentNode
			newNode.Next = currentNode.Next
			newNode.Next.Prev = newNode
			currentNode.Next = newNode
		}
	}

	Dl._numberStored++
	return nil
}

func (Dl *DoublyLinkedList) Add(elements ...interface{}) error {
	return Dl.Append(elements...)
}

func (Dl *DoublyLinkedList) Prepend(elements ...interface{}) error {
	for _, element := range elements {
		err := Dl.Insert(0, element)
		if err != nil {
			return err
		}
	}
	return nil
}

func (Dl *DoublyLinkedList) Append(elements ...interface{}) error {
	for _, element := range elements {
		err := Dl.Insert(Dl.Length(), element)
		if err != nil {
			return err
		}
	}
	return nil
}

func (Dl *DoublyLinkedList) Contains(element interface{}) bool {
	currentNode := Dl._headNode
	for currentNode != nil {
		if currentNode.Value == element {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

func (Dl *DoublyLinkedList) IndexOf(element interface{}) int {
	currentNode := Dl._headNode
	for i := 0; currentNode != nil; i++ {
		if currentNode.Value == element {
			return i
		}
		currentNode = currentNode.Next
	}
	return -1
}

func (Dl *DoublyLinkedList) Clear() {
	Dl._headNode = nil
	Dl._tailNode = nil
	Dl._middleNode = nil
	Dl._numberStored = 0
}

func (Dl *DoublyLinkedList) Empty() bool {
	return Dl.Length() == 0
}

func (Dl *DoublyLinkedList) Remove(index int) (interface{}, error) {
	if 0 > index || index >= Dl._numberStored {
		return nil, errors.New("the index is out of bounds")
	}
	var result interface{}
	if Dl._numberStored == 1 {
		result = Dl._headNode.Value
		Dl._headNode = nil
		Dl._tailNode = nil
	} else {
		if index == 0 {
			result = Dl._headNode.Value
			Dl._headNode = Dl._headNode.Next
			Dl._headNode.Prev = nil
		} else if index == Dl._numberStored-1 {
			result = Dl._tailNode.Value
			Dl._tailNode = Dl._tailNode.Prev
			Dl._tailNode.Next = nil
		} else if index >= Dl._numberStored/2 {
			currentNode := Dl._tailNode
			for i := Dl._numberStored - 1; index != i; i-- {
				currentNode = currentNode.Prev
			}
			result = currentNode.Value
			currentNode.Prev.Next = currentNode.Next
			currentNode.Next.Prev = currentNode.Prev
		} else {
			currentNode := Dl._headNode
			for i := 0; i != index; i++ {
				currentNode = currentNode.Next
			}
			result = currentNode.Value
			currentNode.Prev.Next = currentNode.Next
			currentNode.Next.Prev = currentNode.Prev
		}
	}
	Dl._numberStored--
	return result, nil
}

func (Dl *DoublyLinkedList) Length() int { return Dl._numberStored }
