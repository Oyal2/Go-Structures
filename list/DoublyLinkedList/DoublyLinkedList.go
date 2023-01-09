package DoublyLinkedList

import (
	"github.com/oyal2/Go-Structures/list"
	"github.com/pkg/errors"
)

type DoublyLinkedList[T comparable] struct {
	tailNode     *list.Node[T]
	middleNode   *list.Node[T]
	headNode     *list.Node[T]
	numberStored int
}

func New[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (Dl *DoublyLinkedList[T]) Get(index int) (value T, err error) {
	if 0 > index || index >= Dl.numberStored {
		return value, errors.New("the index is out of bounds")
	}
	if index == Dl.numberStored-1 {
		//Retrieve the tail node value
		return Dl.tailNode.Value, nil
	}

	if Dl.Length()-index < index {
		currentNode := Dl.tailNode
		for i := Dl.Length() - 1; i != index; i-- {
			currentNode = currentNode.Prev
		}
		return currentNode.Value, nil
	} else {
		currentNode := Dl.headNode
		for i := 0; i < index; i++ {
			currentNode = currentNode.Next
		}
		return currentNode.Value, nil
	}

}

func (Dl *DoublyLinkedList[T]) Update(index int, element T) (T, error) {
	var empty T
	if 0 > index || index >= Dl.numberStored {
		return empty, errors.New("the index is out of bounds")
	}
	if index == Dl.numberStored-1 {
		empty = Dl.tailNode.Value
		Dl.tailNode.Value = element
		return empty, nil
	}

	if index == 0 {
		empty = Dl.headNode.Value
		Dl.headNode.Value = element
		return empty, nil
	}

	currentNode := Dl.headNode
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}
	empty = currentNode.Value
	currentNode.Value = element

	return empty, nil
}

func (Dl *DoublyLinkedList[T]) Insert(index int, element T) error {
	if 0 > index || index > Dl.numberStored {
		return errors.New("the index is out of bounds")
	}
	newNode := &list.Node[T]{
		Value: element,
	}
	//If the DL is empty
	if Dl.numberStored == 0 {
		Dl.headNode = newNode
		Dl.middleNode = newNode
		Dl.tailNode = newNode
	} else {
		if index == 0 {
			newNode.Next = Dl.headNode
			Dl.headNode.Prev = newNode
			Dl.headNode = newNode
		} else if index == Dl.numberStored {
			newNode.Prev = Dl.tailNode
			Dl.tailNode.Next = newNode
			Dl.tailNode = newNode
		} else if index >= Dl.numberStored/2 {
			currentNode := Dl.tailNode
			for i := Dl.numberStored - 1; index != i; i-- {
				currentNode = currentNode.Prev
			}
			newNode.Next = currentNode
			newNode.Prev = currentNode.Prev
			newNode.Prev.Next = newNode
			currentNode.Prev = newNode
		} else {
			currentNode := Dl.headNode
			for i := 0; i < index-1; i++ {
				currentNode = currentNode.Next
			}
			newNode.Prev = currentNode
			newNode.Next = currentNode.Next
			newNode.Next.Prev = newNode
			currentNode.Next = newNode
		}
	}

	Dl.numberStored++
	return nil
}

func (Dl *DoublyLinkedList[T]) Add(elements ...T) error {
	return Dl.Append(elements...)
}

func (Dl *DoublyLinkedList[T]) Prepend(elements ...T) error {
	for _, element := range elements {
		err := Dl.Insert(0, element)
		if err != nil {
			return err
		}
	}
	return nil
}

func (Dl *DoublyLinkedList[T]) Append(elements ...T) error {
	for _, element := range elements {
		err := Dl.Insert(Dl.Length(), element)
		if err != nil {
			return err
		}
	}
	return nil
}

func (Dl *DoublyLinkedList[T]) Contains(element T) bool {
	currentNode := Dl.headNode
	for currentNode != nil {
		if currentNode.Value == element {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

func (Dl *DoublyLinkedList[T]) IndexOf(element T) int {
	currentNode := Dl.headNode
	for i := 0; currentNode != nil; i++ {
		if currentNode.Value == element {
			return i
		}
		currentNode = currentNode.Next
	}
	return -1
}

func (Dl *DoublyLinkedList[T]) Clear() {
	Dl.headNode = nil
	Dl.tailNode = nil
	Dl.middleNode = nil
	Dl.numberStored = 0
}

func (Dl *DoublyLinkedList[T]) Empty() bool {
	return Dl.Length() == 0
}

func (Dl *DoublyLinkedList[T]) Remove(index int) (value T, err error) {
	if 0 > index || index >= Dl.numberStored {
		return value, errors.New("the index is out of bounds")
	}
	var result T
	if Dl.numberStored == 1 {
		result = Dl.headNode.Value
		Dl.headNode = nil
		Dl.tailNode = nil
	} else {
		if index == 0 {
			result = Dl.headNode.Value
			Dl.headNode = Dl.headNode.Next
			Dl.headNode.Prev = nil
		} else if index == Dl.numberStored-1 {
			result = Dl.tailNode.Value
			Dl.tailNode = Dl.tailNode.Prev
			Dl.tailNode.Next = nil
		} else if index >= Dl.numberStored/2 {
			currentNode := Dl.tailNode
			for i := Dl.numberStored - 1; index != i; i-- {
				currentNode = currentNode.Prev
			}
			result = currentNode.Value
			currentNode.Prev.Next = currentNode.Next
			currentNode.Next.Prev = currentNode.Prev
		} else {
			currentNode := Dl.headNode
			for i := 0; i != index; i++ {
				currentNode = currentNode.Next
			}
			result = currentNode.Value
			currentNode.Prev.Next = currentNode.Next
			currentNode.Next.Prev = currentNode.Prev
		}
	}
	Dl.numberStored--
	return result, nil
}

func (Dl *DoublyLinkedList[T]) Length() int { return Dl.numberStored }
